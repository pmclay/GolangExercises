/*
 *  © 2020 Michael Page. All rights reserved.
 */

package main

import (
	"strconv"
	"testing"
)

func getTestDictionaryData() [24]*Element {

	testData := [24]*Element{
		NewElement("bdac"),
		NewElement("abdc"),
		NewElement("cbda"),
		NewElement("adbc"),
		NewElement("dcab"),
		NewElement("acdb"),
		NewElement("dbca"),
		NewElement("adcb"),
		NewElement("bacd"),
		NewElement("acbd"),
		NewElement("bcda"),
		NewElement("bdca"),
		NewElement("cadb"),
		NewElement("dbac"),
		NewElement("cdab"),
		NewElement("badc"),
		NewElement("cabd"),
		NewElement("dcba"),
		NewElement("cdba"),
		NewElement("dabc"),
		NewElement("abcd"),
		NewElement("dacb"),
		NewElement("cbad"),
		NewElement("bcad")}

	return testData
}

// all test cases expect the Dictionary under test
// to be a freshly instantiated object

func testDictionaryAddEmpty(d Dictionary, t *testing.T) {
	obj := NewElement("abc")

	d.Add(obj)

	s := d.Size()
	if s != 1 {
		t.Errorf("Adding one item to dictionary caused; size=%d", s)
	}

	item, err := d.Get(obj.Hash())

	if err != nil {
		t.Errorf("Getting the added item caused an error; %v", err)
	} else if item.String() != obj.String() {
		t.Errorf("Getting the added item returned; %v", item)
	} else {
		// all good
	}
}

func testDictionaryAddNonEmpty(d Dictionary, t *testing.T) {
	obj1 := NewElement("abc")
	d.Add(obj1)

	obj2 := NewElement("def")
	d.Add(obj2)

	s := d.Size()
	if s != 2 {
		t.Errorf("Adding an item to dictionary caused; size=%d", s)
	}

	item, err := d.Get(obj2.Hash())

	if err != nil {
		t.Errorf("Getting the added item caused an error; %v", err)
	} else if item.String() != obj2.String() {
		t.Errorf("Getting the added item returned; %v", item)
	} else {
		// all good
	}

	item, err = d.Get(obj1.Hash())

	if err != nil {
		t.Errorf("Getting the pre-existing item caused an error; %v", err)
	} else if item.String() != obj1.String() {
		t.Errorf("Getting the pre-existing item returned; %v", item)
	} else {
		// all good
	}
}

func testDictionaryAddOnExisting(d Dictionary, t *testing.T) {
	obj1 := NewElement("abc")
	d.Add(obj1)

	obj2 := NewElement("abc")
	d.Add(obj2)

	s := d.Size()
	if s != 1 {
		t.Errorf("Adding an item to dictionary caused; size=%d", s)
	}

	item, err := d.Get(obj2.Hash())

	if err != nil {
		t.Errorf("Getting the added item caused an error; %v", err)
	} else if item.ID() != obj2.ID() {
		t.Errorf("New item with identical key didn't replace existing one; %v", item)
	} else {
		// all good
	}

	item, err = d.Get(obj1.Hash())

	if err != nil {
		t.Errorf("Getting the added item caused an error; %v", err)
	} else if item.ID() != obj2.ID() {
		t.Errorf("New item with identical key didn't replace existing one; %v", item)
	} else {
		// all good
	}
}

func testDictionaryDeleteEmpty(d Dictionary, t *testing.T) {
	obj := NewElement("abc")

	err := d.Delete(obj.Hash())

	if err == nil {
		t.Errorf("Deleting non-existent item didn't cause an error")
	}

	s := d.Size()
	if s != 0 {
		t.Errorf("Deleting from empty dictionary caused; size=%d", s)
	}
}

func testDictionaryDeleteExistsAlone(d Dictionary, t *testing.T) {
	obj := NewElement("abc")
	d.Add(obj)

	err := d.Delete(obj.Hash())

	if err != nil {
		t.Errorf("Deleting item caused an error; %v", err)
	}

	s := d.Size()
	if s != 0 {
		t.Errorf("Deleting item caused; size=%d", s)
	}

	item, err := d.Get(obj.Hash())

	if err == nil {
		t.Errorf("Getting deleted item didn't cause an error")
	}

	if item != nil {
		t.Errorf("Getting deleted item returned; %v", item)
	}
}

func testDictionaryDeleteExisting(d Dictionary, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("def")
	d.Add(obj1)
	d.Add(obj2)

	err := d.Delete(obj2.Hash())

	if err != nil {
		t.Errorf("Deleting item caused an error; %v", err)
	}

	s := d.Size()
	if s != 1 {
		t.Errorf("Deleting item caused; size=%d", s)
	}

	err = d.Delete(obj2.Hash())

	if err == nil {
		t.Errorf("Deleting non-existent item didn't cause an error")
	}

	item, err := d.Get(obj1.Hash())

	if err != nil {
		t.Errorf("Getting pre-existing item errored after deleting new item; %v", err)
	} else if item.ID() != obj1.ID() {
		t.Errorf("Deleting new item mutated pre-existing item; %v", item)
	} else {
		// all good
	}
}

