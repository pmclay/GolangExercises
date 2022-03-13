/*
 *  © 2020 Michael Page. All rights reserved.
 */

package main

import (
	"strconv"
	"testing"
)

func getTestSearchTreeData() [20]*Element {

	testData := [20]*Element{}

	for i := 0; i < 20; i++ {
		testData[i] = NewElement(strconv.Itoa(i))
	}

	return testData
}

func testSearchTreeAddEmpty(st SearchTree, t *testing.T) {
	obj := NewElement("abc")

	st.Add(obj)

	size := st.Size()
	if size != 1 {
		t.Errorf("Adding one item to tree caused; size=%d", size)
	}

	if !st.Contains(obj.Hash()) {
		t.Errorf("Item was not added to emtpy tree")
	}
}

func testSearchTreeAddNonEmpty(st SearchTree, t *testing.T) {
	obj1 := NewElement("abc")
	st.Add(obj1)

	obj2 := NewElement("def")
	st.Add(obj2)

	size := st.Size()
	if size != 2 {
		t.Errorf("Adding an item to tree caused; size=%d", size)
	}

	if !st.Contains(obj1.Hash()) {
		t.Errorf("Pre-existing item was not found in tree")
	}

	if !st.Contains(obj2.Hash()) {
		t.Errorf("Added item was not found in tree")
	}
}

func testSearchTreeDeleteEmpty(st SearchTree, t *testing.T) {
	obj := NewElement("abc")

	err := st.Delete(obj.Hash())

	if err == nil {
		t.Errorf("Deleting non-existent item from tree did not error")
	}

	size := st.Size()
	if size != 0 {
		t.Errorf("Deleting from empty tree caused; size=%d", size)
	}
}

func testSearchTreeDeleteExistsAlone(st SearchTree, t *testing.T) {
	obj := NewElement("abc")
	st.Add(obj)

	err := st.Delete(obj.Hash())

	if err != nil {
		t.Errorf("Deleting lone item from tree caused; %v", err)
	}

	size := st.Size()
	if size != 0 {
		t.Errorf("Deleting lone item from tree caused; size=%d", size)
	}

	if st.Contains(obj.Hash()) {
		t.Errorf("Deleted item was still found in tree")
	}
}

func testSearchTreeDeleteExisting(st SearchTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("def")
	st.Add(obj1)
	st.Add(obj2)

	err := st.Delete(obj1.Hash())

	if err != nil {
		t.Errorf("Deleting item from tree caused; %v", err)
	}

	size := st.Size()
	if size != 1 {
		t.Errorf("Deleting item from tree caused; size=%d", size)
	}

	if st.Contains(obj1.Hash()) {
		t.Errorf("Deleted item was still found in tree")
	}

	if !st.Contains(obj2.Hash()) {
		t.Errorf("Expected item was not found in tree")
	}
}

func testSearchTreeDeleteDoesNotExist(st SearchTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("def")
	st.Add(obj1)

	err := st.Delete(obj2.Hash())

	if err == nil {
		t.Errorf("Deleting non-existent item from tree did not error")
	}

	size := st.Size()
	if size != 1 {
		t.Errorf("Deleting from empty tree caused; size=%d", size)
	}

	if !st.Contains(obj1.Hash()) {
		t.Errorf("Expected item was not found in tree")
	}
}

func testSearchTreeRemoveEmpty(st SearchTree, t *testing.T) {
	obj := NewElement("abc")

	item, err := st.Remove(obj.Hash())

	if err == nil {
		t.Errorf("Removing non-existent item from tree did not error")
	}

	if item != nil {
		t.Errorf("Removing non-existent item returned %v", item)
	}

	size := st.Size()
	if size != 0 {
		t.Errorf("Removing from empty tree caused; size=%d", size)
	}
}

func testSearchTreeRemoveExistsAlone(st SearchTree, t *testing.T) {
	obj := NewElement("abc")
	st.Add(obj)

	item, err := st.Remove(obj.Hash())

	if err != nil {
		t.Errorf("Removing lone item from tree caused; %v", err)
	}

	if item.ID() != obj.ID() {
		t.Errorf("Removing item returned unexpected item  %v", item)
	}

	size := st.Size()
	if size != 0 {
		t.Errorf("Removing lone item from tree caused; size=%d", size)
	}

	if st.Contains(obj.Hash()) {
		t.Errorf("Removed item was still found in tree")
	}
}

func testSearchTreeRemoveExisting(st SearchTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("def")
	st.Add(obj1)
	st.Add(obj2)

	item, err := st.Remove(obj1.Hash())

	if err != nil {
		t.Errorf("Removing item from tree caused; %v", err)
	}

	if item.ID() != obj1.ID() {
		t.Errorf("Removing item returned unexpected item  %v", item)
	}

	size := st.Size()
	if size != 1 {
		t.Errorf("Removing item from tree caused; size=%d", size)
	}

	if st.Contains(obj1.Hash()) {
		t.Errorf("Removed item was still found in tree")
	}

	if !st.Contains(obj2.Hash()) {
		t.Errorf("Expected item was not found in tree")
	}
}

