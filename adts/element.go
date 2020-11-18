/*
 *  © 2020 Michael Page. All rights reserved.
 */

package main

var idCounter int

type Element struct {
	id   int
	data string
}

func NewElement(data string) *Element {
	idCounter += 1

	e := new(Element)
	e.id = idCounter
	e.data = data

	return e
}

func (e Element) Hash() int {
	hash := 0

	n := len(e.data)
	for i, char := range e.data {
		multiplier := 1
		for j := 0; j < n-(i+1); j++ {
			multiplier *= 31
		}

		hash += int(char) * multiplier
	}

	return hash
}

func (e Element) ID() int {
	return e.id
}

func (e Element) String() string {
	return e.data
}
