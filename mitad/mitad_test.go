package mitad

import "testing"

func Test_mitad(t *testing.T) {
	type args struct {
		node  *Node
		other *Node
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
			if got := mitad(tt.args.node, tt.args.other); got != tt.want {
				t.Errorf("mitad() = %v, want %v", got, tt.want)
			}
		})
	}
}
