package main

import ()

type trienode struct {
	endEntry bool
	children map[rune]*trienode
}

type Trie struct {
	root *trienode
	size int
}

func NewTrie() *Trie {
	return new(Trie)
}

func (t *Trie) Add(obj Item) {
	t.size += 1
}

func (t *Trie) Contains(key string) bool {
	return false
}

func (t *Trie) Size() int {
	return t.size
}
