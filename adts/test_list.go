/*
 *  © 2020 Michael Page. All rights reserved.
 */

package main

import (
	"strconv"
	"testing"
)

// all test cases expect the List under test
// to be a freshly instantiated object

func testListAppendEmpty(l List, t *testing.T) {
	l.Append(NewElement("abc"))

	s := l.Size()
	if s != 1 {
		t.Errorf("Appending to empty list caused list size to be %d not 1", s)
	}

	if item, _ := l.Get(s - 1); item.String() != "abc" {
		t.Error("Expected item was not appended to the empty list")
	}
}

func testListAppendNonEmpty(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Append(NewElement("123"))

	s := l.Size()
	if s != 2 {
		t.Errorf("List size=%d is not correct after appending new item", s)
	}

	if item, _ := l.Get(s - 1); item.String() != "123" {
		t.Error("Expected item was not appended to the list")
	}
}

func testListPrependEmpty(l List, t *testing.T) {
	l.Prepend(NewElement("abc"))

	s := l.Size()
	if s != 1 {
		t.Errorf("Prepending to empty list caused list size to be %d not 1", s)
	}

	if item, _ := l.Get(0); item.String() != "abc" {
		t.Error("Expected item was not prepended to the empty list")
	}
}

func testListPrependNonEmpty(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Prepend(NewElement("123"))

	s := l.Size()
	if s != 2 {
		t.Errorf("List size=%d is not correct after prepending new item", s)
	}

	if item, _ := l.Get(0); item.String() != "123" {
		t.Error("Expected item was not prepended to the list")
	}
}

func testListInsertEmpty(l List, t *testing.T) {
	if err := l.Insert(NewElement("abc"), 0); err != nil {
		t.Errorf("Inserting into empty list caused an error : %v", err)
	}

	s := l.Size()
	if s != 1 {
		t.Errorf("Inserting into empty list caused list size to be %d not 1", s)
	}

	if item, _ := l.Get(0); item.String() != "abc" {
		t.Error("Expected item was not inserted into the empty list")
	}
}

func testListInsertOutofRange(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Prepend(NewElement("123"))

	// index too small
	if err := l.Insert(NewElement("def"), -1); err == nil {
		t.Error("Inserting at out of bounds index did not error")
	}

	s := l.Size()
	if s != 2 {
		t.Errorf("Failed list insert caused; size=%d", s)
	}

	if item, _ := l.Get(0); item.String() != "123" {
		t.Errorf("Failed list insert caused; L[0]=%v", item)
	}

	if item, _ := l.Get(1); item.String() != "abc" {
		t.Errorf("Failed list insert caused; L[1]=%v", item)
	}

	// index too large
	if err := l.Insert(NewElement("def"), 5); err == nil {
		t.Error("Inserting at out of bounds index did not error")
	}

	s = l.Size()
	if s != 2 {
		t.Errorf("Failed list insert caused; size=%d", s)
	}

	if item, _ := l.Get(0); item.String() != "123" {
		t.Errorf("Failed list insert caused; L[0]=%v", item)
	}

	if item, _ := l.Get(1); item.String() != "abc" {
		t.Errorf("Failed list insert caused; L[1]=%v", item)
	}
}

func testListInsertMiddle(l List, t *testing.T) {
	l.Append(NewElement("123"))
	l.Prepend(NewElement("abc"))
	l.Append(NewElement("456"))

	if err := l.Insert(NewElement("def"), 1); err != nil {
		t.Errorf("Inserting into list caused an error : %v", err)
	}

	s := l.Size()
	if s != 4 {
		t.Errorf("List size=%d is not correct after inserting new item", s)
	}

	if item, _ := l.Get(0); item.String() != "abc" {
		t.Errorf("List insert had unexpected effect; L[0]=%v", item)
	}

	if item, _ := l.Get(1); item.String() != "def" {
		t.Errorf("List insert had unexpected effect; L[1]=%v", item)
	}

	if item, _ := l.Get(2); item.String() != "123" {
		t.Errorf("List insert had unexpected effect; L[2]=%v", item)
	}

	if item, _ := l.Get(3); item.String() != "456" {
		t.Errorf("List insert had unexpected effect; L[3]=%v", item)
	}
}

func testListInsertAppend(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Prepend(NewElement("123"))

	if err := l.Insert(NewElement("def"), 2); err != nil {
		t.Errorf("Inserting into list caused an error : %v", err)
	}

	s := l.Size()
	if s != 3 {
		t.Errorf("List size=%d is not correct after inserting new item", s)
	}

	if item, _ := l.Get(0); item.String() != "123" {
		t.Errorf("List insert had unexpected effect; L[0]=%v", item)
	}

	if item, _ := l.Get(1); item.String() != "abc" {
		t.Errorf("List insert had unexpected effect; L[1]=%v", item)
	}

	if item, _ := l.Get(2); item.String() != "def" {
		t.Errorf("List insert had unexpected effect; L[2]=%v", item)
	}
}

