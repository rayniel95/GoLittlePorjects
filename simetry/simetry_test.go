package simetry

import (
	"testing"
)

func Test_doubleEntreOrden2(t *testing.T) {
	root := &treeNode{
		left: &treeNode{
			left: &treeNode{
				left:  &treeNode{},
				right: &treeNode{},
			},
		},
		right: &treeNode{
			right: &treeNode{
				left: &treeNode{
					right: &treeNode{},
				},
			},
		},
	}
	want := 2
	got := doubleEntreOrden2((*root).left, (*root).right)

	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func Test_doubleEntreOrden(t *testing.T) {
	root := &treeNode{
		left: &treeNode{
			left: &treeNode{
				left:  &treeNode{},
				right: &treeNode{},
			},
		},
		right: &treeNode{
			right: &treeNode{
				left: &treeNode{
					right: &treeNode{},
				},
			},
		},
	}
	want := 2
	got := doubleEntreOrden((*root).left, (*root).right)

	if got != want {
		t.Errorf("Hello() = %d, want %d", got, want)
	}
}
