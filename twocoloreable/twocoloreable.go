package twocoloreable

import "errors"

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

func (list *LinkedList) popFirst() (*LinkedListNode, error) {
	return list.popIndex(0)
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

func (list *LinkedList) addList(other *LinkedList) {
	if (*other).lenght == 0 {
		return
	}
	(*list).lenght += (*other).lenght
	(*((*list).end)).next = (*other).start
	(*list).end = (*other).end
}

type Vertex struct {
	value int
	index int
	color int
}

const black = 0
const white = 1
const noColor = 2

func bfs(graph []*LinkedList, initial *Vertex) {
	queue := &LinkedList{}

	(*initial).color = white

	queue.add(initial)

	for (*queue).lenght > 0 {
		actualNode, _ := queue.popFirst()
		actualVertex := (*actualNode).value

		for node := (*graph[(*actualVertex).index]).start; node != nil; node = (*node).next {
			vertex := (*node).value

			if (*vertex).color == noColor {
				if (*actualVertex).color == white {
					(*vertex).color = black
				} else {
					(*vertex).color = white
				}
				queue.add(vertex)
			}
		}
	}
}

func isTwoColor(graph []*LinkedList, vertex []*Vertex) bool {
	for index, actualList := range graph {
		for node := (*actualList).start; node != nil; node = (*node).next {
			actualVertex := (*node).value
			if (*actualVertex).color != (*vertex[index]).color {
				return false
			}
		}
	}
	return true
}

func isTwoColoreable(graph []*LinkedList) bool {
	vertexs := make([]*Vertex, len(graph))
	for index := 0; index < len(vertexs); index++ {
		vertexs[index] = &Vertex{index: index}
	}

	bfs(graph, vertexs[0])
	return isTwoColor(graph, vertexs)
}
