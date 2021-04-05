package arboreando

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value interface{}
}

type QueueNode struct {
	value interface{}
	next  *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func (q *Queue) enqueue(value interface{}) {
	if q.len() == 0 {
		node := &QueueNode{value: value}
		(*q).head = node
		(*q).tail = node
		return
	}
	newNode := &QueueNode{value: value}
	(*((*q).tail)).next = newNode
	(*q).tail = newNode
}

func (q *Queue) pop() interface{} {
	value := (*((*q).head)).value
	(*q).head = (*((*q).head)).next
	return value
}

func (q *Queue) len() int {
	count := 0
	pointer := (*q).head
	for ; pointer != nil; pointer = (*pointer).next {
		count++
	}
	return count
}

type operation func(interface{}, interface{}) interface{}

type IOperation interface {
	operation(interface{}, interface{}) interface{}
}

func operaInt(a interface{}, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func (node *TreeNode) copy() *TreeNode {
	newNode := &TreeNode{value: (*node).value}
	if (*node).left != nil {
		(*newNode).left = (*node).left.copy()
	}
	if (*node).right != nil {
		(*newNode).right = (*node).right.copy()
	}
	return newNode
}

func applyOp(left *TreeNode, right *TreeNode, op operation) {
	(*left).value = op((*left).value, (*right).value)
	if (*left).left != nil && (*right).left != nil {
		applyOp((*left).left, (*right).left, op)
	}
	if (*left).right != nil && (*right).right != nil {
		applyOp((*left).right, (*right).right, op)
	}
}

func orborear(left *TreeNode, right *TreeNode, op operation) *TreeNode {
	newTree := left.copy()
	queue := &Queue{}

	queue.enqueue(newTree)
	for queue.len() > 0 {
		root := (queue.pop()).(*TreeNode)
		applyOp(root, right, op)
		if !root.isLeaf() {
			if (*root).left != nil {
				queue.enqueue((*root).left)
			}
			if (*root).right != nil {
				queue.enqueue((*root).right)
			}
		}
	}
	return newTree
}

func (node *TreeNode) isLeaf() bool {
	return (*node).left == nil && (*node).right == nil
}
