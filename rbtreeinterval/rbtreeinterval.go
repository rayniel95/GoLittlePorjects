package rbtreeinterval

import (
	"errors"
)

// NOTE - simulating enums with static variables
var red int8 = 0
var black int8 = 1

type RBTreeNode struct {
	color  int8
	parent *RBTreeNode
	left   *RBTreeNode
	right  *RBTreeNode
	value  int
}

type RBTree struct {
	root *RBTreeNode
}

func initRBTree() *RBTree {
	return &RBTree{}
}

func (tree *RBTree) leftRotation(node *RBTreeNode) {
	temp := (*node).right
	(*node).right = (*temp).left

	if (*temp).left != nil {
		(*((*temp).left)).parent = node
	}
	(*temp).parent = (*node).parent

	if (*node).parent == nil {
		(*tree).root = temp
	} else if node == (*node).parent.left {
		(*node).parent.left = temp
	} else {
		(*node).parent.right = temp
	}
	(*temp).left = node
	(*node).parent = temp
}

func (tree *RBTree) rightRotation(node *RBTreeNode) {
	temp := (*node).left         // temp: x
	(*node).left = (*temp).right // node: y
	// basicamente es un reajuste de los punteros de los nodos
	if (*temp).right != nil {
		(*((*temp).right)).parent = node
	}
	(*temp).parent = (*node).parent

	if (*node).parent == nil {
		(*tree).root = temp
	} else if node == (*node).parent.left {
		(*node).parent.left = temp
	} else {
		(*node).parent.right = temp
	}
	(*temp).right = node
	(*node).parent = temp
}

func (tree *RBTree) insert(key int) {
	var parentNode *RBTreeNode // y
	actualNode := (*tree).root // x
	// lo tipico nos movemos en el arbol hasta llegar al final para insertar
	// el nodo
	for actualNode != nil {
		parentNode = actualNode
		if key < (*actualNode).value {
			actualNode = (*actualNode).left
		} else {
			actualNode = (*actualNode).right
		}
	}
	// se crea el nodo
	newNode := &RBTreeNode{ // z
		color:  red,
		parent: parentNode,
		value:  key,
	}
	// ajuste de punteros
	if parentNode == nil {
		(*tree).root = newNode
	} else if (*newNode).value < (*parentNode).value {
		(*parentNode).left = newNode
	} else {
		(*parentNode).right = newNode
	}
	// se resuelven los problemas que causa insertar un nodo rojo
	tree.insertFixup(newNode)
}

// TODO - revisar bien los casos del nil
func (tree *RBTree) insertFixup(node *RBTreeNode) {
	// mientras el padre es rojo hay un conflicto entre hijo rojo y padre
	// rojo, notar que si el nodo es la raiz su padre va a ser nil y nil es negro,
	// no se entra al ciclo, por tanto preguntamos tambien si es distinto de nil
	// el padre, no solo para no tener un error de una referencia a un puntero nil
	// sino tambien pq nil es negro
	for (*node).parent != nil && (*((*node).parent)).color == red {
		if (*node).parent == (*((*((*node).parent)).parent)).left {
			uncle := (*((*((*node).parent)).parent)).right // y
			// si el tio es rojo hay un nodo abuelo con dos hijos rojos
			// se recolorea y pasamos al abuelo para seguir viendo conflictos
			// notar que al tio se le debe colorear de negro junto al padre, si
			// el tio permaneciera rojo se incumple la propiedad de cantidad de nodos
			// negros para el subarbol con raiz en el abuelo, luego el tio ha de
			// pasar a negro junto con el padre, pero si el abuelo tambien fuera negro
			// se incumple la misma propiedad para el subarbol con raiz en el
			// tatarabuelo por tanto el abuelo pasa a rojo, pero el tatara puede ser rojo
			// por tanto se sigue buscando hacia arriba posibles problemas con el color
			if (*uncle).color == red {
				(*((*node).parent)).color = black // padre y tio seran negros
				(*uncle).color = black
				(*((*((*node).parent)).parent)).color = red // abuelo rojo
				node = (*((*node).parent)).parent           // pasamos al abuelo
			} else if node == (*((*node).parent)).right { // tio negro
				// se hace una rotacion a la izquierda, la idea es convertir
				// este caso en otro para que sea tratado por el proximo elif, con
				// la rotacion el nodo pasa ahora a ser hijo izquierdo, notar que tanto
				// el nodo como el padre permanecen rojos, puesto que el padre es rojo
				// por la condicion de parada del ciclo y el hijo es rojo puesto que o
				// bien fue insertado como rojo o bien es el resultado de aplicar el
				// caso 1 (primer if) y es un abuelo que se pinto de rojo, luego ambos
				// nodos son rojos y despues de rotar permanecen rojos y el padre pasa
				// a ser el hijo y viceversa, notar que obligatoriamente el abuelo
				// en este caso es negro, si fuere rojo siendo el tio negro, entonces
				// el padre tambien tendria que ser negro, notar que la rotacion
				// preserva la cantidad de negros en los path desde la raiz del abuelo
				node = (*node).parent
				tree.leftRotation(node)
			} else if node == (*((*node).parent)).left {
				// this not according with the book but I think its a patch
				// bajo las condiciones de bloque anterior, siendo el nodo hijo
				// izquierdo, el abuelo negro, el tio negro y el padre rojo, al
				// colorear y rotar no solamente se mantiene la propiedad de arbol
				// binario de busqueda sino tambien la de los colores en los nodos
				(*((*node).parent)).color = black
				(*((*((*node).parent)).parent)).color = red
				tree.rightRotation((*((*node).parent)).parent)
			}
		} else {
			uncle := (*((*((*node).parent)).parent)).left // y
			if (*uncle).color == red {
				(*((*node).parent)).color = black
				(*uncle).color = black
				(*((*((*node).parent)).parent)).color = red
				node = (*((*node).parent)).parent
			} else if node == (*((*node).parent)).left {
				node = (*node).parent
				tree.rightRotation(node)
			} else if node == (*((*node).parent)).right {
				// this not according with the book but I think its a patch
				(*((*node).parent)).color = black
				(*((*((*node).parent)).parent)).color = red
				tree.leftRotation((*((*node).parent)).parent)
			}
		}
	}
	(*((*tree).root)).color = black
}

