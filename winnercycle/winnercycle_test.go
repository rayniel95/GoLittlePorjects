package winnercycle

import (
	"reflect"
	"testing"
)

// func Test_cashReceive(t *testing.T) {
// 	type args struct {
// 		matrix   [][]float32
// 		cycle    []int
// 		currency int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want float32
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := cashReceive(tt.args.matrix, tt.args.cycle, tt.args.currency); got != tt.want {
// 				t.Errorf("cashReceive() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_findCycle(t *testing.T) {
// 	type args struct {
// 		matrix      [][]float32
// 		taken       []bool
// 		actualCycle []int
// 		best        []int
// 		currency    int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			findCycle(tt.args.matrix, tt.args.taken, tt.args.actualCycle, tt.args.best, tt.args.currency)
// 		})
// 	}
// }

func Test_winnerCycle(t *testing.T) {
	type args struct {
		matrix   [][]float32
		currency int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{
				[][]float32{
					{1.00, 0.50, 0.90, 10.25},
					{1.25, 1, 0.90, 17.0},
					{1.03, 1.50, 1.00, 15.10},
					{0.01, 0.04, 0.05, 1.00},
				},
				1,
			},
			[]int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerCycle(tt.args.matrix, tt.args.currency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("winnerCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
