package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}

func main() {

}
