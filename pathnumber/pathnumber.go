package pathnumber

import (
	"errors"
	"log"
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

func (list *LinkedList) reverse() *LinkedList {
	if (*list).lenght == 0 {
		return nil
	}
	newList := &LinkedList{}

	firstNode := (*list).start
	newList.add((*firstNode).value)

	actualNode := (*firstNode).next
	for count := (*list).lenght - 1; count > 0; count-- {
		newList.insert(0, (*actualNode).value)
		actualNode = (*actualNode).next
	}
	return newList
}

type Vertex struct {
	index   int
	start   int
	end     int
	visited bool
	paths   int
}

func topologicalSort(graph []*LinkedList, sort *LinkedList, actualNode *Vertex, time *int) {
	(*time)++
	(*actualNode).start = (*time)
	actualLinkedNode := (*graph[(*actualNode).index]).start
	for count := (*graph[(*actualNode).index]).lenght; count > 0; count-- {
		newVertex := (*actualLinkedNode).value
		if !(*newVertex).visited {
			(*newVertex).visited = !(*newVertex).visited
			topologicalSort(graph, sort, newVertex, time)
		}
		actualLinkedNode = (*actualLinkedNode).next
	}
	(*time)++
	(*actualNode).end = *time
	sort.add(actualNode)
}

func trasposeG(graph []*LinkedList, vertexs []*Vertex) []*LinkedList {
	newGraph := make([]*LinkedList, len(graph))
	for index := 0; index < len(newGraph); index++ {
		newGraph[index] = &LinkedList{}
	}

	for index, ady := range graph {
		actualLinkedNode := (*ady).start
		for count := (*ady).lenght; count > 0; count-- {
			actualVertex := (*actualLinkedNode).value
			newGraph[(*actualVertex).index].add(vertexs[index])
			actualLinkedNode = (*actualLinkedNode).next
		}
	}
	return newGraph
}

func pathCount(graph []*LinkedList, a *Vertex, b *Vertex, vertexs []*Vertex) int {
	traspose := trasposeG(graph, vertexs)
	sort := &LinkedList{}
	time := 0

	topologicalSort(graph, sort, a, &time)
	// printSortList(sort)
	sort = sort.reverse()
	printSortList(sort)

	(*a).paths = 1
	actualNode := (*((*sort).start)).next
	for count := (*sort).lenght - 1; count > 0; count-- {
		vertex := (*actualNode).value

		pointer := (*traspose[(*vertex).index]).start
		for times := (*traspose[(*vertex).index]).lenght; times > 0; times-- {
			incidor := (*pointer).value
			(*vertex).paths += (*incidor).paths
			pointer = (*pointer).next
		}
		actualNode = (*actualNode).next
	}
	for index, vertex := range vertexs {
		log.Println(index, (*vertex).paths)
	}
	return (*b).paths
}

// TODO - organize code in new files accesible from each other file
func printSortList(list *LinkedList) {
	actualNode := (*list).start
	for count := (*list).lenght; count > 0; count-- {
		vertex := (*actualNode).value

		log.Println((*vertex).index)
		actualNode = (*actualNode).next
	}
}
