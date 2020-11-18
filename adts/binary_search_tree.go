/*
 *  © 2020 Michael Page. All rights reserved.
 */

package main

import ()

type treenode struct {
	key         int
	data        Item
	left, right *treenode
}

func (n *treenode) find(key int) **treenode {
	if n.left == nil && n.right == nil {
		return nil
	}

	var next *treenode
	var target **treenode = nil

	if key < n.key {
		next = n.left
	} else {
		next = n.right
	}

	if next != nil {
		if next.key == key {
			target = &next
		} else {
			target = next.find(key)
		}
	}

	return target
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

func (n *treenode) delete() *treenode {
	if n.right != nil {
		n.key = n.right.key
		n.data = n.right.data
		n.right = n.right.delete()
	} else if n.left != nil {
		n.key = n.left.key
		n.data = n.left.data
		n.left = n.left.delete()
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

	if t.root.key == key {
		t.root = t.root.delete()
	} else {
		target := t.root.find(key)

		if target == nil {
			return NewDoesNotExistError()
		}

		*target = (*target).delete()
	}

	t.size -= 1

	return nil
}

func (t *BinarySearchTree) Remove(key int) (Item, error) {
	if t.root == nil {
		return nil, NewDoesNotExistError()
	}

	var item Item

	if t.root.key == key {
		item = t.root.data
		t.root = t.root.delete()
	} else {
		target := t.root.find(key)

		if target == nil {
			return nil, NewDoesNotExistError()
		}

		item = (*target).data
		*target = (*target).delete()
	}

	t.size -= 1

	return item, nil
}

func (t *BinarySearchTree) Get(key int) (Item, error) {
	if t.root == nil {
		return nil, NewDoesNotExistError()
	}

	var item Item

	if t.root.key == key {
		item = t.root.data
	} else {
		target := t.root.find(key)

		if target == nil {
			return nil, NewDoesNotExistError()
		}

		item = (*target).data
	}

	return item, nil
}

func (t *BinarySearchTree) Contains(key int) bool {
	if t.root == nil {
		return false
	} else if t.root.key == key {
		return true
	} else {
		return t.root.find(key) != nil
	}
}

func (t *BinarySearchTree) Size() int {
	return t.size
}
