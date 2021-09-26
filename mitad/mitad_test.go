package mitad

import "testing"

func Test_mitad(t *testing.T) {
	subtree:= &Node{
		taken: true,
		value: 4,
		sons: []*Node{
			&Node{
				value: 4,
				sons: []*Node{},
			},
			&Node{
				value: 2,
				sons: []*Node{
					&Node{
						value: 1,
						sons: []*Node{},
					},
					&Node{
						value: 7,
						sons: []*Node{},
					},
					&Node{
						value: 9,
						sons: []*Node{},
					},
				},
			},
		},
	}
	
	tree:= &Node{
		value: 4,
		sons: []*Node{
			&Node{
				value: 4,
				sons: []*Node{},
			},
			&Node{
				value: 2,
				sons: []*Node{
					&Node{
						value: 1,
						sons: []*Node{},
					},
					subtree,
					&Node{
						value: 7,
						sons: []*Node{},
					},
					&Node{
						value: 9,
						sons: []*Node{},
					},
				},
			},
		},
	}

	type args struct {
		node  *Node
		other *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				tree,
				subtree,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mitad(tt.args.node, tt.args.other); got != tt.want {
				t.Errorf("mitad() = %v, want %v", got, tt.want)
			}
		})
	}
}
