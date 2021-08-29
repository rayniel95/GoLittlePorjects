package poptopologicalsort

import (
	"errors"
)

type LinkedListNode struct {
	next  *LinkedListNode
	value *Vertex
}

type LinkedList struct {
	start  *LinkedListNode
	end    *LinkedListNode
	lenght int
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

type Vertex struct {
	index    int
	indegree int
	visited  bool
}

func inDegree(graph []*LinkedList) {
	for _, ady := range graph {
		actualNode := (*ady).start

		for times := (*ady).lenght; times > 0; times-- {
			vertex := (*actualNode).value
			(*vertex).indegree++
			actualNode = (*actualNode).next
		}
	}
}

func topologicalSortInDegree(graph []*LinkedList, vertexs []*Vertex) *LinkedList {
	topologicalSort := &LinkedList{}
	queue := &LinkedList{}

	for _, vertex := range vertexs {
		if (*vertex).indegree == 0 {
			queue.add(vertex)
		}
	}
	for (*queue).lenght > 0 {
		actualNode, _ := queue.popIndex(0)
		actualVertex := (*actualNode).value

		topologicalSort.add(actualVertex)
		actualNode = (*graph[(*actualVertex).index]).start
		for count := (*graph[(*actualVertex).index]).lenght; count > 0; count-- {
			vertex := (*actualNode).value
			(*vertex).indegree--
			if (*vertex).indegree == 0 {
				queue.add(vertex)
			}
			actualNode = (*actualNode).next
		}
	}
	return topologicalSort
}
