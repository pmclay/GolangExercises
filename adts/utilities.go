package main

import (
	"fmt"
	"math"
)

func getDepth(node *treenode, depth int) int {
	if node == nil {
		return depth
	}

	leftDepth := getDepth(node.left, depth+1)
	rightDepth := getDepth(node.right, depth+1)

	if leftDepth < rightDepth {
		return rightDepth
	} else {
		return leftDepth
	}
}

func (t *BinarySearchTree) Print() {
	depth := getDepth(t.root, 0)
	fmt.Printf("depth=%d\n", depth)

	dodo := &treenode{}
	q := []*treenode{t.root, dodo}

	level := 0
	for len(q) != 0 {
		current := q[0]
		q = q[1:]

		if current == dodo {
			fmt.Printf("\n")
			for i := 0; i < int(math.Pow(2.0, float64(depth-level)))-1; i++ {
				fmt.Printf(" ")
			}

			if len(q) != 0 {
				q = append(q, dodo)
				level++
			}
		} else {
			offset := int(math.Floor((math.Pow(2.0, float64(depth+1)) - 1) / math.Pow(2.0, float64(level))))
			for i := 0; i < offset; i++ {
				fmt.Printf(" ")
			}

			if current != nil {
				fmt.Printf(current.data.String())
				q = append(q, current.left)
				q = append(q, current.right)
			} else {
				fmt.Printf("null")
			}

			for i := 0; i < (depth-level)+1; i++ {
				fmt.Printf(" ")
			}
		}
	}
}

func (h *HashTable) PrintTable() {
	for i, item := range h.table {
		fmt.Printf("%d: ", i)
		current := item

		for current != nil {
			fmt.Printf("{key=%d value=%s} -> ", current.key, current.value.String())
			current = current.next
		}

		fmt.Printf("nil\n")
	}
}