func min(subtree *RBTreeNode) *RBTreeNode {
	if (*subtree).left == nil {
		return subtree
	}
	return min((*subtree).left)
}

// metodo utilizado para sustituir un nodo por otro, el otro pudiera ser nil
func (tree *RBTree) transplant(deletedNode, node *RBTreeNode) {
	// si el padre es nil entonces el nodo que se va a eliminar era la raiz
	// y hay que actualizar el valor con el nuevo nodo que vamos a poner
	if (*deletedNode).parent == nil {
		(*tree).root = node
		// si el nodo que se va a eliminar es hijo izquierdo se le pone al padre como
		// hijo izquierdo el nuevo nodo
	} else if deletedNode == (*((*deletedNode).parent)).left {
		(*((*deletedNode).parent)).left = node
		// igual que el anterior pero en simetrico
	} else {
		(*((*deletedNode).parent)).right = node
	}
	// en caso de que el nodo a insertar como nuevo sea distinto del nil, entonces
	// se le puede poner como padre el padre del nodo que vamos a eliminar
	if node != nil {
		(*node).parent = (*deletedNode).parent
	}
}

func (tree *RBTree) delete(node *RBTreeNode) {
	temp := node // y
	nodeOriginalColor := (*temp).color
	var child *RBTreeNode
	// si el hijo izquierdo es nil (da igual si el derecho tambien lo es)
	// ponemos el hijo derecho en la posicion del nodo a eliminar, notar que el
	// hijo derecho puede ser nil, en este caso el nodo a eliminar seria hoja, se
	// elimina sin problemas, en caso de que el derecho sea distinto de nil esto
	// significa que ese hijo derecho es hoja, se hace el transplante sin problemas
	if (*node).left == nil {
		child = (*node).right // x
		tree.transplant(node, (*node).right)
		//en caso de que el derecho sea nil y el izquierdo no nuevamente similar
		// al caso anterior, el nodo tiene un solo hijo, el hijo izquierdo que es
		// hoja, se hace el trasplante sin problemas
	} else if (*node).right == nil {
		child = (*node).left
		tree.transplant(node, (*node).left)
		// de otra forma el nodo tiene los dos hijos, se busca el sucesor del mismo
		// que vendra siendo el minimo entre sus hijos derechos. se almacena el
		// color del nodo que se va a eliminar pq se le va a poner al sucesor
		// que sera quien tome su lugar, tambien almacenamos el hijo derecho del
		// sucesor (quien no tiene hijo izquierdo) para subirlo a que tome el lugar
		// de este.
	} else {
		temp = min((*node).right)
		nodeOriginalColor = (*temp).color
		child = (*temp).right
		// estas dos lineas nunca las entendi, no les veo sentido mas alla que el
		// de explicar mejor el codigo, si el padre del sucesor es el nodo que
		// vamos a eliminar, esto significa que una vez pongamos el sucesor (hijo
		// derecho del nodo que eliminaremos) en su nuevo lugar, su hijo derecho
		// va a seguir con la referencia correcta, o sea, el hijo derecho del
		// sucesor va a apuntar al nodo correcto si es el sucesor el hijo del
		// nodo a eliminar, no hay que setear nuevamente, no me queda claro pq lo
		// hacen mas alla de separar este caso para mostrar mejor el algoritmo
		if (*temp).parent == node {
			(*child).parent = temp
			// se realiza el trasplante para la posicion del sucesor a el hijo
			// derecho de este (no tiene hijo izquierdo) y se arreglan los punteros
			// para el caso del sucesor con el hijo derecho del nodo que vamos
			// a eliminar, notar que los punteros del anterior padre del sucesor
			// con el anterior hijo derecho de este se arreglan en el trasplante
		} else {
			tree.transplant(temp, (*temp).right)
			(*temp).right = (*node).right
			(*((*temp).right)).parent = temp
		}
		// en caso de que el nodo a eliminar tenga un hijo que no es el propio
		// sucesor se pone el sucesor en el lugar del nodo a eliminar y se
		// hace el trasplante, se arreglan los punteros del hijo izquierdo, al
		// sucesor en su nuevo lugar se le pone el mismo color que el nodo que
		// eliminamos, de esta forma no hay desbalance de blacks, sin embargo
		// si el sucesor es negro si hay desbalance de blacks entre su rama y la
		// de su hermano
		tree.transplant(node, temp)
		(*temp).left = (*node).left
		(*((*temp).left)).parent = temp
		(*temp).color = (*node).color
	}
	// si el sucesor era black hay que arreglar el desbalance de negros creado
	if nodeOriginalColor == black {
		tree.deleteFixup(child)
	}
}

