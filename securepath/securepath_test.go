package securepath

import (
	"testing"
)

// func Test_warning(t *testing.T) {
// 	type args struct {
// 		table [8][8]int
// 		cell  [2]int
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
// 			if got := warning(tt.args.table, tt.args.cell); got != tt.want {
// 				t.Errorf("warning() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_bestPath(t *testing.T) {
// 	type args struct {
// 		table  [8][8]int
// 		cell   [2]int
// 		number int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{
// 			"test1",
// 			args{
// 				table: [8][8]int{
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 					{0, 0, -1, 0, -1, 0, 0, 0},
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 					{0, 0, -1, 0, 0, 0, 0, 0},
// 					{-1, 0, 0, 0, 0, -1, 0, -1},
// 					{0, 0, 0, 0, 0, 0, 0, 0},
// 					{0, 0, -1, 0, 0, 0, 0, 0},
// 					{0, 0, 0, 0, -1, 0, 0, 0},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			bestPath(tt.args.table, tt.args.cell, tt.args.number)
// 		})
// 	}
// }

// func Test_warning(t *testing.T) {
// 	type args struct {
// 		table [8][8]int
// 		cell  [2]int
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
// 			if got := warning(tt.args.table, tt.args.cell); got != tt.want {
// 				t.Errorf("warning() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_path(t *testing.T) {
	type args struct {
		table [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				table: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, -1, 0, -1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, -1, 0, 0, 0, 0, 0},
					{-1, 0, 0, 0, 0, -1, 0, -1},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, -1, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, -1, 0, 0, 0},
				},
			},
			9,
		},
		{
			"test2",
			args{
				table: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0},
					{-1, -1, -1, -1, -1, -1, -1, -1},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			0,
		},
		{
			"test1",
			args{
				table: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{-1, -1, -1, -1, -1, -1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, -1},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, -1},
					{0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := path(tt.args.table); got != tt.want {
				t.Errorf("path() = %v, want %v", got, tt.want)
			}
		})
	}
}
