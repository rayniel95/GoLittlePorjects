package ABBE

import "fmt"

type IComparable interface {
	compare(other interface{}) int
}

type abbeNode struct {
	left  *abbeNode
	right *abbeNode
	value IComparable
	/*
		NOTE - por algun motivo tiene mas sentido que no sea un puntero a
		IComparable, generalmente la interfaz la va a implementar un puntero a
		un struct o a un tipo. Entonces el valor concreto de esa interfaz seria
		un puntero.
	*/
}

type frequency struct {
	value   IComparable
	cantity int
}

func (node *abbeNode) insert(treeNode *abbeNode) {
	if (*node).left == nil && (*node).right == nil {
		if (*node).value.compare((*treeNode).value) == 1 {
			(*node).left = &abbeNode{value: (*node).value}
			(*node).right = treeNode
			return
		}
		if (*node).value.compare((*treeNode).value) == -1 {
			(*node).left = treeNode
			(*node).right = &abbeNode{value: (*node).value}
			return
		}
		(*node).left = treeNode
		(*node).right = &abbeNode{value: (*node).value}
		return
	}
	if (*node).value.compare((*treeNode).value) == 1 ||
		(*node).value.compare((*treeNode).value) == 0 {

		(*node).right.insert(treeNode)
		return
	}
	if (*node).value.compare((*treeNode).value) == -1 {
		(*node).left.insert(treeNode)
		return
	}
}

func (node *abbeNode) frecuencyPerValue() map[IComparable]int {
	frecuency := make(map[IComparable]int)

	for _, value := range node.values() {
		frecuency[value]++
	}
	return frecuency
}

func (node *abbeNode) values() []IComparable {
	if (*node).left == nil && (*node).right == nil {
		return []IComparable{(*node).value}
	}
	list := make([]IComparable, 0)
	list = append(list, (*node).left.values()...)

	list = append(list, (*node).value)

	list = append(list, (*node).right.values()...)
	return list
}

type myInt int

func (inte *myInt) compare(other interface{}) int {
	value, ok := other.(*myInt)
	if !ok {
		// panic
	}
	if *inte < *value {
		return 1
	}
	if *inte == *value {
		return 0
	}
	return -1
}

func main() {
	var treeVal4 myInt = 4
	var treeVal3 myInt = 3
	var treeVal8 myInt = 8
	var treeVal5 myInt = 5
	var treeVal2 myInt = 2

	tree := &abbeNode{value: &treeVal4}
	tree.insert(&abbeNode{value: &treeVal3})
	tree.insert(&abbeNode{value: &treeVal8})
	tree.insert(&abbeNode{value: &treeVal5})
	tree.insert(&abbeNode{value: &treeVal2})

	pointers := tree.values()
	values := make([]myInt, len(pointers))

	for index := 0; index < len(values); index++ {
		values[index] = *pointers[index].(*myInt)
	}
	fmt.Println(values)
}
