package heap

import (
	"reflect"
	"testing"
)

func Test_newHeapNode(t *testing.T) {
	type args struct {
		value    *interface{}
		priority uint32
	}
	tests := []struct {
		name string
		args args
		want *HeapNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHeapNode(tt.args.value, tt.args.priority); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHeapNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapNode_heapifyDown(t *testing.T) {
	type fields struct {
		cell   *Cell
		left   *HeapNode
		right  *HeapNode
		parent *HeapNode
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
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
			if err := node.heapifyDown(); (err != nil) != tt.wantErr {
				t.Errorf("HeapNode.heapifyDown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
			tree.update(tt.args.node)
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
			got, err := tree.peek()
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
			tree.add(tt.args.value, tt.args.priority)
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
			got, err := tree.deleteMin()
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
