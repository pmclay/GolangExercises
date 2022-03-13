/*
 *  © 2020 Michael Page. All rights reserved.
 */

package main

import (
	"strconv"
	"testing"
)

// all test cases expect the PrefixTree under test
// to be a freshly instantiated object

func testPrefixTreeAddEmpty(pt PrefixTree, t *testing.T) {
	obj := NewElement("abcde")
	pt.Add(obj)

	s := pt.Size()
	if s != 1 {
		t.Errorf("Adding an item to empty prefix tree caused; size=%d", s)
	}

	if !pt.Contains(obj.String()) {
		t.Errorf("Item was not added to empty prefix tree")
	}
}

func testPrefixTreeAddNonEmpty(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("fghi")
	obj2 := NewElement("abcde")
	pt.Add(obj1)
	pt.Add(obj2)

	s := pt.Size()
	if s != 2 {
		t.Errorf("Adding an item to prefix tree caused; size=%d", s)
	}

	if !pt.Contains(obj2.String()) {
		t.Errorf("Item was not added")
	}

	if !pt.Contains(obj1.String()) {
		t.Errorf("Item was not added")
	}
}

func testPrefixTreeAddOnExisting(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("abcde")
	obj2 := NewElement("abcde")
	pt.Add(obj1)
	pt.Add(obj2)

	s := pt.Size()
	if s != 2 {
		t.Errorf("Adding an item to prefix tree caused; size=%d", s)
	}

	if !pt.Contains(obj2.String()) {
		t.Errorf("Item was not added")
	}

	if !pt.Contains(obj1.String()) {
		t.Errorf("Item was not added")
	}
}

func testPrefixTreeAddPrefix(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("abcde")
	obj2 := NewElement("abc")
	pt.Add(obj1)
	pt.Add(obj2)

	s := pt.Size()
	if s != 2 {
		t.Errorf("Adding an item to prefix tree caused; size=%d", s)
	}

	if !pt.Contains(obj2.String()) {
		t.Errorf("Item which is a prefix of an existing item was not added")
	}

	if !pt.Contains(obj1.String()) {
		t.Errorf("Item was not added")
	}
}

func testPrefixTreeAddToPrefix(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("abcde")
	pt.Add(obj1)
	pt.Add(obj2)

	s := pt.Size()
	if s != 2 {
		t.Errorf("Adding an item to prefix tree caused; size=%d", s)
	}

	if !pt.Contains(obj2.String()) {
		t.Errorf("Item which extends an existing item was not added")
	}

	if !pt.Contains(obj1.String()) {
		t.Errorf("Item was not added")
	}
}

func testPrefixTreeContainsEmpty(pt PrefixTree, t *testing.T) {
	obj := NewElement("abcde")

	if pt.Contains(obj.String()) {
		t.Errorf("Item was found in empty prefix tree")
	}
}

func testPrefixTreeContainsDoesNotExist(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("fghi")
	obj2 := NewElement("abcde")
	pt.Add(obj1)

	if pt.Contains(obj2.String()) {
		t.Errorf("Non-existing item was found")
	}
}

func testPrefixTreeContainsExistsAlone(pt PrefixTree, t *testing.T) {
	obj := NewElement("abcde")
	pt.Add(obj)

	if !pt.Contains(obj.String()) {
		t.Errorf("Existing item was not found")
	}
}

func testPrefixTreeContainsExists(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("abcde")
	obj2 := NewElement("fghij")
	pt.Add(obj1)
	pt.Add(obj2)

	if !pt.Contains(obj2.String()) {
		t.Errorf("Existing item was not found")
	}
}

func testPrefixTreeContainsPrefixDoesNotExist(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("abcde")
	obj2 := NewElement("abc")
	pt.Add(obj1)

	if pt.Contains(obj2.String()) {
		t.Errorf("Non-existing prefix of item was found")
	}
}

func testPrefixTreeContainsPrefixExists(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("abcde")
	obj2 := NewElement("abc")
	pt.Add(obj1)
	pt.Add(obj2)

	if !pt.Contains(obj2.String()) {
		t.Errorf("Existing prefix of item was not found")
	}
}

func testPrefixTreeContainsPrefixExistsForPrefix(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("abcde")
	obj3 := NewElement("abcd")
	pt.Add(obj1)
	pt.Add(obj2)
	pt.Add(obj3)

	if !pt.Contains(obj1.String()) {
		t.Errorf("Existing prefix of item was not found")
	}
}

func testPrefixTreeContainsPostfixDoesNotExist(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("abcde")
	pt.Add(obj1)

	if pt.Contains(obj2.String()) {
		t.Errorf("Non-existing item was found")
	}
}

func testPrefixTreeContainsPostfixExists(pt PrefixTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("abcde")
	pt.Add(obj1)
	pt.Add(obj2)

	if !pt.Contains(obj2.String()) {
		t.Errorf("Existing item was not found")
	}
}

func testPrefixTreeSize(pt PrefixTree, t *testing.T) {
	i := 0
	for ; i < 25; i++ {
		if pt.Size() != i {
			t.Errorf("A prefix tree with %d items had size=%d", i, pt.Size())
		}

		pt.Add(NewElement(strconv.Itoa(i)))
	}

	if pt.Size() != i {
		t.Errorf("A prefix tree with %d items had size=%d", i, pt.Size())
	}
}
