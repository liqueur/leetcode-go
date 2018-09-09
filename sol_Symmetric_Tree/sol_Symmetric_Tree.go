package sol_Symmetric_Tree

import (
	"fmt"
	)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return comp(root.Left, root.Right)
}

func comp(left *TreeNode, right *TreeNode) bool {
	if left == nil {
		return right == nil
	}

	if right == nil {
		return left == nil
	}

	if left.Val != right.Val {
		return false
	}

	return comp(left.Left, right.Right) && comp(left.Right, right.Left)
}

func preorder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
}

func inorder(root *TreeNode) {
	if root != nil {
		inorder(root.Left)
		fmt.Print(root.Val)
		inorder(root.Right)
	}
}

func postorder(root *TreeNode) {
	if root != nil {
		postorder(root.Left)
		postorder(root.Right)
		fmt.Print(root.Val)
	}
}

/*
    1
   / \
  2   2
 / \ / \
3  4 4  3
 */
func main() {
	node1 := & TreeNode{1, nil, nil}
	node2 := & TreeNode{2, nil, nil}
	node3 := & TreeNode{2, nil, nil}
	node4 := & TreeNode{3, nil, nil}
	node5 := & TreeNode{4, nil, nil}
	node6 := & TreeNode{4, nil, nil}
	node7 := & TreeNode{3, nil, nil}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	fmt.Println("preorder")
	preorder(node1)
	fmt.Println("\ninorder")
	inorder(node1)
	fmt.Println("\npostorder")
	postorder(node1)
}
