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
