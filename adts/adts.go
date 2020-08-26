package main

import (
	"fmt"
)

func NewOutOfBoundsError(idx int) error {
	return fmt.Errorf("index %d is out of bounds", idx)
}

func NewIsEmptyError() error {
	return fmt.Errorf("data structure is empty")
}

func NewDoesNotExistError() error {
	return fmt.Errorf("no such item exists")
}

type Item interface {
	Hash() int
	ID() int
	String() string
}

type List interface {
	Prepend(obj Item)
	Append(obj Item)
	Insert(obj Item, idx int) error
	Get(idx int) (Item, error)
	Set(obj Item, idx int) error
	RemoveFirst() (Item, error)
	RemoveLast() (Item, error)
	Remove(idx int) (Item, error)
	Size() int
	Capacity() int
	String() string
}

type Dictionary interface {
	Add(obj Item)
	Delete(key int) error
	Get(key int) (Item, error)
	Contains(key int) bool
	Size() int
	Capacity() int
}

type SearchTree interface {
	Add(obj Item)
	Delete(key int) error
	Remove(key int) (Item, error)
	Get(key int) (Item, error)
	Contains(key int) bool
	Size() int
}

type PrefixTree interface {
	Add(obj Item)
	Contains(key string) bool
	Size() int
}
