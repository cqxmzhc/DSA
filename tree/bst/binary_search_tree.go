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

//PrintBST print bst
func PrintBST(root *node) {
	if root == nil {
		return
	}

	fmt.Println(root.value)
	fmt.Println(root.leftChild)
	fmt.Println(root.rightChild)
}
