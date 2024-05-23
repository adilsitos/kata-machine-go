package algorithms

func BreadthFirstSearch(binTree *BinaryTree, value int) bool {
	q := NewQueue()

	q.Enqueue(binTree.Root)

	for q.size != 0 {
		val := q.Dequeue().(*BinaryNode)

		if val.Value == value {
			return true
		}

		if val.Left != nil {
			q.Enqueue(val.Left)
		}

		if val.Right != nil {
			q.Enqueue(val.Right)
		}
	}

	return false
}
