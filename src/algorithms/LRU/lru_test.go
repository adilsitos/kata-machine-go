package lru_test

import (
	"testing"

	lru "github.com/adilsitos/kata-machine-golang/src/algorithms/LRU"
	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	lru := lru.NewLRU(3)

	assert.Equal(t, nil, lru.Get("foo"))
	lru.Update("foo", 1)
	assert.Equal(t, 1, lru.Get("foo"))

	lru.Update("bar", 420)
	assert.Equal(t, 420, lru.Get("bar"))

	lru.Update("baz", 1337)
	assert.Equal(t, 1337, lru.Get("baz"))

	lru.Update("ball", 1234)
	assert.Equal(t, 1234, lru.Get("ball"))

	assert.Equal(t, nil, lru.Get("foo"))

	assert.Equal(t, 420, lru.Get("bar"))

	lru.Update("foo", 69)

	assert.Equal(t, 420, lru.Get("bar"))
	assert.Equal(t, 69, lru.Get("foo"))

	assert.Equal(t, nil, lru.Get("baz"))
}