func (tree *RBTree) deleteFixup(node *RBTreeNode) {
	// los casos no son excluyentes
	// la idea de este algoritmo es arreglar el desbalance (ahora hay un negro menos)
	// de negros creado
	// al mover y (sucesor del nodo eliminado) a la posicion del nodo eliminado
	// si y es negro se crea un desbalance entre su rama y la rama de su hermano
	// asi como entre las ramas que parten de los ancestros de y, para arreglar este
	// desbalance hay dos ideas basicas, de alguna forma manteniendo las propiedades
	// del rbtree se anade otro negro a la rama donde estaba y de esta forma es
	// como si nunca se hubiera quitado ninguno otro metodo es quitar un negro de la
	// rama del hermano de y, emparejando ambas ramas en la cantidad de negros y
	// poniendo un negro mas arriba, o sea, subir el negro a su padre y ver
	// si se puede dejar ahi o se debe seguir realizando ajustes, si se
	// realiza la primera solucion podemos parar, sin embargo, de realizar la
	// segunda, subir el negro a su padre, aunque se haya arreglado el desbalance
	// creado entre las ramas de y y su hermano, es posible que si el padre es
	// negro entonces en nuevo negro no se pueda poner en ese nodo, por lo que
	// tenemos el mismo problema, estamos en un nodo negro y tenemos que poner un
	// negro en esa parte del arbol para garantizar la igualdad de negros entre
	// la rama del padre y la de su hermano y las demas ramas que parten de sus
	// ancestros, por lo que nuevamente, intentamos aplicar las soluciones
	// anteriores, o ponemos un negro extra o quitamos un negro a la otra rama
	// y lo ponemos al padre (abuelo de y) si este no es negro, si es negro repetir.
	// la idea utilizada es que al quitar a y de su posicion, si como y era negro
	// su hijo ahora, quien toma su posicion, tendra dos negros, si el hijo es rojo
	// diremos que es rojo-negro y de ser negro diremos que es negro-negro, la idea
	// es intentar poner en alguna parte de esa rama el negro que quitamos o ir
	// subiendo por el arbol intentando igualar la cantidad de negros, si el hijo
	// es rojo, no entramos al ciclo y lo coloreamos de negro, haciendo igual la
	// cantidad de negros, si el nodo ya es negro, entonces este seria negro-negro
	// y aun tenemos que poner el negro en alguna parte. si x, el hijo de y y que
	// tomo el lugar de este, es negro y su hermano es negro con dos hijos negros
	// se puede quitar el negro tanto de x como del hermano, dejando en x un solo
	// negro y al hermano rojo emparejando ahora la cantidad de negros en ambas
	// ramas, pero aun asi la rama del padre ahora tiene un nero menos que la rama
	// del tio de x, por tanto ponemos el negro en el padre, ahora x sera el padre,
	// nuevamente si el
	// padre es rojo lo podemos colorear de negro, si es negro empezamos de nuevo,
	// si el hermano es negro y su hijo derecho rojo, podemos ponerle al hermano
	// el color del padre, al padre y al hijo los pintamos de negro y hacemos una
	// rotacion a la izquierda en el padre, esto genera un nuevo nodo negro padre de
	// x en esta rama, igualando nuevamente la cantidad de negros, terminamos. Si
	// el hermano es negro con dos hijos negros realizamos lo anterior de ponerlo
	// de rojo y subir el negro para el padre en donde repetimos el ciclo, si de
	// esta forma llegamos a la raiz ya esta es negra y abremos terminado, notar
	// que si llegamos a la raiz de esta forma fuimos eliminando negros mientras
	// subiamos. los casos 1 se transforman en casos 2, 3, 4 mediante recolorear
	// y rotaciones y los casos 3 se transforman en casos 4.
	for node != (*tree).root && (*node).color == black {
		if node == (*((*node).parent)).left {
			brother := (*((*node).parent)).right // w
			// este caso mediante recolorear y rotar se transforma en uno 2, 3 o 4
			if (*brother).color == red { // caso 1
				(*brother).color = black
				(*((*node).parent)).color = red
				tree.leftRotation((*node).parent)
				brother = (*((*node).parent)).right
			}
			// en este caso brother es negro, caso 2
			// pintamos el nodo hermano de rojo y le pasamos el negro al padre
			// si el padre es rojo pues se termina el ciclo y al final lo pintamos
			// de negro, manteniendo la cantidad de negros que habian anteriormente
			// si el nodo padre es negro pues entonces sera doble negro y tenemos
			// el mismo problema que al inicio donde x era tambien doble negro
			if (*((*brother).left)).color == black && (*((*brother).right)).color == black {
				(*brother).color = red
				node = (*node).parent
				// el hijo derecho de brother es negro y el izquierdo rojo, brother
				// es negro, despues de recolorear y hacer la rotacion este caso se
				// transforma en un caso 4 manteniendo la cantidad de negros
				// existente anteriormente
			} else if (*((*brother).right)).color == black { // caso 3
				(*((*brother).left)).color = black
				(*brother).color = red
				tree.rightRotation(brother)
				brother = (*((*node).parent)).right
			}
			// en esta caso se pone el color del padre al hermano, el padre pasa a
			// ser negro junto al hijo derecho, se hace una rotacion en el padre
			// a la izquierda, esto garantiza poner un nodo negro como padre de x,
			// en la posicion donde antes estaba y, sin alterar la cantidad de
			// negros en la otra rama ni las propiedades del rbtree
			if (*((*brother).right)).color == red { // caso 4
				(*brother).color = (*((*node).parent)).color
				(*((*node).parent)).color = black
				(*((*brother).right)).color = black
				tree.leftRotation((*node).parent)
				node = (*tree).root
			}
		} else {
			brother := (*((*node).parent)).left // w
			if (*brother).color == red {        // caso 1
				(*brother).color = black
				(*((*node).parent)).color = red
				tree.rightRotation((*node).parent)
				brother = (*((*node).parent)).left
			}
			// en este caso brother es negro
			if (*((*brother).right)).color == black && (*((*brother).left)).color == black {
				(*brother).color = red
				node = (*node).parent
				// el hijo derecho de brother es negro y el izquierdo rojo, brother
				// es negro, despues de recolorear y hacer la rotacion este caso se
				// transforma en un caso 4 manteniendo la cantidad de negros
				// existente anteriormente
			} else if (*((*brother).left)).color == black { // caso 3
				(*((*brother).right)).color = black
				(*brother).color = red
				tree.leftRotation(brother)
				brother = (*((*node).parent)).left
			}

			if (*((*brother).left)).color == red { // caso 4
				(*brother).color = (*((*node).parent)).color
				(*((*node).parent)).color = black
				(*((*brother).left)).color = black
				tree.rightRotation((*node).parent)
				node = (*tree).root
			}
		}
	}
	(*node).color = black
}

