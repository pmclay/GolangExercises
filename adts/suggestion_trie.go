/*
 *  © 2022 Michael Page. All rights reserved.
 */

package main

import ()

type alphabeticPostfixes struct {
	key         trieNode
	predecessor *alphabeticPostfixes
	successor   *alphabeticPostfixes
}

func (head *alphabeticPostfixes) findInsert(r rune) *alphabeticPostfixes {
	prev := head.predecessor
	curr := head
	for curr != nil && curr.key.value < r {
		prev = curr
		curr = curr.successor
	}

	return prev
}

func (a *alphabeticPostfixes) insert(str []rune) {
	if len(str) == 0 || a == nil {
		return
	}

	insertPosition := a.findInsert(str[0])

	if insertPosition != nil {
		newPostfix := new(alphabeticPostfixes)
		newPostfix.key = trieNode{value: str[0]}

		newPostfix.successor = insertPosition.successor
		insertPosition.successor.predecessor = newPostfix
		insertPosition.successor = newPostfix
		newPostfix.predecessor = insertPosition

		insertPosition = newPostfix
	} else {
		if a.key.value == str[0] {
			insertPosition = insertPosition.successor
		} else {
			insertPosition = new(alphabeticPostfixes)
			insertPosition.key = trieNode{value: str[0]}
			insertPosition.successor = a
			a.predecessor = insertPosition
		}
	}

	insertPosition.key.build(str[1:])
}

type trieNode struct {
	value     rune
	postfixes *alphabeticPostfixes
}

func (n *trieNode) build(str []rune) {
	if len(str) == 0 || n == nil {
		return
	}
	if n.postfixes == nil {
		n.postfixes = new(alphabeticPostfixes)
	}

	if n.value == str[0] {
		n.postfixes.insert(str[1:])
	} else {

	}
}

type SuggestionTrie struct {
	root alphabeticPostfixes
	size int
}

func (s *SuggestionTrie) Add(obj Item) {
	s.root.insert([]rune(obj.String()))
	s.size += 1
}

func (s *SuggestionTrie) Contains(key string) bool {

}

func (s *SuggestionTrie) Size() int {

}
