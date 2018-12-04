package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// Pre-order go through binary tree
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.val)
	preOrder(node.left)
	preOrder(node.right)
}

func buildTree() *TreeNode {
	n1 := &TreeNode{val: 1}
	n2 := &TreeNode{val: 2}
	n3 := &TreeNode{val: 3}
	n4 := &TreeNode{val: 4}
	n5 := &TreeNode{val: 5}
	n6 := &TreeNode{val: 6}
	n7 := &TreeNode{val: 7}

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5
	n3.left = n6
	n3.right = n7

	// n1 is root
	return n1
}

func main() {
	root := buildTree()
	preOrder(root)
}
