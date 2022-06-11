package neveras

import (
	"testing"
)

func Test_toMerge(t *testing.T) {
	type args struct {
		nevs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nevs: [][]int{{4}, {1}, {6}}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toMerge(tt.args.nevs); got != tt.want {
				t.Errorf("toMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeNevs(t *testing.T) {
	type args struct {
		nevs     [][]int
		cost     int
		bestCost *int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeNevs(tt.args.nevs, tt.args.cost, tt.args.bestCost)
		})
	}
}

// func Test_merge(t *testing.T) {
// 	type args struct {
// 		nevs  [][]int
// 		index int
// 		dir   int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  [][]int
// 		want1 int
// 	}{
// 		{
// 			name: "test1",
// 			args: args{
// 				nevs:  [][]int{{4}, {1}, {6}},
// 				index: 0,
// 				dir:   1,
// 			},
// 			want: [][]int{{4, 1}, {6}},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := merge(tt.args.nevs, tt.args.index, tt.args.dir)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("merge() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("merge() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

func Test_sum(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.array); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
