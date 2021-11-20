package heap

type Cell struct {
	value    interface{}
	priority uint32
}

type HeapNode struct {
	cell   Cell
	left   *HeapNode
	right  *HeapNode
	parent *HeapNode
}
