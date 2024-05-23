package algorithms

func DFS(binTree *BinaryTree, value int) bool {
	stack := NewStack()

	stack.Push(binTree.Root)

	for stack.size != 0 {
		aux := stack.Pop().(*BinaryNode)

		if aux.Value == value {
			return true
		}

		if aux.Right != nil {
			stack.Push(aux.Right)
		}

		if aux.Left != nil {
			stack.Push(aux.Left)
		}
	}

	return false
}
