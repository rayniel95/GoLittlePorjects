package mitad

type Node struct {
	value int
	sons  []*Node
	token bool
}

func (node *Node) equal(other *Node) bool {
	length_other := 0
	for _, son := range (*other).sons {
		if !(*son).token {
			length_other++
		}
	}
	length_this := 0
	for _, son := range (*node).sons {
		if !(*son).token {
			length_this++
		}
	}
	return length_other == length_this && (*node).value == (*other).value
}

// TODO - restruncture the workspace for import modules
func (node *Node) recolectSubTrees() []*Node {
	array := make([]*Node, len((*node).sons))
	array = append(array, (*node).sons...)

	for _, son := range (*node).sons {
		array = append(array, son.recolectSubTrees()...)
	}
	return array
}

func mitad(node *Node, other *Node) bool {
	if !node.equal(other) {
		return false
	}
	has_taken := false
	for index, son := range (*node).sons {
		if (*son).token {
			has_taken = true
		}
		if has_taken && !mitad(son, (*other).sons[index-1]) {
			return false
		} else if !has_taken && !mitad(son, (*other).sons[index]) {
			return false
		}
	}
	return true
}
