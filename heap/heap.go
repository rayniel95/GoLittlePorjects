package heap

import (
	"errors"
	"fmt"
)

type Cell struct {
	value    interface{}
	priority uint32
}

func (cell *Cell) String() string {
	return fmt.Sprintf("[%v, %d]", (*cell).value, (*cell).priority)
}

type HeapNode struct {
	cell   *Cell
	left   *HeapNode
	right  *HeapNode
	parent *HeapNode
}

func NewHeapNode(value interface{}, priority uint32) *HeapNode {
	return &HeapNode{
		cell: &Cell{value: value, priority: priority},
	}
}

func (node *HeapNode) heapifyDown() error {
	var small *HeapNode

	if (*node).left != nil && (*node).right != nil {
		if (*((*((*node).left)).cell)).priority < (*((*((*node).right)).cell)).priority {
			small = (*node).left
		} else {
			small = (*node).right
		}
	} else if (*node).left != nil {
		small = (*node).left
	} else if (*node).right != nil {
		return errors.New("tree do not have correct structure")
	}

	if small != nil && (*((*node).cell)).priority > (*((*small).cell)).priority {
		temp := (*small).cell
		(*small).cell = (*node).cell
		(*node).cell = temp
		if small.heapifyDown() != nil {
			return errors.New("tree do not have correct structure")
		}
	}
	return nil
}

func (node *HeapNode) heapifyUp() {
	if (*node).parent != nil && (*((*((*node).parent)).cell)).priority > (*((*node).cell)).priority {
		temp := (*((*node).parent)).cell
		(*((*node).parent)).cell = (*node).cell
		(*node).cell = temp
		(*node).parent.heapifyUp()
	}
}

func (tree *HeapTree) Update(node *HeapNode) {
	node.heapifyDown()
	node.heapifyUp()
}

type LinkedNode struct {
	next *LinkedNode
	prev *LinkedNode
	val  *HeapNode
}

type HeapTree struct {
	start        *LinkedNode
	end          *LinkedNode
	parentOfLast *LinkedNode
	len          uint32
}

func (tree *HeapTree) Peek() (error, *Cell) {
	if (*tree).start != nil {
		return nil, (*((*((*tree).start)).val)).cell
	}
	return errors.New("Empty Heap"), nil
}

func (tree *HeapTree) Add(value interface{}, priority uint32) {
	if (*tree).start == nil {
		newLinkedNode := &LinkedNode{
			val: NewHeapNode(value, priority),
		}
		(*tree).start = newLinkedNode
		(*tree).end = newLinkedNode
		(*tree).len++
		return
	}
	if (*tree).start == (*tree).end {
		(*tree).parentOfLast = (*tree).start
		parentHeapNode := (*((*tree).start)).val
		newHeapNode := &HeapNode{
			cell: &Cell{
				value:    value,
				priority: priority,
			},
			parent: parentHeapNode,
		}
		(*parentHeapNode).left = newHeapNode
		newLinkedNode := &LinkedNode{
			prev: (*tree).start,
			val:  newHeapNode,
		}
		(*((*tree).start)).next = newLinkedNode
		(*tree).end = newLinkedNode
		(*tree).len++
		newHeapNode.heapifyUp()
		return
	}

	newHeapNode := &HeapNode{
		cell: &Cell{
			value:    value,
			priority: priority,
		},
	}

	parent := (*((*tree).parentOfLast)).val
	if (*parent).right == nil {
		(*parent).right = newHeapNode
	} else {
		(*tree).parentOfLast = (*((*tree).parentOfLast)).next
		parent = (*((*tree).parentOfLast)).val
		(*parent).left = newHeapNode
	}
	(*newHeapNode).parent = parent

	newLinkedNode := &LinkedNode{
		prev: (*tree).end,
		val:  newHeapNode,
	}
	(*((*tree).end)).next = newLinkedNode
	(*tree).end = newLinkedNode
	(*tree).len++
	newHeapNode.heapifyUp()
}

func (tree *HeapTree) DeleteMin() (error, *Cell) {
	if (*tree).start == nil {
		return errors.New("empty heap"), nil
	}
	(*tree).len--
	_, min := (*tree).Peek()
	lastNode := (*((*tree).end)).val
	valueOfLast := (*lastNode).cell

	parentLink := (*tree).parentOfLast
	if parentLink != nil {
		parent := (*parentLink).val
		if (*parent).right == lastNode {
			(*parent).right = nil
		} else if (*parent).left == lastNode {
			(*parent).left = nil
			(*tree).parentOfLast = (*((*tree).parentOfLast)).prev
		} else {
			return errors.New("bad heap structure"), nil
		}
	}
	(*tree).end = (*((*tree).end)).prev
	if (*tree).end != nil {
		(*((*tree).end)).next = nil
	} else {
		(*tree).start = nil
	}

	if (*tree).start != nil {
		(*((*((*tree).start)).val)).cell = valueOfLast
		(*((*tree).start)).val.heapifyDown()
	}
	return nil, min
}

func (tree *HeapTree) Len() uint32 {
	return (*tree).len
}
