// Package bst provides ...
package bst

import (
	"fmt"
)

type node struct {
	value      int
	leftChild  *node
	rightChild *node
}

//New return new node
func New(initial int) *node {
	return &node{
		value:      initial,
		leftChild:  nil,
		rightChild: nil,
	}
}

//Insert insert value
func Insert(root *node, value int) bool {
	if root == nil {
		return false
	}

	if value <= root.value {
		if root.leftChild == nil {
			root.leftChild = New(value)
			return true
		} else {
			return Insert(root.leftChild, value)
		}
	} else {
		if root.rightChild == nil {
			root.rightChild = New(value)
			return true
		} else {
			return Insert(root.rightChild, value)
		}
	}
}

//Delete delete value from tree
func Delete(root *node, value int) bool {
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
	} else if root.value > value {
		return Delete(root.leftChild, value)
	} else {
		return Delete(root.rightChild, value)
	}

	return false
}

//Find find value
func Find(root *node, value int) bool {
	if root == nil {
		return false
	}

	if root.value == value {
		return true
	} else if root.value > value {
		return Find(root.leftChild, value)
	} else {
		return Find(root.rightChild, value)
	}
}

//PrintBST print bst
func PrintBST(root *node) {
	if root == nil {
		return
	}

	if root.leftChild != nil {
		PrintBST(root.leftChild)
	}

	fmt.Println(root.value)

	if root.rightChild != nil {
		PrintBST(root.rightChild)
	}
}