func insideInterval(actualNode *RBTreeNode, interval *[2]int, inside *LinkedList) {
	if (*actualNode).value < (*interval)[0] && (*actualNode).right != nil {
		insideInterval((*actualNode).right, interval, inside)
	} else if (*actualNode).value > (*interval)[1] && (*actualNode).left != nil {
		insideInterval((*actualNode).left, interval, inside)
	} else {
		inside.add(actualNode)
		if (*actualNode).left != nil {
			left((*actualNode).left, interval, inside)
		}
		if (*actualNode).right != nil {
			right((*actualNode).right, interval, inside)
		}
	}
}

func addAll(actualNode *RBTreeNode, inside *LinkedList) {
	inside.add(actualNode)
	if (*actualNode).left != nil {
		addAll((*actualNode).left, inside)
	}
	if (*actualNode).right != nil {
		addAll((*actualNode).right, inside)
	}
}

func right(actualNode *RBTreeNode, interval *[2]int, inside *LinkedList) {
	if (*actualNode).value <= (*interval)[1] && (*actualNode).value >= (*interval)[0] {
		inside.add(actualNode)
		if (*actualNode).left != nil {
			addAll((*actualNode).left, inside)
		}
		if (*actualNode).right != nil {
			right((*actualNode).right, interval, inside)
		}
		return
	}
	if (*actualNode).left != nil {
		right((*actualNode).left, interval, inside)
	}
}

