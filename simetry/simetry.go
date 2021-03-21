package simetry

type treeNode struct {
	left  *treeNode
	right *treeNode
}

func (self *treeNode) countNodes() int {
	if (*self).right == nil && (*self).left == nil {
		return 1
	}
	sum := 0
	if (*self).left != nil {
		sum += (*self).left.countNodes()
	}
	if (*self).right != nil {
		sum += (*self).right.countNodes()
	}
	return sum
}
