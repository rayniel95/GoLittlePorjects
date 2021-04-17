package closingislands

import "testing"

func Test_cerca(t *testing.T) {
	type args struct {
		table [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1", args{
				[][]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 1, 0, 0, 2, 2, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 3, 3, 3},
					{0, 0, 4, 0, 0, 0, 0, 3, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			20,
		},
		{
			"test2", args{
				[][]int{
					{0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0},
					{0, 0, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0},
				},
			},
			1,
		},
		{
			"test3", args{
				[][]int{
					{0, 0, 0, 0, 0, 0},
					{0, 1, 1, 0, 2, 2},
					{0, 1, 0, 0, 0, 2},
					{0, 1, 1, 0, 2, 2},
					{0, 0, 0, 0, 0, 0},
				},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cerca(tt.args.table); got != tt.want {
				t.Errorf("cerca() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_allIn(t *testing.T) {
// 	type args struct {
// 		table [][]int
// 		x1    int
// 		x2    int
// 		y1    int
// 		y2    int
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
// 			if got := allIn(tt.args.table, tt.args.x1, tt.args.x2, tt.args.y1, tt.args.y2); got != tt.want {
// 				t.Errorf("allIn() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