func testListGetEmpty(l List, t *testing.T) {
	item, err := l.Get(0)

	if err == nil {
		t.Error("Getting from an empty list did not cause an error")
	}

	if item != nil {
		t.Errorf("Getting from an empty list returned; item=%v", item)
	}

	if s := l.Size(); s != 0 {
		t.Errorf("Getting from empty list caused; size=%d", s)
	}
}

func testListGetOutofRange(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Prepend(NewElement("123"))

	// index too small
	item, err := l.Get(-1)

	if err == nil {
		t.Error("Getting at invalid list index did not cause an error")
	}

	if item != nil {
		t.Errorf("Getting at invalid list index returned; item=%v", item)
	}

	if s := l.Size(); s != 2 {
		t.Errorf("Getting from invalid list index caused; size=%d", s)
	}

	// index too large
	item, err = l.Get(2)

	if err == nil {
		t.Error("Getting at invalid list index did not cause an error")
	}

	if item != nil {
		t.Errorf("Getting at invalid list index returned; item=%v", item)
	}

	if s := l.Size(); s != 2 {
		t.Errorf("Getting from invalid list index caused; size=%d", s)
	}
}

func testListGetValid(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Prepend(NewElement("123"))

	item, err := l.Get(1)

	if err != nil {
		t.Error("Getting at valid list index caused an error")
	}

	if item.String() != "abc" {
		t.Errorf("Getting at valid list index returned wrong item; item=%v", item)
	}

	if s := l.Size(); s != 2 {
		t.Errorf("Getting from valid list index caused; size=%d", s)
	}
}

func testListSetEmpty(l List, t *testing.T) {
	err := l.Set(NewElement("abc"), 0)

	if err == nil {
		t.Error("Setting item for an empty list did not cause an error")
	}

	if s := l.Size(); s != 0 {
		t.Errorf("Setting item for empty list caused; size=%d", s)
	}
}

func testListSetOutofRange(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Prepend(NewElement("123"))

	// index too small
	err := l.Set(NewElement("def"), -1)

	if err == nil {
		t.Error("Setting item at invalid list index did not cause an error")
	}

	if s := l.Size(); s != 2 {
		t.Errorf("Setting item at invalid list index caused; size=%d", s)
	}

	if item, _ := l.Get(0); item.String() != "123" {
		t.Errorf("Setting item at invalid list index caused; L[0]=%v", item)
	}

	if item, _ := l.Get(1); item.String() != "abc" {
		t.Errorf("Setting item at invalid list index caused; L[1]=%v", item)
	}

	// index too large
	err = l.Set(NewElement("def"), 2)

	if err == nil {
		t.Error("Setting item at invalid list index did not cause an error")
	}

	if s := l.Size(); s != 2 {
		t.Errorf("Setting item at invalid list index caused; size=%d", s)
	}

	if item, _ := l.Get(0); item.String() != "123" {
		t.Errorf("Setting item at invalid list index caused; L[0]=%v", item)
	}

	if item, _ := l.Get(1); item.String() != "abc" {
		t.Errorf("Setting item at invalid list index caused; L[1]=%v", item)
	}
}

func testListSetValid(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Prepend(NewElement("123"))

	err := l.Set(NewElement("def"), 1)

	if err != nil {
		t.Error("Setting item at valid list index caused an error")
	}

	if s := l.Size(); s != 2 {
		t.Errorf("Setting item at valid list index caused; size=%d", s)
	}

	if item, _ := l.Get(0); item.String() != "123" {
		t.Errorf("Setting item at valid list index caused; L[0]=%v", item)
	}

	if item, _ := l.Get(1); item.String() != "def" {
		t.Errorf("Setting item at valid list index caused; L[1]=%v", item)
	}
}

func testListRemoveFirstEmpty(l List, t *testing.T) {
	item, err := l.RemoveFirst()

	if err == nil {
		t.Error("Removing first item from empty list did not cause an error")
	}

	if item != nil {
		t.Errorf("Removing first item from empty list returned; item=%v", item)
	}

	if s := l.Size(); s != 0 {
		t.Errorf("Removing first item from empty list caused; size=%d", s)
	}
}

func testListRemoveFirstNonEmpty(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	item, err := l.RemoveFirst()

	if err != nil {
		t.Error("Removing first item from non-empty list caused an error")
	}

	if item.String() != "abc" {
		t.Errorf("Removing first item from non-empty list returned; item=%v", item)
	}

	if s := l.Size(); s != 0 {
		t.Errorf("Removing first item from non-empty list caused; size=%d", s)
	}
}

