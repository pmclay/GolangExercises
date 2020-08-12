package main

import (
	"fmt"
)

func testDA() {
	var x List
	x = NewDynamicArray(2)

	x.Append(NewElement("1"))
	printList(x)

	x.Append(NewElement("2"))
	printList(x)

	x.Append(NewElement("3"))
	printList(x)

	x.Append(NewElement("4"))
	printList(x)

	x.Remove(2)
	printList(x)

	x.Prepend(NewElement("6"))
	printList(x)

	y, err := x.RemoveLast()
	fmt.Printf("%v %v\n", y, err)
	printList(x)
}

func printList(l List) {
	fmt.Printf("%s  size=%d\n", l, l.Size())
}

func testLL() {
	var x List
	x = NewLinkedList()

	x.Append(NewElement("7"))

	x.Append(NewElement("8"))
	printList(x)

	x.Prepend(NewElement("9"))
	printList(x)

	x.Remove(0)
	printList(x)

	x.Insert(NewElement("4"), 1)
	printList(x)
}

func testHT() {
	ht := NewHashTable(8)
	data := getTestDictionaryData()

	i := 0
	for ; ; i++ {
		fmt.Printf("%d items added to table...\n", i)
		fmt.Printf("table size=%d\n", ht.Size())
		ht.PrintTable()
		fmt.Println()

		if i == 6 {
			break
		}
		ht.Add(data[i])
	}

	ht.Add(data[i])
	i++
	fmt.Printf("table size=%d\n", ht.Size())
	ht.PrintTable()
	fmt.Println()
}

func testBST() {
	bst := NewBinarySearchTree()
	data := getTestDictionaryData()

	for i := 0; i < 7; i++ {
		bst.Add(data[i])
	}
	bst.Print()
	fmt.Printf("size=%d", bst.Size())
}

func main() {
	testBST()
	/*fmt.Println("Playing with DynamicArray")
	testHT()
	testDA()
	fmt.Println()
	fmt.Println("Playing with LinkedList")
	testLL()*/
}
