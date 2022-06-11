package rbtree

import (
	"testing"
)

// func Test_initRBTree(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want *RBTree
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := initRBTree(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("initRBTree() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRBTree_leftRotation(t *testing.T) {
// 	type fields struct {
// 		root *RBTreeNode
// 	}
// 	type args struct {
// 		node *RBTreeNode
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tree := &RBTree{
// 				root: tt.fields.root,
// 			}
// 			tree.leftRotation(tt.args.node)
// 		})
// 	}
// }

// func TestRBTree_rightRotation(t *testing.T) {
// 	type fields struct {
// 		root *RBTreeNode
// 	}
// 	type args struct {
// 		node *RBTreeNode
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tree := &RBTree{
// 				root: tt.fields.root,
// 			}
// 			tree.rightRotation(tt.args.node)
// 		})
// 	}
// }

func TestRBTree_insert(t *testing.T) {
	tree := initRBTree()

	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
		size int
	}{
		{
			"test1", args{2}, 1,
		},
		{
			"test2", args{50}, 2,
		},
		{
			"test3", args{100}, 3,
		},
		{
			"test4", args{60}, 4,
		},
		{
			"test5", args{20}, 5,
		},
		{
			"test6", args{1}, 6,
		},
		{
			"test7", args{70}, 7,
		},
		{
			"test8", args{55}, 8,
		},
		{
			"test9", args{56}, 9,
		},
		{
			"test10", args{57}, 10,
		},
		{
			"test11", args{58}, 11,
		},
		{
			"test12", args{59}, 12,
		},
		{
			"test13", args{13}, 13,
		},
		{
			"test14", args{12}, 14,
		},
		{
			"test15", args{11}, 15,
		},
		{
			"test16", args{10}, 16,
		},
		{
			"test17", args{9}, 17,
		},
		{
			"test18", args{8}, 18,
		},
		{
			"test19", args{7}, 19,
		},
		{
			"test20", args{100}, 20,
		},
		{
			"test21", args{99}, 21,
		},
		{
			"test22", args{98}, 22,
		},
		{
			"test23", args{101}, 23,
		},
		{
			"test24", args{102}, 24,
		},
		{
			"test25", args{97}, 25,
		},
		{
			"test26", args{96}, 26,
		},
		{
			"test27", args{103}, 27,
		},
		{
			"test28", args{95}, 28,
		},
		{
			"test29", args{104}, 29,
		},
		{
			"test30", args{105}, 30,
		},
		{
			"test31", args{54}, 31,
		},
		{
			"test32", args{53}, 32,
		},
		{
			"test33", args{56}, 33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree.insert(tt.args.key)
			if tree.size() != tt.size {
				t.Errorf("size() = %v, want %v", tree.size(), tt.size)
			} else if !isABB((*tree).root) {
				t.Errorf("is not an ABB")
			} else if !rootBlack(tree) {
				t.Errorf("root is not black")
			} else if !RBTreeProperty4((*tree).root) {
				t.Errorf("rbtree property 4 not found")
			} else if !RBTreeTestBlackNodes((*tree).root) {
				t.Errorf("black nodes not follow rules")
			} else if !RBTreeTestRedNodes((*tree).root) {
				t.Errorf("red nodes not follow rules")
			}
		})
	}
	// printTree((*tree).root, 0, "root")
}

// func TestRBTree_insertFixup(t *testing.T) {
// 	type fields struct {
// 		root *RBTreeNode
// 	}
// 	type args struct {
// 		node *RBTreeNode
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tree := &RBTree{
// 				root: tt.fields.root,
// 			}
// 			tree.insertFixup(tt.args.node)
// 		})
// 	}
// }

// func Test_min(t *testing.T) {
// 	type args struct {
// 		subtree *RBTreeNode
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *RBTreeNode
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := min(tt.args.subtree); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("min() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRBTree_transplant(t *testing.T) {
// 	type fields struct {
// 		root *RBTreeNode
// 	}
// 	type args struct {
// 		deletedNode *RBTreeNode
// 		node        *RBTreeNode
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tree := &RBTree{
// 				root: tt.fields.root,
// 			}
// 			tree.transplant(tt.args.deletedNode, tt.args.node)
// 		})
// 	}
// }

func TestRBTree_delete(t *testing.T) {
	values := []int{
		2, 50, 100, 60, 20, 1, 70, 55, 56, 57, 58, 59, 13, 12, 11,
		10, 9, 8, 7, 99, 98, 101, 102, 97, 96, 103, 95, 104, 105,
		54, 53,
	}
	tree := initRBTree()
	for _, value := range values {
		tree.insert(value)
	}
	type args struct {
		node *RBTreeNode
	}
	tests := []struct {
		name string
		args args
		size int
	}{
		{
			"test1", args{tree.search(56)}, 30,
		},
		{
			"test2", args{tree.search(56)}, 29,
		},
		{
			"test3", args{tree.search(7)}, 28,
		},
		{
			"test4", args{tree.search(100)}, 27,
		},
		{
			"test5", args{tree.search(103)}, 26,
		},
		{
			"test6", args{tree.search(60)}, 25,
		},
		{
			"test7", args{tree.search(50)}, 24,
		},
		{
			"test8", args{tree.search(13)}, 23,
		},
		{
			"test9", args{tree.search(54)}, 22,
		},
		{
			"test10", args{tree.search(55)}, 21,
		},
		{
			"test11", args{tree.search(9)}, 20,
		},
		{
			"test12", args{tree.search(58)}, 19,
		},
		{
			"test13", args{tree.search(70)}, 18,
		},
		{
			"test14", args{tree.search(98)}, 17,
		},
		{
			"test15", args{tree.search(101)}, 16,
		},
		{
			"test16", args{tree.search(105)}, 15,
		},
		{
			"test18", args{tree.search(59)}, 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree.delete(tt.args.node)
			if tree.size() != tt.size {
				// t.Errorf("%v", tree.search(56))
				t.Errorf("size() = %v, want %v", tree.size(), tt.size)
			} else if !isABB((*tree).root) {
				t.Errorf("is not an ABB")
			} else if !rootBlack(tree) {
				t.Errorf("root is not black")
			} else if !RBTreeProperty4((*tree).root) {
				t.Errorf("rbtree property 4 not found")
			} else if !RBTreeTestBlackNodes((*tree).root) {
				t.Errorf("black nodes not follow rules")
			} else if !RBTreeTestRedNodes((*tree).root) {
				t.Errorf("red nodes not follow rules")
			}
		})
	}
	// printTree((*tree).root, 0, "root")
}

// func TestRBTree_deleteFixup(t *testing.T) {
// 	type fields struct {
// 		root *RBTreeNode
// 	}
// 	type args struct {
// 		node *RBTreeNode
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tree := &RBTree{
// 				root: tt.fields.root,
// 			}
// 			tree.deleteFixup(tt.args.node)
// 		})
// 	}
// }

// func Test_isABB(t *testing.T) {
// 	type args struct {
// 		actualNode *RBTreeNode
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := isABB(tt.args.actualNode); got != tt.want {
// 				t.Errorf("isABB() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRBTreeTestRedNodes(t *testing.T) {
// 	type args struct {
// 		actualNode *RBTreeNode
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := RBTreeTestRedNodes(tt.args.actualNode); got != tt.want {
// 				t.Errorf("RBTreeTestRedNodes() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRBTreeTestBlackNodes(t *testing.T) {
// 	type args struct {
// 		actualNode *RBTreeNode
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := RBTreeTestBlackNodes(tt.args.actualNode); got != tt.want {
// 				t.Errorf("RBTreeTestBlackNodes() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRBTreeProperty4(t *testing.T) {
// 	type args struct {
// 		actualNode *RBTreeNode
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := RBTreeProperty4(tt.args.actualNode); got != tt.want {
// 				t.Errorf("RBTreeProperty4() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
