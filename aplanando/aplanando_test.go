package aplanando

import "testing"

func Test_aplana(t *testing.T) {
	type args struct {
		stacks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{stacks: []int{0, 1, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aplana(tt.args.stacks); got != tt.want {
				t.Errorf("aplana() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aplanando(t *testing.T) {
	type args struct {
		stacks []int
		movs   []mov
		min    int
		best   *[]mov
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aplanando(tt.args.stacks, tt.args.movs, tt.args.min, tt.args.best)
		})
	}
}

func Test_isAplaned(t *testing.T) {
	type args struct {
		stacks []int
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
			if got := isAplaned(tt.args.stacks); got != tt.want {
				t.Errorf("isAplaned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arrayMin(t *testing.T) {
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
			if got := arrayMin(tt.args.array); got != tt.want {
				t.Errorf("arrayMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
