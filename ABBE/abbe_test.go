package ABBE

import (
	"reflect"
	"testing"
)

func Test_abbeNode_insert(t *testing.T) {
	type fields struct {
		left  *abbeNode
		right *abbeNode
		value IComparable
	}
	type args struct {
		treeNode *abbeNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &abbeNode{
				left:  tt.fields.left,
				right: tt.fields.right,
				value: tt.fields.value,
			}
			node.insert(tt.args.treeNode)
		})
	}
}

func Test_abbeNode_frecuencyPerValue(t *testing.T) {
	type fields struct {
		left  *abbeNode
		right *abbeNode
		value IComparable
	}
	tests := []struct {
		name   string
		fields fields
		want   map[IComparable]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &abbeNode{
				left:  tt.fields.left,
				right: tt.fields.right,
				value: tt.fields.value,
			}
			if got := node.frecuencyPerValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("abbeNode.frecuencyPerValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abbeNode_values(t *testing.T) {
	type fields struct {
		left  *abbeNode
		right *abbeNode
		value IComparable
	}
	tests := []struct {
		name   string
		fields fields
		want   []IComparable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &abbeNode{
				left:  tt.fields.left,
				right: tt.fields.right,
				value: tt.fields.value,
			}
			if got := node.values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("abbeNode.values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myInt_compare(t *testing.T) {
	type args struct {
		other interface{}
	}
	tests := []struct {
		name string
		inte *myInt
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.inte.compare(tt.args.other); got != tt.want {
				t.Errorf("myInt.compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_tester1(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{
// 			"test1",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tester1()
// 		})
// 	}
// }