func testListRemoveLastEmpty(l List, t *testing.T) {
	item, err := l.RemoveLast()

	if err == nil {
		t.Error("Removing last item from empty list did not cause an error")
	}

	if item != nil {
		t.Errorf("Removing last item from empty list returned; item=%v", item)
	}

	if s := l.Size(); s != 0 {
		t.Errorf("Removing last item from empty list caused; size=%d", s)
	}
}

func testListRemoveLastNonEmpty(l List, t *testing.T) {
	l.Prepend(NewElement("abc"))
	item, err := l.RemoveLast()

	if err != nil {
		t.Error("Removing last item from non-empty list caused an error")
	}

	if item.String() != "abc" {
		t.Errorf("Removing last item from non-empty list returned; item=%v", item)
	}

	if s := l.Size(); s != 0 {
		t.Errorf("Removing last item from non-empty list caused; size=%d", s)
	}
}

func testListRemoveEmpty(l List, t *testing.T) {
	item, err := l.Remove(0)

	if err == nil {
		t.Error("Removing item from empty list did not cause an error")
	}

	if item != nil {
		t.Errorf("Removing item from empty list returned; item=%v", item)
	}

	if s := l.Size(); s != 0 {
		t.Errorf("Removing item from empty list caused; size=%d", s)
	}
}

func testListRemoveOutofRange(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Prepend(NewElement("123"))

	// index too small
	item, err := l.Remove(-1)

	if err == nil {
		t.Error("Removing item at invalid list index did not cause an error")
	}

	if item != nil {
		t.Errorf("Removing item at invalid list index returned; item=%v", item)
	}

	if s := l.Size(); s != 2 {
		t.Errorf("Removing item at invalid list index caused; size=%d", s)
	}

	if item, _ := l.Get(0); item.String() != "123" {
		t.Errorf("Removing item at invalid list index caused; L[0]=%v", item)
	}

	if item, _ := l.Get(1); item.String() != "abc" {
		t.Errorf("Removing item at invalid list index caused; L[1]=%v", item)
	}

	// index too large
	item, err = l.Remove(2)

	if err == nil {
		t.Error("Removing item at invalid list index did not cause an error")
	}

	if item != nil {
		t.Errorf("Removing item at invalid list index returned; item=%v", item)
	}

	if s := l.Size(); s != 2 {
		t.Errorf("Removing item at invalid list index caused; size=%d", s)
	}

	if item, _ := l.Get(0); item.String() != "123" {
		t.Errorf("Removing item at invalid list index caused; L[0]=%v", item)
	}

	if item, _ := l.Get(1); item.String() != "abc" {
		t.Errorf("Removing item at invalid list index caused; L[1]=%v", item)
	}
}

func testListRemoveValid(l List, t *testing.T) {
	l.Append(NewElement("abc"))
	l.Prepend(NewElement("123"))

	item, err := l.Remove(0)

	if err != nil {
		t.Error("Removing item at valid list index caused an error")
	}

	if item.String() != "123" {
		t.Errorf("Removing item at valid list index returned; item=%v", item)
	}

	if s := l.Size(); s != 1 {
		t.Errorf("Removing item at valid list index caused; size=%d", s)
	}

	if item, _ := l.Get(0); item.String() != "abc" {
		t.Errorf("Removing item at valid list index caused; L[0]=%v", item)
	}
}

func testListSize(l List, t *testing.T) {
	size := l.Size()
	if size != 0 {
		t.Errorf("Empty list size=%d", size)
	}

	for i := 1; i <= 15; i++ {
		l.Prepend(NewElement(strconv.Itoa(i)))

		size = l.Size()
		if size != i {
			t.Errorf("%d items should be in the list but size=%d", i, size)
		}
	}

	for i := size; i > 0; i-- {
		item, _ := l.RemoveFirst()

		if item.String() != strconv.Itoa(i) {
			t.Errorf("Removing item %d from the list returned; item=%v", i, item)
		}

		size = l.Size()
		if size != (i - 1) {
			t.Errorf("%d items should be in the list but size=%d", i-1, size)
		}
	}

	if size != 0 {
		t.Errorf("Empty list size=%d", size)
	}
}

func testListString(l List, t *testing.T) {
	for i := 1; i <= 15; i++ {
		l.Append(NewElement(strconv.Itoa(i)))
	}

	if str := l.String(); str != "[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]" {
		t.Errorf("Invalid string representation for list of integers 1-15; str=%s", str)
	}
}
