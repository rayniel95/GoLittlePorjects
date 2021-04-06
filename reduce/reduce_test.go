package reduce

import (
	"reflect"
	"testing"
)

// func Test_isReduced(t *testing.T) {
// 	type args struct {
// 		array []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := isReduced(tt.args.array); got != tt.want {
// 				t.Errorf("isReduced() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_reduce(t *testing.T) {
	type args struct {
		array []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {"test1", args{[]int{1, 1, 2, 5, 8, 4, 4, 5, 5, 1}, 0}, []int{1, 1, 2, 5, 8, 4, 4, 5, 5, 1}},
		// {"test2", args{[]int{2, 3, 3, 3, 8, 5, 0, 0, 0, 0}, 0}, []int{2, 8, 5}},
		{"test3", args{[]int{2, 2, 2, 3, 4, 4, 4, 3, 3, 2, 2, 1}, 0}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduce(tt.args.array, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
