package g2fromG

import (
	"reflect"
	"testing"
)

// func TestLinkedList_add(t *testing.T) {
// 	type fields struct {
// 		start  *LinkedListNode
// 		end    *LinkedListNode
// 		lenght int
// 	}
// 	type args struct {
// 		value *Vertex
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
// 			list := &LinkedList{
// 				start:  tt.fields.start,
// 				end:    tt.fields.end,
// 				lenght: tt.fields.lenght,
// 			}
// 			list.add(tt.args.value)
// 		})
// 	}
// }

// func TestLinkedList_pop(t *testing.T) {
// 	type fields struct {
// 		start  *LinkedListNode
// 		end    *LinkedListNode
// 		lenght int
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    *LinkedListNode
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			list := &LinkedList{
// 				start:  tt.fields.start,
// 				end:    tt.fields.end,
// 				lenght: tt.fields.lenght,
// 			}
// 			got, err := list.pop()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("LinkedList.pop() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("LinkedList.pop() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestLinkedList_popIndex(t *testing.T) {
// 	type fields struct {
// 		start  *LinkedListNode
// 		end    *LinkedListNode
// 		lenght int
// 	}
// 	type args struct {
// 		popIndex int
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *LinkedListNode
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			list := &LinkedList{
// 				start:  tt.fields.start,
// 				end:    tt.fields.end,
// 				lenght: tt.fields.lenght,
// 			}
// 			got, err := list.popIndex(tt.args.popIndex)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("LinkedList.popIndex() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("LinkedList.popIndex() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestLinkedList_insert(t *testing.T) {
// 	type fields struct {
// 		start  *LinkedListNode
// 		end    *LinkedListNode
// 		lenght int
// 	}
// 	type args struct {
// 		insertIndex int
// 		value       *Vertex
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			list := &LinkedList{
// 				start:  tt.fields.start,
// 				end:    tt.fields.end,
// 				lenght: tt.fields.lenght,
// 			}
// 			if err := list.insert(tt.args.insertIndex, tt.args.value); (err != nil) != tt.wantErr {
// 				t.Errorf("LinkedList.insert() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestLinkedList_addList(t *testing.T) {
// 	type fields struct {
// 		start  *LinkedListNode
// 		end    *LinkedListNode
// 		lenght int
// 	}
// 	type args struct {
// 		other *LinkedList
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
// 			list := &LinkedList{
// 				start:  tt.fields.start,
// 				end:    tt.fields.end,
// 				lenght: tt.fields.lenght,
// 			}
// 			list.addList(tt.args.other)
// 		})
// 	}
// }

// func Test_foundIncidenceList(t *testing.T) {
// 	type args struct {
// 		graph  [][]bool
// 		vertex []*Vertex
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []*LinkedList
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := foundIncidenceList(tt.args.graph, tt.args.vertex); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("foundIncidenceList() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_secondLevelIncidence(t *testing.T) {
// 	type args struct {
// 		incidenceList []*LinkedList
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []*LinkedList
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := secondLevelIncidence(tt.args.incidenceList); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("secondLevelIncidence() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestDFS(t *testing.T) {
// 	type args struct {
// 		list          *LinkedList
// 		index         int
// 		g2            [][]bool
// 		incidenceList []*LinkedList
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			DFS(tt.args.list, tt.args.index, tt.args.g2, tt.args.incidenceList)
// 		})
// 	}
// }

// func Test_buildG2(t *testing.T) {
// 	type args struct {
// 		secondLevel   []*LinkedList
// 		vertex        []*Vertex
// 		incidenceList []*LinkedList
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want [][]bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := buildG2(tt.args.secondLevel, tt.args.vertex, tt.args.incidenceList); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("buildG2() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_g2(t *testing.T) {
	type args struct {
		graph [][]bool
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{
			name: "test1", args: args{
				graph: [][]bool{
					[]bool{false, false, false, false, false, false},
					[]bool{true, false, false, false, false, false},
					[]bool{false, false, false, false, false, true},
					[]bool{false, false, false, false, false, false},
					[]bool{true, true, true, true, false, false},
					[]bool{false, false, false, false, true, false},
				},
			},
			// NOTE - this testcase is not correct
			want: [][]bool{
				[]bool{false, false, false, false, false},
				[]bool{true, false, false, false, false},
				[]bool{false, false, false, true, true},
				[]bool{false, false, false, false, false},
				[]bool{true, true, true, false, true},
				[]bool{true, true, true, true, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g2(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("g2() = %v, want %v", got, tt.want)
			}
		})
	}
}
