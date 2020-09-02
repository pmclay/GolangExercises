package main

import (
	"fmt"
	"strings"
)

func bstPostOrderLabel(node *treenode, id *int, out *strings.Builder) string {
	var leftID, rightID string

	if node.left != nil {
		leftID = bstPostOrderLabel(node.left, id, out)
	} else {
		leftID = fmt.Sprintf("n%d", *id)
		(*id)++
		out.WriteString(fmt.Sprintf("%s [label=\"nil\"]\n", leftID))
	}

	if node.right != nil {
		rightID = bstPostOrderLabel(node.right, id, out)
	} else {
		rightID = fmt.Sprintf("n%d", *id)
		(*id)++
		out.WriteString(fmt.Sprintf("%s [label=\"nil\"]\n", rightID))
	}

	myID := fmt.Sprintf("n%d", *id)
	(*id)++

	out.WriteString(fmt.Sprintf("%s [label=\"%s\"]\n",
		myID,
		node.data.String()))

	out.WriteString(fmt.Sprintf("%s -> %s\n", myID, leftID))
	out.WriteString(fmt.Sprintf("%s -> %s\n", myID, rightID))

	return myID
}

func (t *BinarySearchTree) String() string {
	buffer := strings.Builder{}
	buffer.WriteString("digraph BST {")

	if t.root == nil {
		buffer.WriteString("nil}")
	} else {
		id := 0
		bstPostOrderLabel(t.root, &id, &buffer)
		buffer.WriteString("}")
	}

	return buffer.String()
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
