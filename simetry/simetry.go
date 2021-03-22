package simetry

type treeNode struct {
	left  *treeNode
	right *treeNode
}

func (self *treeNode) countNodes() int {
	sum := 1
	if (*self).left != nil {
		sum += (*self).left.countNodes()
	}
	if (*self).right != nil {
		sum += (*self).right.countNodes()
	}
	return sum
}

func doubleEntreOrden(left *treeNode, right *treeNode) int {
	sum := 0
	if (*left).left != nil && (*right).right != nil {
		sum += doubleEntreOrden((*left).left, (*right).right)
	}
	if (*left).right != nil && (*right).left != nil {
		sum += doubleEntreOrden((*left).right, (*right).left)
	}
	if (*left).left != nil && (*right).right == nil {
		sum += (*left).left.countNodes()
	}
	if (*left).left == nil && (*right).right != nil {
		sum += (*right).right.countNodes()
	}
	if (*left).right != nil && (*right).left == nil {
		sum += (*left).right.countNodes()
	}
	if (*left).right == nil && (*right).left != nil {
		sum += (*right).left.countNodes()

	}
	return sum
}

func doubleEntreOrden2(left *treeNode, right *treeNode) int {
	if left == nil && right == nil {
		return 0
	}
	if left != nil && right == nil {
		return left.countNodes()
	}
	if left == nil && right != nil {
		return right.countNodes()
	}
	sum := 0
	sum += doubleEntreOrden2((*left).left, (*right).right)
	sum += doubleEntreOrden2((*left).right, (*right).left)
	return sum
}

func (node *treeNode) isLeaf() bool {
	return (*node).right == nil && (*node).left == nil
}
