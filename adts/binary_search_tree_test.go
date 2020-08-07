package main

import (
	"testing"
)

func TestBinarySearchTreeAddEmpty(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeAddEmpty(st, t)
}

func TestBinarySearchTreeAddNonEmpty(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeAddNonEmpty(st, t)
}

func TestBinarySearchTreeDeleteEmpty(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeDeleteEmpty(st, t)
}

func TestBinarySearchTreeDeleteExistsAlone(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeDeleteExistsAlone(st, t)
}

func TestBinarySearchTreeDeleteExisting(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeDeleteExisting(st, t)
}

func TestBinarySearchTreeDeleteDoesNotExist(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeDeleteDoesNotExist(st, t)
}

func TestBinarySearchTreeRemoveEmpty(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeRemoveEmpty(st, t)
}

func TestBinarySearchTreeRemoveExistsAlone(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeRemoveExistsAlone(st, t)
}

func TestBinarySearchTreeRemoveExisting(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeRemoveExisting(st, t)
}

func TestBinarySearchTreeRemoveDoesNotExist(t *testing.T) {
	var st SearchTree = NewBinarySearchTree()
	testSearchTreeRemoveDoesNotExist(st, t)
}
