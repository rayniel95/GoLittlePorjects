package pathnumber

import (
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

// func Test_topologicalSort(t *testing.T) {
// 	type args struct {
// 		graph      []*LinkedList
// 		sort       *LinkedList
// 		actualNode *Vertex
// 		time       *int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			topologicalSort(tt.args.graph, tt.args.sort, tt.args.actualNode, tt.args.time)
// 		})
// 	}
// }

// func Test_trasposeG(t *testing.T) {
// 	type args struct {
// 		graph   []*LinkedList
// 		vertexs []*Vertex
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
// 			if got := trasposeG(tt.args.graph, tt.args.vertexs); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("trasposeG() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_pathCount(t *testing.T) {
	type args struct {
		graph   []*LinkedList
		a       *Vertex
		b       *Vertex
		vertexs []*Vertex
	}

	vertexs := make([]*Vertex, 14)
	for count := 0; count < 14; count++ {
		vertexs[count] = &Vertex{
			index: count, start: 0, end: 0, visited: false, paths: 0,
		}
	}
	graph := make([]*LinkedList, 14)
	for count := 0; count < 14; count++ {
		graph[count] = &LinkedList{}
	}

	graph[0].add(vertexs[4])
	graph[0].add(vertexs[5])
	graph[0].add(vertexs[11])

	graph[1].add(vertexs[2])
	graph[1].add(vertexs[4])
	graph[1].add(vertexs[8])

	graph[2].add(vertexs[5])
	graph[2].add(vertexs[6])
	graph[2].add(vertexs[9])

	graph[3].add(vertexs[2])
	graph[3].add(vertexs[6])
	graph[3].add(vertexs[13])

	graph[4].add(vertexs[7])

	graph[5].add(vertexs[8])
	graph[5].add(vertexs[12])

	graph[6].add(vertexs[5])

	graph[8].add(vertexs[7])

	graph[9].add(vertexs[10])
	graph[9].add(vertexs[11])

	graph[10].add(vertexs[13])

	graph[12].add(vertexs[9])

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1", args{
				graph:   graph,
				a:       vertexs[3],
				b:       vertexs[9],
				vertexs: vertexs,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathCount(tt.args.graph, tt.args.a, tt.args.b, tt.args.vertexs); got != tt.want {
				t.Errorf("pathCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
