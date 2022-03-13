/*
 *  © 2020 Michael Page. All rights reserved.
 */

package main

import (
	"bytes"
	"fmt"
)

type listnode struct {
	data       Item
	prev, next *listnode
}

type LinkedList struct {
	head, tail *listnode
	size       int
}

func NewLinkedList() *LinkedList {
	l := new(LinkedList)
	l.head = nil
	l.tail = nil
	l.size = 0

	return l
}

func AsSlice(l *LinkedList) []Item {
	output := make([]Item, l.Size())

	current := l.head
	for current != nil {
		output = append(output, current.data)
		current = current.next
	}
	return output
}

func (l *LinkedList) createFirstNode(obj Item) {
	l.head = new(listnode)
	l.tail = l.head

	l.head.data = obj
	l.head.prev = nil
	l.head.next = nil
}

func (l *LinkedList) Prepend(obj Item) {
	if l.head == nil {
		l.createFirstNode(obj)
	} else {
		l.head.prev = new(listnode)
		l.head.prev.data = obj
		l.head.prev.prev = nil
		l.head.prev.next = l.head

		l.head = l.head.prev
	}

	l.size += 1
}

func (l *LinkedList) Append(obj Item) {
	if l.head == nil {
		l.createFirstNode(obj)
	} else {
		l.tail.next = new(listnode)
		l.tail.next.data = obj
		l.tail.next.prev = l.tail
		l.tail.next.next = nil

		l.tail = l.tail.next
	}

	l.size += 1
}

func (l *LinkedList) Insert(obj Item, idx int) error {
	if idx < 0 || idx > l.size {
		return NewOutOfBoundsError(idx)
	}

	if idx == 0 {
		l.Prepend(obj)
		return nil
	}

	if idx == l.size {
		l.Append(obj)
		return nil
	}

	current := l.head
	for idx > 0 {
		current = current.next
		idx--
	}

	newCurrent := new(listnode)
	newCurrent.data = obj

	current.prev.next = newCurrent
	newCurrent.prev = current.prev
	current.prev = newCurrent
	newCurrent.next = current

	l.size += 1

	return nil
}

func (l *LinkedList) Get(idx int) (Item, error) {
	if idx < 0 || idx >= l.size {
		return nil, NewOutOfBoundsError(idx)
	}

	current := l.head
	for idx > 0 {
		current = current.next
		idx--
	}

	return current.data, nil
}

func (l *LinkedList) Set(obj Item, idx int) error {
	if idx < 0 || idx >= l.size {
		return NewOutOfBoundsError(idx)
	}

	current := l.head
	for idx > 0 {
		current = current.next
		idx--
	}
	current.data = obj

	return nil
}

func (l *LinkedList) RemoveFirst() (Item, error) {
	if l.size == 0 {
		return nil, NewIsEmptyError()
	}

	e := l.head.data

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev.next = nil
		l.head.prev = nil
	}

	l.size -= 1

	return e, nil
}

func (l *LinkedList) RemoveLast() (Item, error) {
	if l.size == 0 {
		return nil, NewIsEmptyError()
	}

	e := l.tail.data

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next.prev = nil
		l.tail.next = nil
	}

	l.size -= 1

	return e, nil
}

func (l *LinkedList) Remove(idx int) (Item, error) {
	if idx < 0 || idx >= l.size {
		return nil, NewOutOfBoundsError(idx)
	}

	if idx == 0 {
		return l.RemoveFirst()
	}

	if idx == l.size-1 {
		return l.RemoveLast()
	}

	current := l.head
	for idx > 0 {
		current = current.next
		idx--
	}

	e := current.data

	current.prev.next = current.next
	current.next.prev = current.prev
	current.prev = nil
	current.next = nil

	l.size -= 1

	return e, nil
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Capacity() int {
	return l.size
}

func (l *LinkedList) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("[")

	current := l.head
	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v", (*current).data))
		current = current.next

		if current != nil {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()
}
