package permutations

import (
	"testing"
)

func Test_permuting(t *testing.T) {
	type args struct {
		original []int
		selected []bool
		permuted []int
		call     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				[]int{1, 2, 3, 4},
				make([]bool, 4),
				make([]int, 4),
				0,
			},
			24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuting(tt.args.original, tt.args.selected, tt.args.permuted, tt.args.call); got != tt.want {
				t.Errorf("permuting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permuting2(t *testing.T) {
	type args struct {
		original []int
		selected []bool
		permuted []int
		call     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				[]int{1, 2, 3, 4},
				make([]bool, 4),
				make([]int, 3),
				0,
			},
			24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuting2(tt.args.original, tt.args.selected, tt.args.permuted, tt.args.call); got != tt.want {
				t.Errorf("permuting2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permutingWithRep(t *testing.T) {
	type args struct {
		original []int
		permuted []int
		call     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				[]int{1, 2, 3, 4},
				make([]int, 3),
				0,
			},
			64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutingWithRep(tt.args.original, tt.args.permuted, tt.args.call); got != tt.want {
				t.Errorf("permutingWithRep() = %v, want %v", got, tt.want)
			}
		})
	}
}
