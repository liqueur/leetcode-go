package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(getDepth(root.Left) - getDepth(root.Right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(getDepth(node.Left), getDepth(node.Right))
}

func main() {

}
