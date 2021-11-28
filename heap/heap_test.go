package heap

import (
	"fmt"
	"math/rand"
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

func TestHeapNode_heapifyDownRootTreeValue(t *testing.T) {
	nodes := createNodes(1001)

	addLeftSon(nodes[1000], nodes[2])
	addRightSon(nodes[1000], nodes[1])

	addLeftSon(nodes[2], nodes[3])
	addRightSon(nodes[2], nodes[4])

	addLeftSon(nodes[1], nodes[5])
	addRightSon(nodes[1], nodes[6])

	nodes[1000].heapifyDown()

	node2 := createNodes(1001)

	addLeftSon(node2[1], node2[2])
	addRightSon(node2[1], node2[5])

	addLeftSon(node2[2], node2[3])
	addRightSon(node2[2], node2[4])

	addLeftSon(node2[5], node2[1000])
	addRightSon(node2[5], node2[6])

	if !reflect.DeepEqual(nodes[1000], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func TestHeapNode_heapifyDownRootTreeValueNotMoved(t *testing.T) {
	nodes := createNodes(32)

	addLeftSon(nodes[1], nodes[2])
	addRightSon(nodes[1], nodes[3])

	addLeftSon(nodes[2], nodes[4])
	addRightSon(nodes[2], nodes[5])

	addLeftSon(nodes[3], nodes[6])
	addRightSon(nodes[3], nodes[7])

	addLeftSon(nodes[4], nodes[8])
	addRightSon(nodes[4], nodes[9])

	addLeftSon(nodes[5], nodes[11])
	addRightSon(nodes[5], nodes[12])

	addLeftSon(nodes[6], nodes[12])
	addRightSon(nodes[6], nodes[13])

	addLeftSon(nodes[7], nodes[14])
	addRightSon(nodes[7], nodes[15])

	addLeftSon(nodes[8], nodes[16])
	addRightSon(nodes[8], nodes[17])

	addLeftSon(nodes[9], nodes[18])
	addRightSon(nodes[9], nodes[19])

	addLeftSon(nodes[10], nodes[20])
	addRightSon(nodes[10], nodes[21])

	addLeftSon(nodes[11], nodes[22])
	addRightSon(nodes[11], nodes[23])

	addLeftSon(nodes[12], nodes[24])
	addRightSon(nodes[12], nodes[25])

	addLeftSon(nodes[13], nodes[26])
	addRightSon(nodes[13], nodes[27])

	addLeftSon(nodes[14], nodes[28])
	addRightSon(nodes[14], nodes[29])

	addLeftSon(nodes[15], nodes[30])
	addRightSon(nodes[15], nodes[31])

	nodes[1].heapifyDown()

	node2 := createNodes(32)

	addLeftSon(node2[1], node2[2])
	addRightSon(node2[1], node2[3])

	addLeftSon(node2[2], node2[4])
	addRightSon(node2[2], node2[5])

	addLeftSon(node2[3], node2[6])
	addRightSon(node2[3], node2[7])

	addLeftSon(node2[4], node2[8])
	addRightSon(node2[4], node2[9])

	addLeftSon(node2[5], node2[11])
	addRightSon(node2[5], node2[12])

	addLeftSon(node2[6], node2[12])
	addRightSon(node2[6], node2[13])

	addLeftSon(node2[7], node2[14])
	addRightSon(node2[7], node2[15])

	addLeftSon(node2[8], node2[16])
	addRightSon(node2[8], node2[17])

	addLeftSon(node2[9], node2[18])
	addRightSon(node2[9], node2[19])

	addLeftSon(node2[10], node2[20])
	addRightSon(node2[10], node2[21])

	addLeftSon(node2[11], node2[22])
	addRightSon(node2[11], node2[23])

	addLeftSon(node2[12], node2[24])
	addRightSon(node2[12], node2[25])

	addLeftSon(node2[13], node2[26])
	addRightSon(node2[13], node2[27])

	addLeftSon(node2[14], node2[28])
	addRightSon(node2[14], node2[29])

	addLeftSon(node2[15], node2[30])
	addRightSon(node2[15], node2[31])

	if !reflect.DeepEqual(nodes[1], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func TestHeapNode_heapifyDownMiddleTreeValueNotMoved(t *testing.T) {
	nodes := createNodes(32)

	addLeftSon(nodes[1], nodes[3])
	addRightSon(nodes[1], nodes[2])

	addLeftSon(nodes[2], nodes[6])
	addRightSon(nodes[2], nodes[7])

	addLeftSon(nodes[3], nodes[4])
	addRightSon(nodes[3], nodes[5])

	addLeftSon(nodes[4], nodes[15])
	addRightSon(nodes[4], nodes[14])

	addLeftSon(nodes[5], nodes[13])
	addRightSon(nodes[5], nodes[12])

	addLeftSon(nodes[6], nodes[11])
	addRightSon(nodes[6], nodes[10])

	addLeftSon(nodes[7], nodes[9])
	addRightSon(nodes[7], nodes[8])

	addLeftSon(nodes[8], nodes[30])
	addRightSon(nodes[8], nodes[31])

	addLeftSon(nodes[9], nodes[28])
	addRightSon(nodes[9], nodes[29])

	addLeftSon(nodes[10], nodes[26])
	addRightSon(nodes[10], nodes[27])

	addLeftSon(nodes[11], nodes[24])
	addRightSon(nodes[11], nodes[25])

	addLeftSon(nodes[12], nodes[22])
	addRightSon(nodes[12], nodes[23])

	addLeftSon(nodes[13], nodes[20])
	addRightSon(nodes[13], nodes[21])

	addLeftSon(nodes[14], nodes[18])
	addRightSon(nodes[14], nodes[19])

	addLeftSon(nodes[15], nodes[16])
	addRightSon(nodes[15], nodes[17])

	nodes[3].heapifyDown()

	node2 := createNodes(32)

	addLeftSon(node2[1], node2[3])
	addRightSon(node2[1], node2[2])

	addLeftSon(node2[2], node2[6])
	addRightSon(node2[2], node2[7])

	addLeftSon(node2[3], node2[4])
	addRightSon(node2[3], node2[5])

	addLeftSon(node2[4], node2[15])
	addRightSon(node2[4], node2[14])

	addLeftSon(node2[5], node2[13])
	addRightSon(node2[5], node2[12])

	addLeftSon(node2[6], node2[11])
	addRightSon(node2[6], node2[10])

	addLeftSon(node2[7], node2[9])
	addRightSon(node2[7], node2[8])

	addLeftSon(node2[8], node2[30])
	addRightSon(node2[8], node2[31])

	addLeftSon(node2[9], node2[28])
	addRightSon(node2[9], node2[29])

	addLeftSon(node2[10], node2[26])
	addRightSon(node2[10], node2[27])

	addLeftSon(node2[11], node2[24])
	addRightSon(node2[11], node2[25])

	addLeftSon(node2[12], node2[22])
	addRightSon(node2[12], node2[23])

	addLeftSon(node2[13], node2[20])
	addRightSon(node2[13], node2[21])

	addLeftSon(node2[14], node2[18])
	addRightSon(node2[14], node2[19])

	addLeftSon(node2[15], node2[16])
	addRightSon(node2[15], node2[17])

	if !reflect.DeepEqual(nodes[1], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func TestHeapNode_heapifyDownMiddleTreeValueheapifyDownOneLevels(t *testing.T) {
	nodes := createNodes(32)

	addLeftSon(nodes[1], nodes[11])
	addRightSon(nodes[1], nodes[2])

	addLeftSon(nodes[2], nodes[6])
	addRightSon(nodes[2], nodes[7])

	addLeftSon(nodes[11], nodes[4])
	addRightSon(nodes[11], nodes[5])

	addLeftSon(nodes[4], nodes[15])
	addRightSon(nodes[4], nodes[14])

	addLeftSon(nodes[5], nodes[13])
	addRightSon(nodes[5], nodes[12])

	addLeftSon(nodes[6], nodes[24])
	addRightSon(nodes[6], nodes[25])

	addLeftSon(nodes[7], nodes[26])
	addRightSon(nodes[7], nodes[27])

	addLeftSon(nodes[12], nodes[22])
	addRightSon(nodes[12], nodes[23])

	addLeftSon(nodes[13], nodes[20])
	addRightSon(nodes[13], nodes[21])

	addLeftSon(nodes[14], nodes[18])
	addRightSon(nodes[14], nodes[19])

	addLeftSon(nodes[15], nodes[16])
	addRightSon(nodes[15], nodes[17])

	nodes[11].heapifyDown()

	node2 := createNodes(32)

	addLeftSon(node2[1], node2[4])
	addRightSon(node2[1], node2[2])

	addLeftSon(node2[2], node2[6])
	addRightSon(node2[2], node2[7])

	addLeftSon(node2[4], node2[11])
	addRightSon(node2[4], node2[5])

	addLeftSon(node2[11], node2[15])
	addRightSon(node2[11], node2[14])

	addLeftSon(node2[5], node2[13])
	addRightSon(node2[5], node2[12])

	addLeftSon(node2[6], node2[24])
	addRightSon(node2[6], node2[25])

	addLeftSon(node2[7], node2[26])
	addRightSon(node2[7], node2[27])

	addLeftSon(node2[12], node2[22])
	addRightSon(node2[12], node2[23])

	addLeftSon(node2[13], node2[20])
	addRightSon(node2[13], node2[21])

	addLeftSon(node2[14], node2[18])
	addRightSon(node2[14], node2[19])

	addLeftSon(node2[15], node2[16])
	addRightSon(node2[15], node2[17])

	// result := make([]byte, 0)
	// for index := 0; index < len(md5.Sum([]byte("https://cu.linkedin.com/in/amaury95"))); index++ {
	// 	result = append(result, md5.Sum([]byte("https://cu.linkedin.com/in/amaury95"))[index])
	// }
	// log.Println(binary.BigEndian.Uint64(result))
	if !reflect.DeepEqual(nodes[1], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func TestHeapNode_heapifyDownMiddleTreeValueheapifyDownTwoLevels(t *testing.T) {
	nodes := createNodes(32)

	addLeftSon(nodes[1], nodes[3])
	addRightSon(nodes[1], nodes[12])

	addLeftSon(nodes[3], nodes[4])
	addRightSon(nodes[3], nodes[5])

	addLeftSon(nodes[12], nodes[6])
	addRightSon(nodes[12], nodes[7])

	addLeftSon(nodes[4], nodes[15])
	addRightSon(nodes[4], nodes[14])

	addLeftSon(nodes[5], nodes[13])
	addRightSon(nodes[5], nodes[22])

	addLeftSon(nodes[6], nodes[11])
	addRightSon(nodes[6], nodes[10])

	addLeftSon(nodes[7], nodes[9])
	addRightSon(nodes[7], nodes[8])

	addLeftSon(nodes[11], nodes[24])
	addRightSon(nodes[11], nodes[25])

	addLeftSon(nodes[10], nodes[26])
	addRightSon(nodes[10], nodes[27])

	addLeftSon(nodes[9], nodes[28])
	addRightSon(nodes[9], nodes[29])

	addLeftSon(nodes[8], nodes[30])
	addRightSon(nodes[8], nodes[31])

	nodes[12].heapifyDown()

	node2 := createNodes(32)

	addLeftSon(node2[1], node2[3])
	addRightSon(node2[1], node2[6])

	addLeftSon(node2[3], node2[4])
	addRightSon(node2[3], node2[5])

	addLeftSon(node2[6], node2[10])
	addRightSon(node2[6], node2[7])

	addLeftSon(node2[4], node2[15])
	addRightSon(node2[4], node2[14])

	addLeftSon(node2[5], node2[13])
	addRightSon(node2[5], node2[22])

	addLeftSon(node2[10], node2[11])
	addRightSon(node2[10], node2[12])

	addLeftSon(node2[7], node2[9])
	addRightSon(node2[7], node2[8])

	addLeftSon(node2[11], node2[24])
	addRightSon(node2[11], node2[25])

	addLeftSon(node2[12], node2[26])
	addRightSon(node2[12], node2[27])

	addLeftSon(node2[9], node2[28])
	addRightSon(node2[9], node2[29])

	addLeftSon(node2[8], node2[30])
	addRightSon(node2[8], node2[31])

	if !reflect.DeepEqual(nodes[1], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func TestHeapNode_heapifyUpToRoot(t *testing.T) {
	nodes := createNodes(32)

	addLeftSon(nodes[2], nodes[6])
	addRightSon(nodes[2], nodes[12])

	addLeftSon(nodes[6], nodes[7])
	addRightSon(nodes[6], nodes[8])

	addLeftSon(nodes[12], nodes[15])
	addRightSon(nodes[12], nodes[1])

	nodes[1].heapifyUp()

	node2 := createNodes(32)

	addLeftSon(node2[1], node2[6])
	addRightSon(node2[1], node2[2])

	addLeftSon(node2[6], node2[7])
	addRightSon(node2[6], node2[8])

	addLeftSon(node2[2], node2[15])
	addRightSon(node2[2], node2[12])

	if !reflect.DeepEqual(nodes[2], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func TestHeapNode_heapifyUpToRootOneLevel(t *testing.T) {
	nodes := createNodes(32)

	addLeftSon(nodes[2], nodes[3])
	addRightSon(nodes[2], nodes[1])

	addLeftSon(nodes[1], nodes[6])
	addRightSon(nodes[1], nodes[7])

	addLeftSon(nodes[3], nodes[4])
	addRightSon(nodes[3], nodes[5])

	nodes[1].heapifyUp()

	node2 := createNodes(32)

	addLeftSon(node2[1], node2[3])
	addRightSon(node2[1], node2[2])

	addLeftSon(node2[2], node2[6])
	addRightSon(node2[2], node2[7])

	addLeftSon(node2[3], node2[4])
	addRightSon(node2[3], node2[5])

	if !reflect.DeepEqual(nodes[2], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func TestHeapNode_heapifyUpToMiddleTwoLevel(t *testing.T) {
	nodes := createNodes(40)

	addLeftSon(nodes[1], nodes[2])
	addRightSon(nodes[1], nodes[19])

	addLeftSon(nodes[2], nodes[4])
	addRightSon(nodes[2], nodes[5])

	addLeftSon(nodes[4], nodes[11])
	addRightSon(nodes[4], nodes[10])

	addLeftSon(nodes[5], nodes[9])
	addRightSon(nodes[5], nodes[8])

	addLeftSon(nodes[19], nodes[6])
	addRightSon(nodes[19], nodes[20])

	addLeftSon(nodes[6], nodes[18])
	addRightSon(nodes[6], nodes[17])

	addLeftSon(nodes[18], nodes[30])
	addRightSon(nodes[18], nodes[31])

	addLeftSon(nodes[17], nodes[32])
	addRightSon(nodes[17], nodes[33])

	addLeftSon(nodes[20], nodes[7])
	addRightSon(nodes[20], nodes[12])

	addLeftSon(nodes[7], nodes[34])
	addRightSon(nodes[7], nodes[35])

	addLeftSon(nodes[12], nodes[36])
	addRightSon(nodes[12], nodes[37])

	nodes[12].heapifyUp()

	node2 := createNodes(40)

	addLeftSon(node2[1], node2[2])
	addRightSon(node2[1], node2[12])

	addLeftSon(node2[2], node2[4])
	addRightSon(node2[2], node2[5])

	addLeftSon(node2[4], node2[11])
	addRightSon(node2[4], node2[10])

	addLeftSon(node2[5], node2[9])
	addRightSon(node2[5], node2[8])

	addLeftSon(node2[19], node2[7])
	addRightSon(node2[19], node2[20])

	addLeftSon(node2[6], node2[18])
	addRightSon(node2[6], node2[17])

	addLeftSon(node2[18], node2[30])
	addRightSon(node2[18], node2[31])

	addLeftSon(node2[17], node2[32])
	addRightSon(node2[17], node2[33])

	addLeftSon(node2[20], node2[36])
	addRightSon(node2[20], node2[37])

	addLeftSon(node2[7], node2[34])
	addRightSon(node2[7], node2[35])

	addLeftSon(node2[12], node2[6])
	addRightSon(node2[12], node2[19])

	if !reflect.DeepEqual(nodes[1], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func TestHeapNode_heapifyUpToLeavesThreeLevels(t *testing.T) {
	nodes := createNodes(40)

	addLeftSon(nodes[1], nodes[2])
	addRightSon(nodes[1], nodes[12])

	addLeftSon(nodes[2], nodes[4])
	addRightSon(nodes[2], nodes[5])

	addLeftSon(nodes[4], nodes[11])
	addRightSon(nodes[4], nodes[10])

	addLeftSon(nodes[5], nodes[9])
	addRightSon(nodes[5], nodes[8])

	addLeftSon(nodes[19], nodes[7])
	addRightSon(nodes[19], nodes[20])

	addLeftSon(nodes[6], nodes[18])
	addRightSon(nodes[6], nodes[17])

	addLeftSon(nodes[18], nodes[30])
	addRightSon(nodes[18], nodes[31])

	addLeftSon(nodes[17], nodes[32])
	addRightSon(nodes[17], nodes[33])

	addLeftSon(nodes[20], nodes[36])
	addRightSon(nodes[20], nodes[3])

	addLeftSon(nodes[7], nodes[34])
	addRightSon(nodes[7], nodes[35])

	addLeftSon(nodes[12], nodes[6])
	addRightSon(nodes[12], nodes[19])

	nodes[3].heapifyUp()

	node2 := createNodes(40)

	addLeftSon(node2[1], node2[2])
	addRightSon(node2[1], node2[3])

	addLeftSon(node2[2], node2[4])
	addRightSon(node2[2], node2[5])

	addLeftSon(node2[4], node2[11])
	addRightSon(node2[4], node2[10])

	addLeftSon(node2[5], node2[9])
	addRightSon(node2[5], node2[8])

	addLeftSon(node2[19], node2[36])
	addRightSon(node2[19], node2[20])

	addLeftSon(node2[6], node2[18])
	addRightSon(node2[6], node2[17])

	addLeftSon(node2[18], node2[30])
	addRightSon(node2[18], node2[31])

	addLeftSon(node2[17], node2[32])
	addRightSon(node2[17], node2[33])

	addLeftSon(node2[12], node2[7])
	addRightSon(node2[12], node2[19])

	addLeftSon(node2[7], node2[34])
	addRightSon(node2[7], node2[35])

	addLeftSon(node2[3], node2[6])
	addRightSon(node2[3], node2[12])

	if !reflect.DeepEqual(nodes[1], node2[1]) {
		(*t).Errorf("heaps are not equals")
	}
}

func (tree *HeapNode) print() {
	fmt.Println((*((*tree).cell)).priority)
	if (*tree).left != nil {
		(*tree).left.print()
	}
	if (*tree).right != nil {
		(*tree).right.print()
	}
}

// TODO: Add fuzzy tests

func TestHeapTree_fuzzyAddDeleteManyNodes(t *testing.T) {
	heap := &HeapTree{}
	for times := 0; times < 10000000; times++ {
		new := rand.Uint32()
		heap.Add(new, new)
	}

	_, cell := heap.DeleteMin()
	min := (*cell).priority
	for times := 1; times < 10000000; times++ {
		err, new_min_cell := heap.DeleteMin()
		if err != nil {
			t.Errorf("not all values unpacked")
		}
		new_min := (*new_min_cell).priority
		if new_min < min {
			t.Error("last poped value is smallest than previous values")
		}
		min = new_min
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

// func TestHeapTree_add(t *testing.T) {
// 	type fields struct {
// 		start        *LinkedNode
// 		end          *LinkedNode
// 		parentOfLast *LinkedNode
// 	}
// 	type args struct {
// 		value    *interface{}
// 		priority uint32
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
// 			tree := &HeapTree{
// 				start:        tt.fields.start,
// 				end:          tt.fields.end,
// 				parentOfLast: tt.fields.parentOfLast,
// 			}
// 			tree.Add(tt.args.value, tt.args.priority)
// 		})
// 	}
// }

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
