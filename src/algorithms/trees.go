package algorithms

// pre order: Apply some logic at the current node and recurse the left side.
// After the left side, recurse the right side

// in order: Recurse the left side, if there isn't more node to go on the left, apply
// the logic to the current node (leaf). Recurse the right side

// post order:
// in order: Recurse the left side, recurse the right side. If is a leaf, apply the logic

type BinaryNode struct {
	Value int
	Right *BinaryNode
	Left  *BinaryNode
}

type BinaryTree struct {
	Root *BinaryNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

func PreOrderWalk(node *BinaryNode, path *[]int) []int {
	if node == nil {
		return nil
	}

	// pre
	*path = append(*path, node.Value)
	PreOrderWalk(node.Left, path)
	PreOrderWalk(node.Right, path)
	// post

	return *path
}

func PreOrder(binTree *BinaryTree) []int {
	path := make([]int, 0)
	return PreOrderWalk(binTree.Root, &path)
}

func InOrderWalk(root *BinaryNode, path *[]int) []int {
	if root == nil {
		return nil
	}
	// pre
	InOrderWalk(root.Left, path)
	//Execution
	*path = append(*path, root.Value)
	// Post
	InOrderWalk(root.Right, path)

	return *path
}

func InOrder(binTree *BinaryTree) []int {
	path := make([]int, 0)
	return InOrderWalk(binTree.Root, &path)
}

// good when you want to free the nodes
func PostOrderWalk(root *BinaryNode, path *[]int) []int {
	if root == nil {
		return nil
	}
	PostOrderWalk(root.Left, path)
	PostOrderWalk(root.Right, path)

	*path = append(*path, root.Value)

	return *path
}

func PostOrder(binTree *BinaryTree) []int {
	path := make([]int, 0)
	return PostOrderWalk(binTree.Root, &path)
}

func (tree *BinaryTree) InsertIntoTree(value int) {
	newNode := BinaryNode{Value: value}

	if tree.Root == nil {
		tree.Root = &newNode
		return
	}

	queue := NewQueue()
	queue.Enqueue(tree.Root)

	for queue.size != 0 {
		temp := queue.Dequeue().(*BinaryNode)

		if temp.Left == nil {
			temp.Left = &newNode
			return
		} else {
			queue.Enqueue(temp.Left)
		}

		if temp.Right == nil {
			temp.Right = &newNode
			return
		} else {
			queue.Enqueue(temp.Right)
		}
	}

}
