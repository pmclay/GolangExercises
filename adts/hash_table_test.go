package main

import (
	"testing"
)

func TestHashTableAddEmpty(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testAddEmpty(h, t)
}

func TestHashTableAddNonEmpty(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testAddNonEmpty(h, t)
}

func TestHashTableAddOnExisting(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testAddOnExisting(h, t)
}

func TestHashTableDeleteEmpty(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDeleteEmpty(h, t)
}

func TestHashTableDeleteExistsAlone(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDeleteExistsAlone(h, t)
}

func TestHashTableDeleteExisting(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDeleteExisting(h, t)
}

func TestHashTableDeleteDoesNotExist(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDeleteDoesNotExist(h, t)
}

func TestHashTableGetExisting(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testGetExisting(h, t)
}

func TestHashTableGetDoesNotExist(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testGetDoesNotExist(h, t)
}

func TestHashTableContainsExisting(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testContainsExisting(h, t)
}

func TestHashTableContainsDoesNotExist(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testContainsDoesNotExist(h, t)
}

func TestHashTableSize(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testSize(h, t)
}

func TestHashTableCapacityGrowing(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testCapacityGrowing(h, t)
}

func TestHashTableCapacityShrinking(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testCapacityShrinking(h, t)
}
