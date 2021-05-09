package ABB

import (
	"errors"
	"log"
)

type ABB struct {
	left      *ABB
	right     *ABB
	successor *ABB
	value     int
}

func (node *ABB) isLeaf() bool {
	return (*node).left == nil && (*node).right == nil
}

func searchAnteccessorLeftSon(root *ABB) *ABB {
	if (*root).right == nil {
		return root
	}
	return searchAnteccessorLeftSon((*root).right)
}

func searchAnteccessorNotLeftSon(
	root *ABB, value int, anteccessor *(*ABB),
) error {
	if (*root).value == value {
		return nil
	}
	if (*root).value < value && (*root).right != nil {
		if *anteccessor == nil || (*root).value > (*(*anteccessor)).value {
			*anteccessor = root
		}
		return searchAnteccessorNotLeftSon((*root).right, value, anteccessor)
	}
	if (*root).value > value && (*root).left != nil {
		return searchAnteccessorNotLeftSon((*root).left, value, anteccessor)
	}
	return errors.New("")
}

func searchParent(root *ABB, value int, parent *(*ABB)) error {
	if (*root).value == value {
		return nil
	}
	if (*root).value > value && (*root).left != nil {
		*parent = root
		return searchParent((*root).left, value, parent)
	}
	if (*root).value < value && (*root).right != nil {
		*parent = root
		return searchParent((*root).right, value, parent)
	}
	return errors.New("")
}

func search(root *ABB, value int) *ABB {
	if (*root).value == value {
		return root
	}
	if (*root).value < value && (*root).right != nil {
		return search((*root).right, value)
	}
	if (*root).value > value && (*root).left != nil {
		return search((*root).left, value)
	}
	return nil
}

func delete(root *(*ABB), value int) {
	node := search(*root, value)

	var temp1 *ABB = nil
	var parent (*(*ABB)) = &temp1
	searchParent(*root, value, parent)

	var temp2 *ABB = nil
	var anteccessor (*(*ABB)) = &temp2

	if node.isLeaf() {
		// si es hoja es muy simple se busca el antecesor que sera ancestro
		searchAnteccessorNotLeftSon(*root, value, anteccessor)
		if *anteccessor != nil {
			(*(*anteccessor)).successor = (*node).successor
		}
		// si es hoja se actualiza la raiz
		if *parent == nil {
			*root = nil
		}
		// se actualiza el padre
		if *parent != nil {
			if (*(*parent)).left == node {
				(*(*parent)).left = nil
				return
			}
			if (*(*parent)).right == node {
				(*(*parent)).right = nil
			}
		}
		return
	}
	// tiene hijo derecho (sino hubiera sido hoja)
	if (*node).left == nil {
		// se busca el antecesor que debe ser ancestro en este caso
		searchAnteccessorNotLeftSon(*root, value, anteccessor)
		if *anteccessor != nil {
			(*(*anteccessor)).successor = (*node).successor
		}
		// si es la raiz se actualiza la raiz
		if *parent == nil {
			*root = (*node).right
			return
		}
		// de no ser la raiz se sube el hijo derecho al lugar del nodo que
		// se desea eliminar
		if (*(*parent)).left == node {
			(*(*parent)).left = (*node).right
		}

		if (*(*parent)).right == node {
			(*(*parent)).right = (*node).right
		}
		return
	}
	// tiene hijo izquierdo, tiene antecesor que es descendiente de dicho
	// hijo o es el propio hijo

	// se busca el antecesor, se actulizan los sucesores y se le pone como
	// hijo derecho al antecesor el hijo derecho del nodo a eliminar, el
	// antecesor del nodo si es descendiente del hijo izquierdo no va a tener
	// hijo derecho y se mantiene la invariante del abb
	*anteccessor = searchAnteccessorLeftSon((*node).left)
	(*(*anteccessor)).successor = (*node).successor
	(*(*anteccessor)).right = (*node).right
	// si es la raiz se actualiza esta poniendo al hijo izquierdo como nueva
	// raiz
	if *parent == nil {
		*root = (*node).left
		log.Println((*(*root)).value)
		return
	}
	// de no ser la raiz se sube al hijo izquierdo en el lugar del nodo
	if (*(*parent)).left == node {
		(*(*parent)).left = (*node).left
	}
	if (*(*parent)).right == node {
		(*(*parent)).right = (*node).left
	}
}

