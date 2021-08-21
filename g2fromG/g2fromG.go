package g2fromG

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

func foundIncidenceList(graph [][]bool, vertex []*Vertex) []*LinkedList {
	incidenceList := make([]*LinkedList, len(graph))

	for index := 0; index < len(incidenceList); index++ {
		incidenceList[index] = &LinkedList{}
	}
	for firstIndex := 0; firstIndex < len(graph); firstIndex++ {
		for secondIndex := 0; secondIndex < len(graph[0]); secondIndex++ {
			if graph[firstIndex][secondIndex] {
				incidenceList[firstIndex].add(vertex[secondIndex])
			}
		}
	}
	return incidenceList
}

func secondLevelIncidence(incidenceList []*LinkedList) []*LinkedList {
	secondLevel := make([]*LinkedList, len(incidenceList))

	for index := 0; index < len(secondLevel); index++ {
		secondLevel[index] = &LinkedList{}
	}
	for index := 0; index < len(secondLevel); index++ {
		actualVertex := (*incidenceList[index]).start
		for cantity := (*incidenceList[index]).lenght; cantity > 0; cantity-- {
			secondLevel[index].addList(incidenceList[(*((*actualVertex).value)).index])
			actualVertex = (*actualVertex).next
		}
	}
	return secondLevel
}

func DFS(list *LinkedList, index int, g2 [][]bool, incidenceList []*LinkedList) {
	actualNode := (*list).start
	for count := (*list).lenght; count > 0; count-- {
		if (*((*actualNode).value)).color == 0 {
			(*((*actualNode).value)).color = 1
			g2[index][(*((*actualNode).value)).index] = true
			DFS(incidenceList[(*((*actualNode).value)).index], index, g2, incidenceList)
		}

	}
}

func buildG2(secondLevel []*LinkedList, vertex []*Vertex, incidenceList []*LinkedList) [][]bool {
	g2 := make([][]bool, len(secondLevel))
	for index := 0; index < len(g2); index++ {
		g2[index] = make([]bool, len(secondLevel))
	}
	for index := 0; index < len(secondLevel); index++ {
		DFS(secondLevel[index], index, g2, incidenceList)
		for _, node := range vertex {
			(*node).color = 0
		}
	}
	return g2
}

func g2(graph [][]bool) [][]bool {
	vertexs := make([]*Vertex, len(graph))
	for index := 0; index < len(vertexs); index++ {
		vertexs[index] = &Vertex{index: index}
	}
	incidenceList := foundIncidenceList(graph, vertexs)
	secondLevelIncidence := secondLevelIncidence(incidenceList)

	return buildG2(secondLevelIncidence, vertexs, incidenceList)
}
