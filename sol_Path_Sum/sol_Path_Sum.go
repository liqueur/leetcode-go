package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	check := false
	pathSum(root, sum, 0, &check)
	return check
}

func pathSum(root *TreeNode, sum int, presum int, check *bool) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if root.Val + presum == sum {
			*check = true
		}
	}

	pathSum(root.Left, sum, root.Val + presum, check)
	pathSum(root.Right, sum, root.Val + presum, check)
}

/*
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
 */
func main() {
	node1 := &TreeNode{5, nil, nil}
	node2 := &TreeNode{4, nil, nil}
	node3 := &TreeNode{8, nil, nil}
	node4 := &TreeNode{11, nil, nil}
	node5 := &TreeNode{13, nil, nil}
	node6 := &TreeNode{4, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node8 := &TreeNode{2, nil, nil}
	node9 := &TreeNode{1, nil, nil}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4

	node3.Left = node5
	node3.Right = node6

	node4.Left = node7
	node4.Right = node8

	node6.Right = node9

	check := hasPathSum(node1, 22)
	fmt.Println(check)
}
