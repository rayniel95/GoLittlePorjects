package linkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLinkedList_add(t *testing.T) {
	type fields struct {
		start  *LinkedListNode
		end    *LinkedListNode
		lenght int
	}
	type args struct {
		value int
	}

	list := &LinkedList{}

	tests := []struct {
		name string
		args args
		want struct {
			list   []int
			start  int
			end    int
			lenght int
		}
	}{
		{
			"test1", args{2}, struct {
				list   []int
				start  int
				end    int
				lenght int
			}{[]int{2}, 2, 2, 1},
		},
		{
			"test2", args{3}, struct {
				list   []int
				start  int
				end    int
				lenght int
			}{[]int{2, 3}, 2, 3, 2},
		},
		{
			"test3", args{5}, struct {
				list   []int
				start  int
				end    int
				lenght int
			}{[]int{2, 3, 5}, 2, 5, 3},
		},
		{
			"test4", args{20}, struct {
				list   []int
				start  int
				end    int
				lenght int
			}{[]int{2, 3, 5, 20}, 2, 20, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list.add(tt.args.value)

			result := linkedListToSlice(list)
			if len(result) != tt.want.lenght ||
				(*((*list).start)).value != tt.want.start ||
				(*((*list).end)).value != tt.want.end ||
				!equalSlices(result, tt.want.list) {
				t.Errorf("wrong")
			}
		})
	}
}

func TestLinkedList_pop(t *testing.T) {
	endNode := &LinkedListNode{value: 100}

	list := &LinkedList{
		start: &LinkedListNode{
			value: 3,
			next: &LinkedListNode{
				value: 6,
				next: &LinkedListNode{
					value: 9,
					next: &LinkedListNode{
						value: 5,
						next:  endNode,
					},
				},
			},
		},
		end:    endNode,
		lenght: 5,
	}
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{
			"test1",
			100, false,
		},
		{
			"test2",
			5, false,
		},
		{
			"test3",
			9, false,
		},
		{
			"test4",
			6, false,
		},
		{
			"test5",
			3, false,
		},
		{
			"test6",
			0, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := list.pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !reflect.DeepEqual((*got).value, tt.want) {
				t.Errorf("LinkedList.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_popIndex(t *testing.T) {
	endNode := &LinkedListNode{value: 100}

	list := &LinkedList{
		start: &LinkedListNode{
			value: 3,
			next: &LinkedListNode{
				value: 6,
				next: &LinkedListNode{
					value: 9,
					next: &LinkedListNode{
						value: 5,
						next:  endNode,
					},
				},
			},
		},
		end:    endNode,
		lenght: 5,
	}

	type args struct {
		popIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"test1", args{1}, 6, false,
		},
		{
			"test2", args{2}, 5, false,
		},
		{
			"test3", args{1}, 9, false,
		},
		{
			"test4", args{1}, 100, false,
		},
		{
			"test5", args{0}, 3, false,
		},
		{
			"test6", args{0}, 100, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := list.popIndex(tt.args.popIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.popIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !reflect.DeepEqual((*got).value, tt.want) {
				t.Errorf("LinkedList.popIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_insert(t *testing.T) {
	endNode := &LinkedListNode{value: 100}

	list := &LinkedList{
		start: &LinkedListNode{
			value: 3,
			next: &LinkedListNode{
				value: 6,
				next: &LinkedListNode{
					value: 9,
					next: &LinkedListNode{
						value: 5,
						next:  endNode,
					},
				},
			},
		},
		end:    endNode,
		lenght: 5,
	}
	type args struct {
		insertIndex int
		value       int
	}
	tests := []struct {
		name    string
		args    args
		list    []int
		wantErr bool
	}{
		{
			"test1", args{0, 4}, []int{4, 3, 6, 9, 5, 100}, false,
		},
		{
			"test2", args{0, 5}, []int{5, 4, 3, 6, 9, 5, 100}, false,
		},
		{
			"test3", args{1, 7}, []int{5, 7, 4, 3, 6, 9, 5, 100}, false,
		},
		{
			"test4", args{5, 10}, []int{5, 7, 4, 3, 6, 10, 9, 5, 100}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := list.insert(tt.args.insertIndex, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !equalSlices(tt.list, linkedListToSlice(list)) {
				fmt.Println(linkedListToSlice(list))
				t.Errorf("wrong")
			}
		})
	}
}

func linkedListToSlice(list *LinkedList) []int {
	if (*list).lenght == 0 {
		return []int{}
	}
	array := make([]int, 0)
	node := (*list).start
	array = append(array, (*node).value)
	for index := 1; index < (*list).lenght; index++ {
		node = (*node).next
		array = append(array, (*node).value)
	}
	return array
}

func equalSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for index := 0; index < len(slice1); index++ {
		if slice1[index] != slice2[index] {
			return false
		}
	}
	return true
}
