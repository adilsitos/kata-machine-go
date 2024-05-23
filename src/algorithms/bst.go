package algorithms

import (
	"fmt"
)

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

type BST struct {
	Root *BSTNode
}

func NewBST() *BST {
	return &BST{}
}

func (node *BSTNode) Find(val int) bool {
	if node == nil {
		return false
	}

	if node.val == val {
		return true
	}

	// value is greater than the current node
	// we need to search on the right side of
	// the current node
	if val > node.val {
		return node.right.Find(val)
	}

	if val <= node.val {
		return node.left.Find(val)
	}

	return false
}

func (bst *BST) Insert(val int) {

	if bst.Root == nil {
		bst.Root = &BSTNode{val: val}
		return
	}

	bst.Root.Insert(val)
}

func (node *BSTNode) Insert(val int) {
	if val > node.val {
		if node.right == nil {
			node.right = &BSTNode{val: val}
			return
		}

		node.right.Insert(val)
	}

	if val <= node.val {
		if node.left == nil {
			node.left = &BSTNode{val: val}
			return
		}

		node.left.Insert(val)
	}
}

// cases
// first case: the node is a leaf, so we can just remove it.
// the tree is still valid

// seconde case: Only one child. Get the parent pointer, and point it
// to the child of the current node. Set parent to child

// third case: The current node has two childs. In this case, we have two approaches: get
// the biggest element on the small side, or get the smallest element on the large side.
// When this element is find, it is necessary to check if it has one child or is a leaf (it does not have any child)
// . If there is a child (it will be a left child) we need to move this left child to the parent node, and the current node
// will be in place of the node that is being deleted.

// question: When choose the largest element on the small side, and when choose the
// smallest element on the large side?
// A: If the current node (the node that is being deleted) knows which side has the
// bigger height, it's better to choose the side with the biggest height, balacing the tree
func (bst *BST) Deletion(val int) {
	remove(bst.Root, val)
}

func remove(node *BSTNode, val int) *BSTNode {
	if node == nil {
		return nil
	}

	fmt.Println(node.val)
	fmt.Printf("%p\n", &node)
	if node.left != nil {
		fmt.Println("left", node.left.val)
		fmt.Printf("%p\n", node.left)
	}
	if node.right != nil {
		fmt.Println("right", node.right.val)
		fmt.Printf("%p\n", node.right)
	}

	if val < node.val {
		node.left = remove(node.left, val)
		return node
	}

	if val > node.val {
		node.right = remove(node.right, val)
		return node
	}

	// the following line is unecessary. If the code gets here,
	// it is only possible to be equal
	// find := node.val == val

	// 1st case, the current node does not have children
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	// 2nd case. Only one branch is nil
	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.right
		return node
	}

	// 3rd case

	biggestOnTheLeft := node.left

	for biggestOnTheLeft != nil {
		// find the biggest value on the left side

		if biggestOnTheLeft.right == nil {
			break
		}

		biggestOnTheLeft = biggestOnTheLeft.right
	}

	node.val = biggestOnTheLeft.val
	node.left = remove(node.left, node.val)

	return node
}

func (bst *BST) Print() {
	fmt.Println("-----------------------------")
	stringify(bst.Root, 0)
	fmt.Println("-----------------------------")
}

func stringify(n *BSTNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "			"
		}
		format += "---[ "
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%d\n", n.val)
		stringify(n.left, level)
	}

}
