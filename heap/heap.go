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

type LinkedNode struct {
	next *LinkedNode
	prev *LinkedNode
	val  *HeapNode
}

type HeapTree struct {
	start        *LinkedNode
	end          *LinkedNode
	parentOfLast *LinkedNode
}

func (tree *HeapTree) peek() (error, *Cell) {
	if (*tree).start != nil {
		return nil, (*((*((*tree).start)).val)).cell
	}
	return errors.New("Empty Heap"), nil
}

func (tree *HeapTree) add(value *interface{}, priority uint32) {
	if (*tree).start == nil {
		newLinkedNode := &LinkedNode{
			val: newHeapNode(value, priority),
		}
		(*tree).start = newLinkedNode
		(*tree).end = newLinkedNode
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

	newHeapNode.heapifyUp()
}
