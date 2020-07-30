package main

import (
	"strconv"
	"testing"
)

func TestDynamicArrayPrependEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListPrependEmpty(l, t)
}

func TestDynamicArrayPrependNonEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListPrependNonEmpty(l, t)
}

func TestDynamicArrayAppendEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListAppendEmpty(l, t)
}

func TestDynamicArrayAppendNonEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListAppendNonEmpty(l, t)
}

func TestDynamicArrayInsertEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListInsertEmpty(l, t)
}

func TestDynamicArrayInsertOutofRange(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListInsertOutofRange(l, t)
}

func TestDynamicArrayInsertMiddle(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListInsertMiddle(l, t)
}

func TestDynamicArrayInsertAppend(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListInsertAppend(l, t)
}

func TestDynamicArrayGetEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListGetEmpty(l, t)
}

func TestDynamicArrayGetOutofRange(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListGetOutofRange(l, t)
}

func TestDynamicArrayGetValid(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListGetValid(l, t)
}

func TestDynamicArraySetEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListSetEmpty(l, t)
}

func TestDynamicArraySetOutofRange(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListSetOutofRange(l, t)
}

func TestDynamicArraySetValid(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListSetValid(l, t)
}

func TestDynamicArrayRemoveFirstEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListRemoveFirstEmpty(l, t)
}

func TestDynamicArrayRemoveFirstNonEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListRemoveFirstNonEmpty(l, t)
}

func TestDynamicArrayRemoveLastEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListRemoveLastEmpty(l, t)
}

func TestDynamicArrayRemoveLastNonEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListRemoveLastNonEmpty(l, t)
}

func TestDynamicArrayRemoveEmpty(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListRemoveEmpty(l, t)
}

func TestDynamicArrayRemoveOutofRange(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListRemoveOutofRange(l, t)
}

func TestDynamicArrayRemoveValid(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListRemoveValid(l, t)
}

func TestDynamicArraySize(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListSize(l, t)
}

func TestDynamicArrayCapacity(t *testing.T) {
	expected := 2
	var l List = NewDynamicArray(expected)

	capacity := l.Capacity()
	if capacity != expected {
		t.Errorf("Unexpected initial list capacity; capacity=%d", capacity)
	}

	for i := 1; i <= 15; i++ {
		if i > expected {
			expected *= 2
		}

		l.Append(NewElement(strconv.Itoa(i)))
		capacity = l.Capacity()
		if capacity != expected {
			t.Errorf("Unexpected capacity after adding %d items; capacity=%d", i, capacity)
		}
	}

	size := l.Size()
	for size > 0 {
		l.RemoveFirst()

		capacity = l.Capacity()
		if capacity != expected {
			t.Errorf("Unexpected list capacity; capacity=%d", capacity)
		}

		size = l.Size()
	}
}

func TestDynamicArrayString(t *testing.T) {
	var l List = NewDynamicArray(2)
	testListString(l, t)
}
