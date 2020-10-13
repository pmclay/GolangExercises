package main

import ()

/*
 * A simple min heap implementation.
 */
type BinaryHeap struct {
	data []Item
}

func (h *BinaryHeap) find(key string) int {
	for i, item := range h.data {
		if item.String() == key {
			return i
		}
	}

	return -1
}

func (h *BinaryHeap) heapifyDown(targetIdx int) {
	for {
		leftIdx := (targetIdx * 2) + 1
		rightIdx := (targetIdx * 2) + 2

		if leftIdx >= len(h.data) {
			break
		} else if h.data[targetIdx].String() > h.data[leftIdx].String() {
			if rightIdx >= len(h.data) {
				h.swap(targetIdx, leftIdx)
			} else if h.data[leftIdx].String() < h.data[rightIdx].String() {
				h.swap(targetIdx, leftIdx)
			} else {
				h.swap(targetIdx, rightIdx)
			}
		} else if rightIdx >= len(h.data) {
			break
		} else if h.data[targetIdx].String() > h.data[rightIdx].String() {
			h.swap(targetIdx, rightIdx)
		} else {
			// the heap is in order
			break
		}
	}
}

func (h *BinaryHeap) swap(i, j int) {
	save := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = save
}

func NewBinaryHeap() *BinaryHeap {
	return new(BinaryHeap)
}

func (h *BinaryHeap) Add(obj Item) {
	h.data = append(h.data, obj)

	// restore heap property
	for i := len(h.data) - 1; ; {
		childIdx := i

		i = (i - 1) / 2
		if i < 0 {
			break
		}

		if h.data[childIdx].String() < h.data[i].String() {
			// out of order swap the values
			h.swap(i, childIdx)
		} else {
			break
		}
	}
}

func (h *BinaryHeap) Delete(key string) error {
	targetIdx := h.find(key)

	if targetIdx == -1 {
		return NewDoesNotExistError()
	}

	h.data[targetIdx] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.heapifyDown(targetIdx)

	return nil
}

func (h *BinaryHeap) Remove(key string) (Item, error) {
	targetIdx := h.find(key)

	if targetIdx == -1 {
		return nil, NewDoesNotExistError()
	}

	removed := h.data[targetIdx]
	h.data[targetIdx] = h.data[len(h.data)]
	h.data = h.data[:len(h.data)-1]

	h.heapifyDown(targetIdx)

	return removed, nil
}

func (h *BinaryHeap) Get(key string) (Item, error) {
	targetIdx := h.find(key)

	if targetIdx == -1 {
		return nil, NewDoesNotExistError()
	} else {
		return h.data[targetIdx], nil
	}
}

func (h *BinaryHeap) Pop() (Item, error) {
	if len(h.data) == 0 {
		return nil, NewDoesNotExistError()
	}

	removed := h.data[0]
	h.data[0] = h.data[len(h.data)]
	h.data = h.data[:len(h.data)-1]

	h.heapifyDown(0)

	return removed, nil
}

func (h *BinaryHeap) Peek() (Item, error) {
	if len(h.data) == 0 {
		return nil, NewDoesNotExistError()
	}

	return h.data[0], nil
}

func (h *BinaryHeap) Contains(key string) bool {
	return h.find(key) != -1
}

func (h *BinaryHeap) Size() int {
	return len(h.data)
}
