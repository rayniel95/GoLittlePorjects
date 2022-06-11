package treemincost

import "math"

func calculus(node *TreeNode) int {
	left := 0
	right := 0
	if (*node).left != nil {
		left = calculus((*node).left)
	}
	if (*node).right != nil {
		right = calculus((*node).right)
	}
	return left + right + (int(math.Pow(2, float64((*node).deep))) * (*node).value)
}

func constructTree(array []int) int {
	MaxUint := ^uint(0)
	MaxInt := int(MaxUint >> 1)

	best := &MaxInt
	levels := make([]bool, len(array))

	for index, value := range array {
		levels[index] = true
		buildTree(array, levels, &TreeNode{value: value, deep: 0}, best)
		levels[index] = false
	}
	return *best
}

type TreeNode struct {
	left, right *TreeNode
	value       int
	deep        int
}

func (node *TreeNode) copy() *TreeNode {
	newNode := &TreeNode{value: (*node).value, deep: (*node).deep}
	if (*node).left != nil {
		(*newNode).left = (*node).left.copy()
	}
	if (*node).right != nil {
		(*newNode).right = (*node).right.copy()
	}
	return newNode
}

func (node *TreeNode) insert(value int) bool {
	if value > (*node).value {
		if (*node).right != nil {
			return (*node).right.insert(value)
		}
		(*node).right = &TreeNode{value: value, deep: (*node).deep + 1}
	} else if value < (*node).value {
		if (*node).left != nil {
			return (*node).left.insert(value)
		}
		(*node).left = &TreeNode{value: value, deep: (*node).deep + 1}
	}
	return true
}

func allTaken(taken []bool) bool {
	for _, value := range taken {
		if !value {
			return false
		}
	}
	return true
}

func buildTree(array []int, taken []bool, root *TreeNode, best *int) {
	if allTaken(taken) {
		result := calculus(root)
		if result < *best {
			*best = result
		}
		return
	}
	for index, value := range array {
		if !taken[index] {
			taken[index] = true
			newRoot := root.copy()
			newRoot.insert(value)
			buildTree(array, taken, newRoot, best)
			taken[index] = false
		}
	}
}
