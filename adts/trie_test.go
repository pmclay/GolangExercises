/*
 *  © 2020 Michael Page. All rights reserved.
 */

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

func TestTrieAddPrefix(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeAddPrefix(pt, t)
}

func TestTrieAddToPrefix(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeAddToPrefix(pt, t)
}

func TestTrieContainsEmpty(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeContainsEmpty(pt, t)
}

func TestTrieContainsDoesNotExist(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeContainsDoesNotExist(pt, t)
}

func TestTrieContainsExistsAlone(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeContainsExistsAlone(pt, t)
}

func TestTrieContainsExists(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeContainsExists(pt, t)
}

func TestTrieContainsPrefixDoesNotExist(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeContainsPrefixDoesNotExist(pt, t)
}

func TestTrieContainsPrefixExists(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeContainsPrefixExists(pt, t)
}

func TestTrieContainsPrefixExistsForPrefix(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeContainsPrefixExistsForPrefix(pt, t)
}

func TestTrieContainsPostfixDoesNotExist(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeContainsPostfixDoesNotExist(pt, t)
}

func TestTrieContainsPostfixExists(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeContainsPostfixExists(pt, t)
}

func TestTrieSize(t *testing.T) {
	var pt PrefixTree = NewTrie()
	testPrefixTreeSize(pt, t)
}
