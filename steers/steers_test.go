package steers

import "testing"

func Test_steers(t *testing.T) {
	type args struct {
		sticks int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{sticks: 2}, want: 3},
		{name: "test2", args: args{sticks: 1}, want: 2},
		{name: "test3", args: args{sticks: 3}, want: 5},
		{name: "test4", args: args{sticks: 4}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := steers(tt.args.sticks); got != tt.want {
				t.Errorf("steers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_steersA(t *testing.T) {
// 	type args struct {
// 		sticks       int
// 		vertical     bool
// 		numberSteers *int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			steersA(tt.args.sticks, tt.args.vertical, tt.args.numberSteers)
// 		})
// 	}
// }
