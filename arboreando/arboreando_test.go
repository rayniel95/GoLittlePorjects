package arboreando

import (
	"testing"
)

func Test_main(t *testing.T) {
	leftTree := &TreeNode{
		value: 5,
		left: &TreeNode{
			value: 8,
			left: &TreeNode{
				value: 13,
			},
		},
		right: &TreeNode{
			value: 9,
			left: &TreeNode{
				value: 2,
			},
			right: &TreeNode{
				value: 6,
			},
		},
	}

	right := &TreeNode{
		value: 2,
		left: &TreeNode{
			value: 1,
		},
		right: &TreeNode{
			value: 4,
		},
	}

	want := &TreeNode{
		value: 7,
		left: &TreeNode{
			value: 11,
			left: &TreeNode{
				value: 16,
			},
		},
		right: &TreeNode{
			value: 15,
			left: &TreeNode{
				value: 5,
			},
			right: &TreeNode{
				value: 12,
			},
		},
	}

	got := arborear(leftTree, right, operaInt)

	if !equal(want, got) {
		t.Errorf("wrong")
	}
}