func left(actualNode *RBTreeNode, interval *[2]int, inside *LinkedList) {
	if (*actualNode).value <= (*interval)[1] && (*actualNode).value >= (*interval)[0] {
		inside.add(actualNode)
		if (*actualNode).right != nil {
			addAll((*actualNode).right, inside)
		}
		if (*actualNode).left != nil {
			left((*actualNode).left, interval, inside)
		}
		return
	}
	if (*actualNode).right != nil {
		left((*actualNode).right, interval, inside)
	}
}

type LinkedListNode struct {
	next  *LinkedListNode
	value *RBTreeNode
}

type LinkedList struct {
	start  *LinkedListNode
	end    *LinkedListNode
	lenght int
}

func (list *LinkedList) add(value *RBTreeNode) {
	node := &LinkedListNode{value: value}
	if (*list).lenght == 0 {
		(*list).start = node
		(*list).end = node
		(*list).lenght = 1
		return
	}
	(*((*list).end)).next = node
	(*list).end = node
	(*list).lenght++
}

func (list *LinkedList) pop() (*LinkedListNode, error) {
	if (*list).lenght == 0 {
		return nil, errors.New("can not pop from a linkedlist with lenght 0")
	}
	node := (*list).start
	for index := 1; index < (*list).lenght-1; index++ {
		node = (*node).next
	}
	(*list).lenght--
	if (*list).lenght == 0 {
		(*list).start = nil
		(*list).end = nil
		return node, nil
	}
	(*list).end = node
	node = node.next
	(*((*list).end)).next = nil
	return node, nil
}

func (list *LinkedList) popIndex(popIndex int) (*LinkedListNode, error) {
	if popIndex >= (*list).lenght || popIndex < 0 {
		return nil, errors.New("index out of range")
	}
	var node *LinkedListNode
	if popIndex == 0 {
		node = (*list).start
		(*list).start = (*node).next
		(*list).lenght--
		if (*list).lenght == 0 {
			(*list).end = nil
		}
		return node, nil
	}
	node = (*list).start
	for index := 1; index < popIndex; index++ {
		node = node.next
	}

	res := (*node).next
	(*node).next = (*res).next
	(*list).lenght--
	if popIndex == (*list).lenght {
		(*list).end = node
	}
	return res, nil
}

func (list *LinkedList) insert(insertIndex int, value *RBTreeNode) error {
	if insertIndex >= (*list).lenght || insertIndex < 0 {
		return errors.New("index out of range")
	}
	newNode := &LinkedListNode{value: value}
	(*list).lenght++
	if insertIndex == 0 {
		(*newNode).next = (*list).start
		(*list).start = newNode
		return nil
	}

	node := (*list).start
	for index := 1; index < insertIndex; index++ {
		node = node.next
	}
	(*newNode).next = (*node).next
	(*node).next = newNode
	return nil
}