func insert(root *ABB, value int) *ABB {
	var anteccessor *ABB
	return insertA(root, value, anteccessor)
}

func insertA(root *ABB, value int, anteccessor *ABB) *ABB {
	// si el hijo izquierdo es nil y es menor lo pongo en el hijo izquierdo
	if (*root).value > value && (*root).left == nil {
		node := &ABB{
			value: value,
		}

		if anteccessor != nil {
			// actualizo sucesores si no soy el menor en el arbol
			(*node).successor = (*anteccessor).successor
			(*anteccessor).successor = node

		} else {
			// soy el menor en el arbol no existe antecesor, eso significa
			// que soy el nodo mas a la izquierda mi padre es mi sucesor
			(*node).successor = root
		}
		(*root).left = node
		return node
	}
	// si el hijo derecho es nil y es mayor lo pongo en el derecho
	if (*root).value < value && (*root).right == nil {
		// si el antecesir en nil o el valor actual es mayor que el del
		// antecesor actualizo mi antecesor
		if anteccessor == nil || (*root).value > (*anteccessor).value {
			anteccessor = root
		}
		node := &ABB{
			value:     value,
			successor: (*anteccessor).successor,
		}
		// actualizo sucesores
		(*anteccessor).successor = node
		(*root).right = node
		return node
	}
	// si el valor es menor y tiene hijo izquierdo me muevo para el mismo
	if (*root).value > value {
		return insertA((*root).left, value, anteccessor)
	}
	// si el valor es menor y tiene hijo derecho me muevo para el mismo
	if (*root).value < value {
		// si el antecesor es nil o el valor del antecesor es menor que el
		// valor del nodo actual tengo un nuevo antecesor y me quedo con el
		if anteccessor == nil || (*anteccessor).value < (*root).value {
			anteccessor = root
		}
		return insertA((*root).right, value, anteccessor)
	}
	return nil
}

func equal(node, other *ABB) bool {
	if ((*node).left != nil && (*other).left == nil) ||
		((*node).right != nil && (*other).right == nil) ||
		((*other).left != nil && (*node).left == nil) ||
		((*other).right != nil && (*node).right == nil) {
		return false
	}

	if (*node).value != (*other).value {
		return false
	}

	if (*node).left != nil && !equal((*node).left, (*other).left) {
		return false
	}

	if (*node).right != nil && !equal((*node).right, (*other).right) {
		return false
	}
	return true
}

func countNodes(root *ABB) int {
	if root.isLeaf() {
		return 1
	}
	cantity := 0
	if (*root).left != nil {
		cantity += countNodes((*root).left)
	}
	if (*root).right != nil {
		cantity += countNodes((*root).right)
	}
	return 1 + cantity
}

func successorCheck(root *ABB) bool {
	actualNode := root
	cantity := 1
	// fmt.Println((*actualNode).value)
	for ; (*actualNode).left != nil; actualNode = (*actualNode).left {
	}
	for ; (*actualNode).successor != nil &&
		(*((*actualNode).successor)).value >=
			(*actualNode).value; actualNode = (*actualNode).successor {
		// fmt.Println((*((*actualNode).successor)).value)
		cantity++
	}
	// fmt.Println(cantity)
	return cantity == countNodes(root)
}

func setSuccessors(root *ABB) {
	list := make([]*ABB, 0)
	entreOrden(root, &list)
	// for _, node := range list {
	// 	log.Println((*node).value)
	// }

	for index := 0; index < len(list)-1; index++ {
		(*list[index]).successor = list[index+1]
		// log.Println((*list[index]).value)
		// log.Println((*((*list[index]).successor)).value)
	}
}

func entreOrden(root *ABB, list *[]*ABB) {
	if (*root).left != nil {
		entreOrden((*root).left, list)
	}

	*list = append(*list, root)

	if (*root).right != nil {
		entreOrden((*root).right, list)
	}
}
