package bfs

import (
	"errors"
)

const white = 0
const grey = 1
const black = 2

type LinkedListNode struct {
	next  *LinkedListNode
	value *Vertex
}

type LinkedList struct {
	start  *LinkedListNode
	end    *LinkedListNode
	lenght int
}

func (list *LinkedList) popFirst() (*LinkedListNode, error) {
	if (*list).lenght == 0 {
		return nil, errors.New("can not pop from a linkedlist with lenght 0")
	}
	(*list).lenght--
	node := (*list).start
	(*list).start = (*node).next
	if (*list).start == nil {
		(*list).end = nil
	}
	return node, nil
}

func (list *LinkedList) add(value *Vertex) {
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

func (list *LinkedList) insert(insertIndex int, value *Vertex) error {
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

type adymatrix struct {
	list [][]*Vertex
}

type Vertex struct {
	value int
	index int
	color int
}

func newAdyGraph() *adymatrix {
	return &adymatrix{make([][]*Vertex, 0)}
}

func bfsAdyMatrix(graph [][]*Vertex, initial *Vertex) []int {
	queue := &LinkedList{}

	distances := make([]int, len(graph))

	queue.add(initial)

	for (*queue).lenght > 0 {
		actualLinkedNode, _ := queue.popFirst()
		actualVertex := (*actualLinkedNode).value

		for _, vertex := range graph[(*actualVertex).index] {
			if vertex != nil && (*vertex).color == white {
				(*vertex).color = grey
				distances[(*vertex).index] = distances[(*actualVertex).index] + 1
				queue.add(vertex)
			}
		}
		(*actualVertex).color = black
	}
	return distances
}
