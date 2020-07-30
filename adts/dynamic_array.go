package main

import (
	"bytes"
	"fmt"
)

// An implementation of a dynamic array which
// ignores the capabilities Go's slice type
// for the sake of academic exercise.
type DynamicArray struct {
	// The dynamic array behavior is implemented
	// via a slice that is managed as a circular array
	data           []Item
	head, tail     int
	size, capacity int
}

func NewDynamicArray(initialCapacity int) *DynamicArray {
	if initialCapacity < 2 {
		initialCapacity = 8
	}

	a := new(DynamicArray)
	a.data = make([]Item, initialCapacity)
	a.head = initialCapacity / 2
	a.tail = a.head - 1
	a.size = 0
	a.capacity = initialCapacity

	return a
}

// expects 'a' is not empty and is at full capacity
func (a *DynamicArray) expand() {
	a.capacity *= 2
	newHead := a.capacity / 4
	newData := make([]Item, a.capacity)

	i := a.head
	for copies := 0; copies < a.size; copies++ {
		if i == cap(a.data) {
			i = 0
		}

		newData[newHead+copies] = a.data[i]
		i++
	}

	a.data = newData
	a.head = newHead
	a.tail = newHead + a.size - 1
}

func (a *DynamicArray) Prepend(obj Item) {
	if a.size == a.capacity {
		a.expand()
	}

	if a.head == 0 {
		a.head = cap(a.data) - 1
	} else {
		a.head -= 1
	}

	a.data[a.head] = obj
	a.size += 1
}

func (a *DynamicArray) Append(obj Item) {
	if a.size == a.capacity {
		a.expand()
	}

	if a.tail == (cap(a.data) - 1) {
		a.tail = 0
	} else {
		a.tail += 1
	}

	a.data[a.tail] = obj
	a.size += 1
}

func (a *DynamicArray) Insert(obj Item, idx int) error {
	if idx < 0 || idx > a.size {
		return NewOutOfBoundsError(idx)
	}

	if a.size == a.capacity {
		a.expand()
	}

	if idx < (a.size / 2) {
		if a.head == 0 {
			a.head = cap(a.data) - 1
		} else {
			a.head -= 1
		}

		i := a.head + 1
		for numShifts := idx + 1; numShifts > 0; numShifts-- {
			if i == cap(a.data) {
				i = 0
				a.data[cap(a.data)-1] = a.data[i]
			} else {
				a.data[i-1] = a.data[i]
			}

			i++
		}

		a.data[i-1] = obj
	} else {
		if a.tail == (cap(a.data) - 1) {
			a.tail = 0
		} else {
			a.tail += 1
		}

		i := a.tail - 1
		for numShifts := a.size - idx; numShifts > 0; numShifts-- {
			if i == -1 {
				i = cap(a.data) - 1
				a.data[0] = a.data[i]
			} else {
				a.data[i+1] = a.data[i]
			}

			i--
		}

		a.data[i+1] = obj
	}
	a.size += 1

	return nil
}

func (a *DynamicArray) Get(idx int) (Item, error) {
	if a.size == 0 || idx < 0 || idx >= a.size {
		return nil, NewOutOfBoundsError(idx)
	}

	return a.data[(a.head+idx)%cap(a.data)], nil
}

func (a *DynamicArray) Set(obj Item, idx int) error {
	if a.size == 0 || idx < 0 || idx >= a.size {
		return NewOutOfBoundsError(idx)
	}

	a.data[(a.head+idx)%cap(a.data)] = obj

	return nil
}

func (a *DynamicArray) RemoveFirst() (Item, error) {
	if a.size == 0 {
		return nil, NewIsEmptyError()
	}

	e := a.data[a.head]

	if a.head == (cap(a.data) - 1) {
		a.head = 0
	} else {
		a.head += 1
	}
	a.size -= 1

	return e, nil
}

func (a *DynamicArray) RemoveLast() (Item, error) {
	if a.size == 0 {
		return nil, NewIsEmptyError()
	}

	e := a.data[a.tail]

	if a.tail == 0 {
		a.tail = cap(a.data) - 1
	} else {
		a.tail -= 1
	}
	a.size -= 1

	return e, nil
}

func (a *DynamicArray) Remove(idx int) (Item, error) {
	if a.size == 0 {
		return nil, NewIsEmptyError()
	}

	if idx < 0 || idx >= a.size {
		return nil, NewOutOfBoundsError(idx)
	}

	realIdx := (a.head + idx) % cap(a.data)

	e := a.data[realIdx]

	if idx < (a.size / 2) {
		i := realIdx - 1
		for numShifts := idx; numShifts > 0; numShifts-- {
			if i == -1 {
				i = cap(a.data) - 1
				a.data[0] = a.data[i]
			} else {
				a.data[i+1] = a.data[i]
			}

			i--
		}

		if a.head == (cap(a.data) - 1) {
			a.head = 0
		} else {
			a.head += 1
		}

	} else {
		i := realIdx + 1
		for numShifts := a.size - idx - 1; numShifts > 0; numShifts-- {
			if i == cap(a.data) {
				i = 0
				a.data[cap(a.data)-1] = a.data[i]
			} else {
				a.data[i-1] = a.data[i]
			}

			i++
		}

		if a.tail == 0 {
			a.tail = cap(a.data) - 1
		} else {
			a.tail -= 1
		}

	}
	a.size -= 1

	return e, nil
}

func (a *DynamicArray) Size() int {
	return a.size
}

func (a *DynamicArray) Capacity() int {
	return a.capacity
}

func (a *DynamicArray) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("[")

	i := a.head
	for count := 0; count < a.size; count++ {
		if i == cap(a.data) {
			i = 0
		}

		buffer.WriteString(fmt.Sprintf("%v", a.data[i]))
		i++

		if count != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()
}
