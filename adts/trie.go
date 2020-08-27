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
	t := new(Trie)
	t.root = new(trienode)
	t.root.children = make(map[rune]*trienode)

	return t
}

func (t *Trie) Add(obj Item) {
	current := t.root
	for _, char := range obj.String() {
		next, ok := current.children[char]

		if !ok {
			next = new(trienode)
			next.children = make(map[rune]*trienode)
			current.children[char] = next
		}

		current = next
	}

	current.endEntry = true

	t.size += 1
}

func (t *Trie) Contains(key string) bool {
	current := t.root
	for _, char := range key {
		next, ok := current.children[char]

		if !ok {
			return false
		}

		current = next
	}

	return current.endEntry
}

func (t *Trie) Size() int {
	return t.size
}