func testDictionaryDeleteDoesNotExist(d Dictionary, t *testing.T) {
	obj1 := NewElement("abc")
	obj2 := NewElement("def")
	d.Add(obj1)

	err := d.Delete(obj2.Hash())

	if err == nil {
		t.Errorf("Deleting non-existent item didn't cause an error")
	}

	s := d.Size()
	if s != 1 {
		t.Errorf("Deleting non-existent item caused; size=%d", s)
	}

	item, err := d.Get(obj1.Hash())

	if err != nil {
		t.Errorf("Getting pre-existing item errored after failed delete; %v", err)
	} else if item.ID() != obj1.ID() {
		t.Errorf("Deleting non-existent item mutated pre-existing item; %v", item)
	} else {
		// all good
	}
}

func testDictionaryGetExisting(d Dictionary, t *testing.T) {
	obj := NewElement("abc")
	d.Add(obj)

	item, err := d.Get(obj.Hash())

	if err != nil {
		t.Errorf("Getting existing item errored; %v", err)
	} else if item.ID() != obj.ID() {
		t.Errorf("Getting existing item returned; %v", item)
	} else {
		// all good
	}
}

func testDictionaryGetDoesNotExist(d Dictionary, t *testing.T) {
	obj := NewElement("abc")

	item, err := d.Get(obj.Hash())

	if err == nil {
		t.Errorf("Getting non-existent item didn't cause an error")
	}

	if item != nil {
		t.Errorf("Getting non-existent item returned; %v", item)
	}
}

func testDictionaryContainsExisting(d Dictionary, t *testing.T) {
	obj := NewElement("abc")
	d.Add(obj)

	hasIt := d.Contains(obj.Hash())

	if !hasIt {
		t.Errorf("Contains reported %v for existing item", hasIt)
	}
}

func testDictionaryContainsDoesNotExist(d Dictionary, t *testing.T) {
	obj := NewElement("abc")

	hasIt := d.Contains(obj.Hash())

	if hasIt {
		t.Errorf("Contains reported %v for non-existing item", hasIt)
	}
}

func testDictionarySize(d Dictionary, t *testing.T) {
	testItems := [30]*Element{nil}

	i := 0
	for ; i < cap(testItems); i++ {
		testItems[i] = NewElement(strconv.Itoa(i))
	}

	for i = 0; i < 30; i++ {
		if d.Size() != i {
			t.Errorf("Dictionary with %d items has size=%d", i, d.Size())
		}

		d.Add(testItems[i])
	}

	for ; i > 0; i-- {
		if d.Size() != i {
			t.Errorf("Dictionary with %d items has size=%d", i, d.Size())
		}

		d.Delete(testItems[i-1].Hash())
	}

	if d.Size() != i {
		t.Errorf("Dictionary with %d items has size=%d", i, d.Size())
	}
}

func testDictionaryCapacityGrowing(d Dictionary, t *testing.T) {
	expectedCapacities := map[int]int{0: 8,
		1:  8,
		2:  8,
		3:  8,
		4:  8,
		5:  8,
		6:  8,
		7:  16,
		8:  16,
		9:  16,
		10: 16,
		11: 16,
		12: 16,
		13: 32,
		14: 32,
		15: 32,
		16: 32,
		17: 32,
		18: 32,
		19: 32,
		20: 32,
		21: 32,
		22: 32,
		23: 32,
		24: 64}

	testData := getTestDictionaryData()
	for i, value := range testData {
		expected := expectedCapacities[i]
		actual := d.Capacity()

		if actual != expected {
			t.Errorf("Dictionary of size=%d has capacity %d not %d", i, actual, expected)
		}

		d.Add(value)
	}

	for _, value := range testData {
		// make sure all the items are really there
		if !d.Contains(value.Hash()) {
			t.Errorf("item %s was not added to dictionary", value.String())
		}
	}

	total := len(testData)
	expected := expectedCapacities[total]
	actual := d.Capacity()
	if actual != expected {
		t.Errorf("Dictionary of size=%d has capacity %d not %d", total, actual, expected)
	}
}

func testDictionaryCapacityShrinking(d Dictionary, t *testing.T) {
	testData := getTestDictionaryData()
	for _, value := range testData {
		d.Add(value)
	}

	expectedCapacities := map[int]int{0: 8,
		1:  8,
		2:  8,
		3:  9,
		4:  9,
		5:  15,
		6:  15,
		7:  15,
		8:  24,
		9:  24,
		10: 24,
		11: 24,
		12: 39,
		13: 39,
		14: 39,
		15: 39,
		16: 39,
		17: 39,
		18: 39,
		19: 39,
		20: 64,
		21: 64,
		22: 64,
		23: 64,
		24: 64}

	for i := len(testData) - 1; i >= 0; i-- {
		expected := expectedCapacities[i+1]
		actual := d.Capacity()

		if expected != actual {
			t.Errorf("Dictionary of size=%d has capacity %d not %d", i, actual, expected)
		}

		if err := d.Delete(testData[i].Hash()); err != nil {
			t.Errorf("Deleting existing item from dictionary caused; %v", err)
		}

		for j := 0; j < i; j++ {
			// make sure all remaining items are really there
			if !d.Contains(testData[j].Hash()) {
				t.Errorf("item %s was lost", testData[j].String())
			}
		}
	}

	expected := expectedCapacities[0]
	actual := d.Capacity()
	if actual != expected {
		t.Errorf("Dictionary of size=%d has capacity %d not %d", 0, actual, expected)
	}
}
