package trie_test

import (
	"testing"

	"github.com/adilsitos/kata-machine-golang/src/algorithms/trie"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := trie.NewTrie()

	trie.Insert("adilson")
	trie.Insert("test")
	trie.Insert("cat")
	trie.Insert("cattle")
	trie.Insert("cats")
	trie.Insert("category")
	trie.Insert("catheterization")

	assert.Equal(t, true, trie.Search("test"))
	assert.Equal(t, false, trie.Search("dont"))

	query := "cat"
	queryRes := []string{"cat", "category", "catheterization", "cats", "cattle"}

	assert.Equal(t, queryRes, trie.Query(query))
}
