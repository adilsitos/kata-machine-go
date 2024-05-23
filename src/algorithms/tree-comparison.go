package algorithms

// check if two binary trees are equal, both in value and in shape

// using DFS, because DFS preserves the shape of the tree

func Compare(a, b *BinaryNode) bool {

	// shape comparison
	if a == nil && b == nil {
		return true
	}

	// shape comparison
	if a == nil || b == nil {
		return false
	}

	// value comparison
	if a.Value != b.Value {
		return false
	}

	return Compare(a.Left, b.Left) && Compare(a.Right, b.Right)
}
