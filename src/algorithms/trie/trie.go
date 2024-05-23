package trie

type Node struct {
	val      rune
	children [26]*Node
	isWord   bool
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: &Node{},
	}
}

func (t *Trie) Insert(word string) {
	currentNode := t.Root

	for i, w := range word {
		runeIdx := word[i] - 'a'

		if currentNode.children[runeIdx] == nil {
			currentNode.children[runeIdx] = &Node{val: w}
		}

		currentNode = currentNode.children[runeIdx]
	}

	currentNode.isWord = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.Root

	for i := range word {
		runeIdx := word[i] - 'a'

		if currentNode.children[runeIdx] == nil {
			return false
		}

		currentNode = currentNode.children[runeIdx]
	}

	return currentNode.isWord
}

func (t *Trie) Query(word string) []string {
	curr := t.Root
	var res []string

	for i := range word {
		idx := word[i] - 'a'

		if curr.children[idx] == nil {
			return res
		}

		curr = curr.children[idx]
	}
	prefix := word[:len(word)-1]

	dfs(curr, &res, prefix)

	return res
}

func dfs(node *Node, res *[]string, prefix string) {
	if node.isWord {
		result := prefix + string(node.val)
		*res = append(*res, result)
	}

	for _, nodeAux := range node.children {
		if nodeAux != nil {
			dfs(nodeAux, res, prefix+string(node.val))
		}
	}
}
