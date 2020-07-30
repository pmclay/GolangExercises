package main

import (
	"math"
)

const minLoad = 0.3
const maxLoad = 0.7
const shrinkFactor = 0.6
const growthFactor = 2.0

type entry struct {
	key   int
	value Item
	next  *entry
}

type HashTable struct {
	table          []*entry
	capacity, size int
}

func NewHashTable(initialCapacity int) *HashTable {
	if initialCapacity < 8 {
		initialCapacity = 8
	}

	h := new(HashTable)
	h.table = make([]*entry, initialCapacity)
	h.capacity = initialCapacity
	h.size = 0

	return h
}

func (h *HashTable) get_load_factor() float32 {
	return float32(h.size) / float32(h.capacity)
}

func (h *HashTable) resize_rehash(grow bool) {
	var newCapacity int
	if grow {
		newCapacity = int(math.Ceil(float64(h.capacity) * growthFactor))
	} else {
		newCapacity = int(math.Ceil(float64(h.capacity) * shrinkFactor))

		if newCapacity < 8 {
			newCapacity = 8
		}
	}

	newTable := make([]*entry, newCapacity)

	for i := 0; i < h.capacity; i++ {
		current := h.table[i]

		for current != nil {
			h.table[i] = current.next
			current.next = nil
			idx := current.key % newCapacity

			if newTable[idx] == nil {
				newTable[idx] = current
			} else {
				tail := newTable[idx]

				for tail.next != nil {
					tail = tail.next
				}

				tail.next = current
			}

			current = h.table[i]
		}
	}

	h.table = newTable
	h.capacity = newCapacity
}

func (h *HashTable) Add(obj Item) {
	if h.get_load_factor() > maxLoad {
		h.resize_rehash(true)
	}

	key := obj.Hash()
	idx := key % h.capacity

	if h.table[idx] == nil {
		h.table[idx] = &entry{key, obj, nil}
	} else {
		tail := h.table[idx]

		for {
			if tail.key == key {
				tail.value = obj
				return
			}

			if tail.next != nil {
				tail = tail.next
			} else {
				break
			}
		}

		tail.next = &entry{key, obj, nil}
	}

	h.size += 1
}

func (h *HashTable) Delete(key int) error {
	idx := key % h.capacity

	var previous *entry = nil
	current := h.table[idx]

	found := false
	for current != nil {
		if current.key == key {
			if previous != nil {
				previous.next = current.next
			} else {
				h.table[idx] = current.next
			}

			current.next = nil
			found = true
			break
		}

		previous = current
		current = current.next
	}

	if found {
		h.size -= 1

		if (h.capacity > 8) && (h.get_load_factor() < minLoad) {
			h.resize_rehash(false)
		}

		return nil
	} else {
		return NewDoesNotExistError()
	}
}

func (h *HashTable) Get(key int) (Item, error) {
	idx := key % h.capacity
	current := h.table[idx]

	var obj Item = nil
	for current != nil {
		if current.key == key {
			obj = current.value
			break
		}

		current = current.next
	}

	if obj != nil {
		return obj, nil
	} else {
		return nil, NewDoesNotExistError()
	}
}

func (h *HashTable) Contains(key int) bool {
	idx := key % h.capacity
	current := h.table[idx]

	found := false
	for current != nil {
		if current.key == key {
			found = true
			break
		}

		current = current.next
	}

	return found
}

func (h *HashTable) Size() int {
	return h.size
}

func (h *HashTable) Capacity() int {
	return h.capacity
}
