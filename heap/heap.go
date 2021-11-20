package heap

import "errors"

type Cell struct {
	value    *interface{}
	priority uint32
}

type HeapNode struct {
	cell   *Cell
	left   *HeapNode
	right  *HeapNode
	parent *HeapNode
}

func newHeapNode(value *interface{}, priority uint32) *HeapNode {
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

func (tree *HeapTree) update(node *HeapNode) {
	node.heapifyDown()
	node.heapifyUp()
}
