package treemincost

import "testing"

// func Test_noTaken(t *testing.T) {
// 	type args struct {
// 		array []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := noTaken(tt.args.array); got != tt.want {
// 				t.Errorf("noTaken() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_selector(t *testing.T) {
// 	type args struct {
// 		array []int
// 		taken []int
// 		index int
// 		taked int
// 		max   int
// 		level int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			selector(tt.args.array, tt.args.taken, tt.args.index, tt.args.taked, tt.args.max, tt.args.level)
// 		})
// 	}
// }

// func Test_sel(t *testing.T) {
// 	type args struct {
// 		array []int
// 		taken []int
// 		level int
// 		best  *int
// 		taked int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sel(tt.args.array, tt.args.taken, tt.args.level, tt.args.best, tt.args.taked)
// 		})
// 	}
// }

// func Test_calculus(t *testing.T) {
// 	type args struct {
// 		array []int
// 		taken []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := calculus(tt.args.array, tt.args.taken); got != tt.want {
// 				t.Errorf("calculus() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_constructTree(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1", args{[]int{2, -1, 3}}, 3,
		},
		{
			"test2", args{[]int{-1, 3, 4, 5, 6, 7}}, 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructTree(tt.args.array); got != tt.want {
				t.Errorf("constructTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
