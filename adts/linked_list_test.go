/*
 *  © 2020 Michael Page. All rights reserved.
 */

package main

import (
	"strconv"
	"testing"
)

func TestLinkedListPrependEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListPrependEmpty(l, t)
}

func TestLinkedListPrependNonEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListPrependNonEmpty(l, t)
}

func TestLinkedListAppendEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListAppendEmpty(l, t)
}

func TestLinkedListAppendNonEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListAppendNonEmpty(l, t)
}

func TestLinkedListInsertEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListInsertEmpty(l, t)
}

func TestLinkedListInsertOutofRange(t *testing.T) {
	var l List = NewLinkedList()
	testListInsertOutofRange(l, t)
}

func TestLinkedListInsertMiddle(t *testing.T) {
	var l List = NewLinkedList()
	testListInsertMiddle(l, t)
}

func TestLinkedListInsertAppend(t *testing.T) {
	var l List = NewLinkedList()
	testListInsertAppend(l, t)
}

func TestLinkedListGetEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListGetEmpty(l, t)
}

func TestLinkedListGetOutofRange(t *testing.T) {
	var l List = NewLinkedList()
	testListGetOutofRange(l, t)
}

func TestLinkedListGetValid(t *testing.T) {
	var l List = NewLinkedList()
	testListGetValid(l, t)
}

func TestLinkedListSetEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListSetEmpty(l, t)
}

func TestLinkedListSetOutofRange(t *testing.T) {
	var l List = NewLinkedList()
	testListSetOutofRange(l, t)
}

func TestLinkedListSetValid(t *testing.T) {
	var l List = NewLinkedList()
	testListSetValid(l, t)
}

func TestLinkedListRemoveFirstEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListRemoveFirstEmpty(l, t)
}

func TestLinkedListRemoveFirstNonEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListRemoveFirstNonEmpty(l, t)
}

func TestLinkedListRemoveLastEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListRemoveLastEmpty(l, t)
}

func TestLinkedListRemoveLastNonEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListRemoveLastNonEmpty(l, t)
}

func TestLinkedListRemoveEmpty(t *testing.T) {
	var l List = NewLinkedList()
	testListRemoveEmpty(l, t)
}

func TestLinkedListRemoveOutofRange(t *testing.T) {
	var l List = NewLinkedList()
	testListRemoveOutofRange(l, t)
}

func TestLinkedListRemoveValid(t *testing.T) {
	var l List = NewLinkedList()
	testListRemoveValid(l, t)
}

func TestLinkedListSize(t *testing.T) {
	var l List = NewLinkedList()
	testListSize(l, t)
}

func TestLinkedListCapacity(t *testing.T) {
	var l List = NewLinkedList()

	// for linked list capacity should = size
	capacity := l.Capacity()
	size := l.Size()
	if capacity != size {
		t.Errorf("Emtpy list capacity does not equal size; capacity=%d size=%d", capacity, size)
	}

	for i := 0; i < 15; i++ {
		l.Append(NewElement(strconv.Itoa(i)))

		capacity = l.Capacity()
		size = l.Size()
		if capacity != size {
			t.Errorf("Capacity does not equal size; capacity=%d size=%d", capacity, size)
		}
	}

	for i := size; i > 0; i-- {
		l.RemoveFirst()

		capacity = l.Capacity()
		size = l.Size()
		if capacity != size {
			t.Errorf("Capacity does not equal size; capacity=%d size=%d", capacity, size)
		}
	}

	if capacity != size {
		t.Errorf("Emtpy list capacity does not equal size; capacity=%d size=%d", capacity, size)
	}
}

func TestLinkedListString(t *testing.T) {
	var l List = NewLinkedList()
	testListString(l, t)
}
