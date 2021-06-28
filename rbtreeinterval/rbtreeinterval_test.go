package rbtreeinterval

import (
	"reflect"
	"testing"
)

func Test_initRBTree(t *testing.T) {
	tests := []struct {
		name string
		want *RBTree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initRBTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initRBTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRBTree_leftRotation(t *testing.T) {
	type fields struct {
		root *RBTreeNode
	}
	type args struct {
		node *RBTreeNode
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
			tree := &RBTree{
				root: tt.fields.root,
			}
			tree.leftRotation(tt.args.node)
		})
	}
}

func TestRBTree_rightRotation(t *testing.T) {
	type fields struct {
		root *RBTreeNode
	}
	type args struct {
		node *RBTreeNode
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
			tree := &RBTree{
				root: tt.fields.root,
			}
			tree.rightRotation(tt.args.node)
		})
	}
}

func TestRBTree_insert(t *testing.T) {
	type fields struct {
		root *RBTreeNode
	}
	type args struct {
		key int
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
			tree := &RBTree{
				root: tt.fields.root,
			}
			tree.insert(tt.args.key)
		})
	}
}

func TestRBTree_insertFixup(t *testing.T) {
	type fields struct {
		root *RBTreeNode
	}
	type args struct {
		node *RBTreeNode
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
			tree := &RBTree{
				root: tt.fields.root,
			}
			tree.insertFixup(tt.args.node)
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		subtree *RBTreeNode
	}
	tests := []struct {
		name string
		args args
		want *RBTreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.subtree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRBTree_transplant(t *testing.T) {
	type fields struct {
		root *RBTreeNode
	}
	type args struct {
		deletedNode *RBTreeNode
		node        *RBTreeNode
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
			tree := &RBTree{
				root: tt.fields.root,
			}
			tree.transplant(tt.args.deletedNode, tt.args.node)
		})
	}
}

func TestRBTree_delete(t *testing.T) {
	type fields struct {
		root *RBTreeNode
	}
	type args struct {
		node *RBTreeNode
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
			tree := &RBTree{
				root: tt.fields.root,
			}
			tree.delete(tt.args.node)
		})
	}
}

func TestRBTree_deleteFixup(t *testing.T) {
	type fields struct {
		root *RBTreeNode
	}
	type args struct {
		node *RBTreeNode
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
			tree := &RBTree{
				root: tt.fields.root,
			}
			tree.deleteFixup(tt.args.node)
		})
	}
}

func Test_insideInterval(t *testing.T) {
	type args struct {
		actualNode *RBTreeNode
		interval   *[2]int
		inside     *LinkedList
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insideInterval(tt.args.actualNode, tt.args.interval, tt.args.inside)
		})
	}
}

func Test_addAll(t *testing.T) {
	type args struct {
		actualNode *RBTreeNode
		inside     *LinkedList
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addAll(tt.args.actualNode, tt.args.inside)
		})
	}
}

func Test_right(t *testing.T) {
	type args struct {
		actualNode *RBTreeNode
		interval   *[2]int
		inside     *LinkedList
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			right(tt.args.actualNode, tt.args.interval, tt.args.inside)
		})
	}
}

func Test_left(t *testing.T) {
	type args struct {
		actualNode *RBTreeNode
		interval   *[2]int
		inside     *LinkedList
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left(tt.args.actualNode, tt.args.interval, tt.args.inside)
		})
	}
}

func TestLinkedList_add(t *testing.T) {
	type fields struct {
		start  *LinkedListNode
		end    *LinkedListNode
		lenght int
	}
	type args struct {
		value *RBTreeNode
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
			list := &LinkedList{
				start:  tt.fields.start,
				end:    tt.fields.end,
				lenght: tt.fields.lenght,
			}
			list.add(tt.args.value)
		})
	}
}

func TestLinkedList_pop(t *testing.T) {
	type fields struct {
		start  *LinkedListNode
		end    *LinkedListNode
		lenght int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *LinkedListNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkedList{
				start:  tt.fields.start,
				end:    tt.fields.end,
				lenght: tt.fields.lenght,
			}
			got, err := list.pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_popIndex(t *testing.T) {
	type fields struct {
		start  *LinkedListNode
		end    *LinkedListNode
		lenght int
	}
	type args struct {
		popIndex int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *LinkedListNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkedList{
				start:  tt.fields.start,
				end:    tt.fields.end,
				lenght: tt.fields.lenght,
			}
			got, err := list.popIndex(tt.args.popIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.popIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.popIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_insert(t *testing.T) {
	type fields struct {
		start  *LinkedListNode
		end    *LinkedListNode
		lenght int
	}
	type args struct {
		insertIndex int
		value       *RBTreeNode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkedList{
				start:  tt.fields.start,
				end:    tt.fields.end,
				lenght: tt.fields.lenght,
			}
			if err := list.insert(tt.args.insertIndex, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
