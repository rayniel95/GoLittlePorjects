package heap

import (
	"reflect"
	"testing"
)

func TestHeapNode_heapifyDownMiddleTreeValue(t *testing.T) {
	nodes := createNodes(1001)

	addLeftSon(nodes[1], nodes[2])
	addRightSon(nodes[1], nodes[1000])

	addLeftSon(nodes[2], nodes[3])
	addRightSon(nodes[2], nodes[4])

	addLeftSon(nodes[3], nodes[7])
	addRightSon(nodes[3], nodes[8])

	addLeftSon(nodes[4], nodes[9])
	addRightSon(nodes[4], nodes[10])

	addLeftSon(nodes[1000], nodes[5])
	addRightSon(nodes[1000], nodes[6])

	addLeftSon(nodes[5], nodes[11])
	addRightSon(nodes[5], nodes[12])

	addLeftSon(nodes[6], nodes[13])
	addRightSon(nodes[6], nodes[14])

	nodes[1000].heapifyDown()

	node2 := createNodes(1001)

	addLeftSon(node2[1], node2[2])
	addRightSon(node2[1], node2[5])

	addLeftSon(node2[2], node2[3])
	addRightSon(node2[2], node2[4])

	addLeftSon(node2[3], node2[7])
	addRightSon(node2[3], node2[8])

	addLeftSon(node2[4], node2[9])
	addRightSon(node2[4], node2[10])

	addLeftSon(node2[5], node2[11])
	addRightSon(node2[5], node2[6])

	addLeftSon(node2[11], node2[1000])
	addRightSon(node2[11], node2[12])

	addLeftSon(node2[6], node2[13])
	addRightSon(node2[6], node2[14])

	if !reflect.DeepEqual(nodes[1], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func addLeftSon(node, son *HeapNode) {
	(*node).left = son
	(*son).parent = node
}

func addRightSon(node, son *HeapNode) {
	(*node).right = son
	(*son).parent = node
}

func createNodes(number int) []*HeapNode {
	nodes := make([]*HeapNode, number)

	for index := 0; index < number; index++ {
		nodes[index] = &HeapNode{
			cell: &Cell{
				value:    index,
				priority: uint32(index),
			},
		}
	}
	return nodes
}

func TestHeapNode_heapifyUp(t *testing.T) {
	type fields struct {
		cell   *Cell
		left   *HeapNode
		right  *HeapNode
		parent *HeapNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &HeapNode{
				cell:   tt.fields.cell,
				left:   tt.fields.left,
				right:  tt.fields.right,
				parent: tt.fields.parent,
			}
			node.heapifyUp()
		})
	}
}

func TestHeapTree_update(t *testing.T) {
	type fields struct {
		start        *LinkedNode
		end          *LinkedNode
		parentOfLast *LinkedNode
	}
	type args struct {
		node *HeapNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &HeapTree{
				start:        tt.fields.start,
				end:          tt.fields.end,
				parentOfLast: tt.fields.parentOfLast,
			}
			tree.Update(tt.args.node)
		})
	}
}

func TestHeapTree_peek(t *testing.T) {
	type fields struct {
		start        *LinkedNode
		end          *LinkedNode
		parentOfLast *LinkedNode
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Cell
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &HeapTree{
				start:        tt.fields.start,
				end:          tt.fields.end,
				parentOfLast: tt.fields.parentOfLast,
			}
			got, err := tree.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("HeapTree.peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapTree.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapTree_add(t *testing.T) {
	type fields struct {
		start        *LinkedNode
		end          *LinkedNode
		parentOfLast *LinkedNode
	}
	type args struct {
		value    *interface{}
		priority uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &HeapTree{
				start:        tt.fields.start,
				end:          tt.fields.end,
				parentOfLast: tt.fields.parentOfLast,
			}
			tree.Add(tt.args.value, tt.args.priority)
		})
	}
}

func TestHeapTree_deleteMin(t *testing.T) {
	type fields struct {
		start        *LinkedNode
		end          *LinkedNode
		parentOfLast *LinkedNode
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Cell
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &HeapTree{
				start:        tt.fields.start,
				end:          tt.fields.end,
				parentOfLast: tt.fields.parentOfLast,
			}
			got, err := tree.DeleteMin()
			if (err != nil) != tt.wantErr {
				t.Errorf("HeapTree.deleteMin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapTree.deleteMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
