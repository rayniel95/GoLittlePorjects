package mutanttree

func constructTree(node *TreeNode, other *TreeNode) int {
	MaxUint := ^uint(0)
	MaxInt := int(MaxUint >> 1)

	best := &MaxInt

	array := make([]*TreeNode, node.cantity())

	index := 0
	createArray(array, &index, node)
	mutateTree(array, 0, 0, best, other)
	return *best
}

func (node *TreeNode) cantity() int {
	if (*node).left == nil && (*node).right == nil {
		return 1
	}
	sum := 1
	if (*node).left != nil {
		sum += (*node).left.cantity()
	}
	if (*node).right != nil {
		sum += (*node).right.cantity()
	}
	return sum
}

func createArray(array []*TreeNode, index *int, node *TreeNode) {
	array[*index] = node
	if (*node).left != nil {
		*index++
		createArray(array, index, (*node).left)
	}
	if (*node).right != nil {
		*index++
		createArray(array, index, (*node).right)
	}
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
}

func (node *TreeNode) isLeaf() bool {
	return (*node).left == nil && (*node).right == nil
}

func equal(node, other *TreeNode) bool {
	if ((*node).left != nil && (*other).left == nil) ||
		((*node).right != nil && (*other).right == nil) ||
		((*other).left != nil && (*node).left == nil) ||
		((*other).right != nil && (*node).right == nil) {
		return false
	}

	if (*node).left != nil && !equal((*node).left, (*other).left) {
		return false
	}

	if (*node).right != nil && !equal((*node).right, (*other).right) {
		return false
	}
	return true
}

func (node *TreeNode) copy() *TreeNode {
	newNode := &TreeNode{}
	if (*node).left != nil {
		(*newNode).left = (*node).left.copy()
	}
	if (*node).right != nil {
		(*newNode).right = (*node).right.copy()
	}
	return newNode
}

func mutateTree(nodes []*TreeNode, actualIndex int, mutations int, best *int, otherTree *TreeNode) {
	if equal(nodes[0], otherTree) {
		if mutations < *best {
			*best = mutations
		}
		return
	}
	for index := actualIndex; index < len(nodes); index++ {
		if !nodes[index].isLeaf() {
			temp := (*nodes[index]).right
			(*nodes[index]).right = (*nodes[index]).left
			(*nodes[index]).left = temp
			mutateTree(nodes, index+1, mutations+1, best, otherTree)
			temp = (*nodes[index]).right
			(*nodes[index]).right = (*nodes[index]).left
			(*nodes[index]).left = temp
		}
	}
}

func initTree() *TreeNode {
	return &TreeNode{}
}
