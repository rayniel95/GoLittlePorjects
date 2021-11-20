package mitad

type Node struct {
	value int
	sons  []*Node
	taken bool
}

func (node *Node) equal(other *Node) bool {
	length_other := 0
	for _, son := range (*other).sons {
		if !(*son).taken {
			length_other++
		}
	}
	length_this := 0
	for _, son := range (*node).sons {
		if !(*son).taken {
			length_this++
		}
	}
	return length_other == length_this && (*node).value == (*other).value
}

// TODO - restruncture the workspace for import modules
func (node *Node) recolectSubTrees() []*Node {
	array := make([]*Node, 0)
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
		if (*son).taken {
			has_taken = true
			continue
		}
		if has_taken && !mitad(son, (*other).sons[index-1]) {
			return false
		} else if !has_taken && !mitad(son, (*other).sons[index]) {
			return false
		}
	}
	return true
}

func solveMitad(node *Node) *Node{
	subtrees:=node.recolectSubTrees()
	for _, sub:=range subtrees{
		(*sub).taken=true
		if mitad(node, sub){
			return sub
		}
		(*sub).taken=false
	}
	return nil
}
 