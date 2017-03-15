package main

import (
	"fmt"
)

type Tree struct {
	value       int
	left, right *Tree
}

func add(root *Tree, val int) *Tree {
	if root == nil {
		root = new(Tree)
		root.value = val
		return root
	}
	if val < root.value {
		root.left = add(root.left, val)
	} else {
		root.right = add(root.right, val)
	}
	return root
}

func appendVal(values []int, t *Tree) []int {
	if t != nil {
		//中序遍历
		values = appendVal(values, t.left)
		values = append(values, t.value)
		values = appendVal(values, t.right)
	}
	return values
}

func Sort(values []int) {
	var root *Tree
	for _, val := range values {
		root = add(root, val)
	}
	appendVal(values[:0], root)
}

func main() {
	values := []int{1, 3, 22, 4, 1, 2}
	Sort(values[:])
	fmt.Println(values)
}
