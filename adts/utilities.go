package main

import (
	"fmt"
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
		return rightDepth
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

			if len(q) != 0 {
				q = append(q, dodo)
				level++
			}
		} else {
			for i := 0; i < (depth - level); i++ {
				fmt.Printf(" ")
			}

			if current != nil {
				fmt.Printf(current.data.String())
				q = append(q, current.left)
				q = append(q, current.right)
			} else {
				fmt.Printf("nil")
			}

			for i := 0; i < (depth-level)+1; i++ {
				fmt.Printf(" ")
			}
		}
	}
}
