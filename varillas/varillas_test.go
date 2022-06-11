package varillas

import (
	"reflect"
	"testing"
)

func Test_rebuild(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				[]int{1, 2, 2, 3, 4, 4, 5, 7},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rebuild(tt.args.array); got != tt.want {
				t.Errorf("rebuild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subsetSum(t *testing.T) {
	type args struct {
		array  []int
		k      int
		assign []int
		index  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetSum(tt.args.array, tt.args.k, tt.args.assign, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subsetEqualSum(t *testing.T) {
	type args struct {
		array  []int
		assign []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetEqualSum(tt.args.array, tt.args.assign, tt.args.k); got != tt.want {
				t.Errorf("subsetEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
