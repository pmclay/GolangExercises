package main

import ()

type treenode struct {
	key         int
	data        Item
	left, right *treenode
}

func (n *treenode) find(key int) *treenode {
	if n == nil {
		return nil
	}

	if key == n.key {
		return n
	} else if key < n.key {
		return n.left.find(key)
	} else {
		return n.right.find(key)
	}
}

func (n *treenode) add(obj Item) {
	key := obj.Hash()

	if key < n.key {
		if n.left == nil {
			n.left = new(treenode)
			n.left.key = key
			n.left.data = obj
		} else {
			n.left.add(obj)
		}
	} else {
		if n.right == nil {
			n.right = new(treenode)
			n.right.key = key
			n.right.data = obj
		} else {
			n.right.add(obj)
		}
	}
}

func (n *treenode) delete(key int) *treenode {
	if n.right != nil {
		n.key = n.right.key
		n.data = n.right.data
		n.right = n.right.delete(key)
	} else if n.left != nil {
		n.key = n.left.key
		n.data = n.left.data
		n.left = n.left.delete(key)
	} else {
		return nil
	}

	return n
}

type BinarySearchTree struct {
	root *treenode
	size int
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

func (t *BinarySearchTree) Add(obj Item) {
	if t.root == nil {
		t.root = new(treenode)
		t.root.key = obj.Hash()
		t.root.data = obj
	} else {
		t.root.add(obj)
	}

	t.size += 1
}

func (t *BinarySearchTree) Delete(key int) error {
	if t.root == nil {
		return NewDoesNotExistError()
	}

	target := t.root.find(key)

	if target == t.root {
		t.root = nil
	}

	if target != nil {
		target.delete(key)
		t.size -= 1

		return nil
	} else {
		return NewDoesNotExistError()
	}
}

func (t *BinarySearchTree) Remove(key int) (Item, error) {
	t.size -= 1
	return nil, nil
}

func (t *BinarySearchTree) Get(key int) (Item, error) {
	return nil, nil
}

func (t *BinarySearchTree) Contains(key int) bool {
	return t.root.find(key) != nil
}

func (t *BinarySearchTree) Size() int {
	return t.size
}
