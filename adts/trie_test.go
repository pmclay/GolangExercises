package main

import (
	"testing"
)

func TestTrieAddEmpty(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeAddEmpty(pt, t)
}

func TestTrieAddNonEmpty(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeAddEmpty(pt, t)
}

func TestTrieAddOnExisting(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeAddEmpty(pt, t)
}