func testSearchTreeRemoveDoesNotExist(st SearchTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("def")
	st.Add(obj1)

	item, err := st.Remove(obj2.Hash())

	if err == nil {
		t.Errorf("Removing non-existent item from tree did not error")
	}

	if item != nil {
		t.Errorf("Removing non-existent item returned %v", item)
	}

	size := st.Size()
	if size != 1 {
		t.Errorf("Removing from empty tree caused; size=%d", size)
	}

	if !st.Contains(obj1.Hash()) {
		t.Errorf("Expected item was not found in tree")
	}
}

func testSearchTreeGetEmpty(st SearchTree, t *testing.T) {
	obj := NewElement("abc")

	item, err := st.Get(obj.Hash())

	if err == nil {
		t.Errorf("Getting non-existent item from tree did not error")
	}

	if item != nil {
		t.Errorf("Getting non-existent item returned %v", item)
	}

	size := st.Size()
	if size != 0 {
		t.Errorf("Getting from empty tree caused; size=%d", size)
	}

	if st.Contains(obj.Hash()) {
		t.Errorf("Unexpected item found in tree")
	}
}

func testSearchTreeGetExistsAlone(st SearchTree, t *testing.T) {
	obj := NewElement("abc")
	st.Add(obj)

	item, err := st.Get(obj.Hash())

	if err != nil {
		t.Errorf("Getting lone item from tree caused; %v", err)
	}

	if item.ID() != obj.ID() {
		t.Errorf("Getting item returned unexpected item  %v", item)
	}

	size := st.Size()
	if size != 1 {
		t.Errorf("Getting lone item from tree caused; size=%d", size)
	}

	if !st.Contains(obj.Hash()) {
		t.Errorf("Retrieved item cannot be found in tree")
	}
}

func testSearchTreeGetExisting(st SearchTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("def")
	st.Add(obj1)
	st.Add(obj2)

	item, err := st.Get(obj1.Hash())

	if err != nil {
		t.Errorf("Getting item from tree caused; %v", err)
	}

	if item.ID() != obj1.ID() {
		t.Errorf("Getting item returned unexpected item  %v", item)
	}

	size := st.Size()
	if size != 2 {
		t.Errorf("Getting item from tree caused; size=%d", size)
	}

	if !st.Contains(obj1.Hash()) {
		t.Errorf("Retrieved item cannot be found in tree")
	}

	if !st.Contains(obj2.Hash()) {
		t.Errorf("Expected item was not found in tree")
	}
}

func testSearchTreeGetDoesNotExist(st SearchTree, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("def")
	st.Add(obj1)

	item, err := st.Get(obj2.Hash())

	if err == nil {
		t.Errorf("Getting non-existent item from tree did not error")
	}

	if item != nil {
		t.Errorf("Getting non-existent item returned %v", item)
	}

	size := st.Size()
	if size != 1 {
		t.Errorf("Getting from empty tree caused; size=%d", size)
	}

	if !st.Contains(obj1.Hash()) {
		t.Errorf("Expected item was not found in tree")
	}

	if st.Contains(obj2.Hash()) {
		t.Errorf("Unexpected item found in tree")
	}
}

func testSearchTreeContains(st SearchTree, t *testing.T) {
	items := getTestSearchTreeData()

	for i := 0; i < len(items); i++ {
		st.Add(items[i])

		for j := 0; j <= i; j++ {
			if !st.Contains(items[j].Hash()) {
				t.Errorf("Expected item was not found in tree")
			}
		}

		for j := i + 1; j < len(items); j++ {
			if st.Contains(items[j].Hash()) {
				t.Errorf("Tree contains unexpected item")
			}
		}
	}

	for i := 0; i < len(items); i++ {
		st.Delete(items[i].Hash())

		for j := 0; j <= i; j++ {
			if st.Contains(items[j].Hash()) {
				t.Errorf("Tree contains unexpected item")
			}
		}

		for j := i + 1; j < len(items); j++ {
			if !st.Contains(items[j].Hash()) {
				t.Errorf("Expected item was not found in tree")
			}
		}
	}
}

func testSearchTreeSize(st SearchTree, t *testing.T) {
	items := getTestSearchTreeData()

	for i := 0; i < len(items); i++ {
		if st.Size() != i {
			t.Errorf("Tree with %d items had size %d", i, st.Size())
		}

		st.Add(items[i])
	}

	for i := len(items); i > 0; i-- {
		if st.Size() != i {
			t.Errorf("Tree with %d items had size %d", i, st.Size())
		}

		st.Delete(items[i-1].Hash())
	}

	if st.Size() != 0 {
		t.Errorf("Empty tree had size %d", st.Size())
	}
}
