// Package bst provides ...
package bst

import (
	"fmt"
)

type BSTree struct {
	root *node
}

type node struct {
	value      int
	leftChild  *node
	rightChild *node
}

//New return new node
func NewTree(initial int) BSTree {
	return BSTree{
		root: NewNode(initial),
	}
}

func NewNode(initial int) *node {
	return &node{
		value:      initial,
		leftChild:  nil,
		rightChild: nil,
	}

}

//Insert insert value
func Insert(t BSTree, value int) bool {
	return insert(t.root, value)
}

// insert
func insert(root *node, value int) bool {
	if root == nil {
		return false
	}

	if value <= root.value {
		if root.leftChild == nil {
			root.leftChild = NewNode(value)
			return true
		} else {
			return insert(root.leftChild, value)
		}
	} else {
		if root.rightChild == nil {
			root.rightChild = NewNode(value)
			return true
		} else {
			return insert(root.rightChild, value)
		}
	}
}

//Delete delete value from tree
func Delete(t BSTree, value int) bool {
	return delete(t.root, value)
}

//delete
func delete(root *node, value int) bool {
	if root == nil {
		return false
	}

	if root.value == value {
		if root.leftChild == nil && root.rightChild == nil {
			root = nil
			return true
		}

		if root.leftChild == nil && root.rightChild != nil {
			root.value = root.rightChild.value
			root.leftChild = root.rightChild.leftChild
			root.rightChild = root.rightChild.rightChild

			return true
		}

		if root.leftChild != nil && root.rightChild == nil {
			root.value = root.leftChild.value
			root.rightChild = root.leftChild.rightChild
			root.leftChild = root.leftChild.leftChild

			return true
		}

		if root.leftChild != nil && root.rightChild != nil {
			smallest, parent := findSmallest(root.rightChild, root)
			parent.rightChild = nil
			root.value = smallest.value

		}
	} else if root.value > value {
		return delete(root.leftChild, value)
	} else {
		return delete(root.rightChild, value)
	}

	return false
}

//Find find value
func Find(t BSTree, value int) bool {
	return find(t.root, value)
}

//find
func find(root *node, value int) bool {
	if root == nil {
		return false
	}

	if root.value == value {
		return true
	} else if root.value > value {
		return find(root.leftChild, value)
	} else {
		return find(root.rightChild, value)
	}
}

//find the smallest value larger than root.value
func findSmallest(root, parent *node) (*node, *node) {
	if root == nil || root.leftChild == nil {
		return root, parent
	}

	return findSmallest(root.leftChild, root)
}

//PrintBST print bst
func PrintBST(t BSTree) {
	printBST(t.root)
}

//printBST
func printBST(root *node) {
	if root == nil {
		return
	}

	if root.leftChild != nil {
		printBST(root.leftChild)
	}

	fmt.Println(root.value)

	if root.rightChild != nil {
		printBST(root.rightChild)
	}
}
