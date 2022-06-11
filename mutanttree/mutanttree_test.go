package mutanttree

import (
	"testing"
)

func Test_main(t *testing.T) {
	tree := &TreeNode{
		left: &TreeNode{
			left:  &TreeNode{},
			right: &TreeNode{},
		},
		right: &TreeNode{
			right: &TreeNode{},
		},
	}

	other := &TreeNode{
		left: &TreeNode{
			right: &TreeNode{},
		},
		right: &TreeNode{
			left:  &TreeNode{},
			right: &TreeNode{},
		},
	}

	want := 1
	got := constructTree(tree, other)

	if got != want {
		t.Errorf("Hello() = %d, want %d", got, want)
	}

	tree = &TreeNode{
		left: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{
					left: &TreeNode{
						left: &TreeNode{},
					},
				},
				right: &TreeNode{},
			},
		},
		right: &TreeNode{
			right: &TreeNode{
				left: &TreeNode{},
				right: &TreeNode{
					right: &TreeNode{
						right: &TreeNode{},
					},
				},
			},
		},
	}

	other = &TreeNode{
		left: &TreeNode{
			left: &TreeNode{
				left: &TreeNode{},
				right: &TreeNode{
					right: &TreeNode{
						left: &TreeNode{},
					},
				},
			},
		},
		right: &TreeNode{
			right: &TreeNode{
				left: &TreeNode{
					right: &TreeNode{
						left: &TreeNode{},
					},
				},
				right: &TreeNode{},
			},
		},
	}
	want = 4
	got = constructTree(tree, other)

	if got != want {
		t.Errorf("Hello() = %d, want %d", got, want)
	}
}
