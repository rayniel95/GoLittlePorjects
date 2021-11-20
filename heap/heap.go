package heap

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
