package main

import (
	"testing"
)

func TestHashTableAddEmpty(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryAddEmpty(h, t)
}

func TestHashTableAddNonEmpty(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryAddNonEmpty(h, t)
}

func TestHashTableAddOnExisting(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryAddOnExisting(h, t)
}

func TestHashTableDeleteEmpty(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryDeleteEmpty(h, t)
}

func TestHashTableDeleteExistsAlone(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryDeleteExistsAlone(h, t)
}

func TestHashTableDeleteExisting(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryDeleteExisting(h, t)
}

func TestHashTableDeleteDoesNotExist(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryDeleteDoesNotExist(h, t)
}

func TestHashTableGetExisting(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryGetExisting(h, t)
}

func TestHashTableGetDoesNotExist(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryGetDoesNotExist(h, t)
}

func TestHashTableContainsExisting(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryContainsExisting(h, t)
}

func TestHashTableContainsDoesNotExist(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryContainsDoesNotExist(h, t)
}

func TestHashTableSize(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionarySize(h, t)
}

func TestHashTableCapacityGrowing(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryCapacityGrowing(h, t)
}

func TestHashTableCapacityShrinking(t *testing.T) {
	var h Dictionary = NewHashTable(1)
	testDictionaryCapacityShrinking(h, t)
}
