package main

/*
Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return makeNode(nums, 0, len(nums) - 1)
}

func makeNode(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	node := &TreeNode{nums[mid], nil, nil}
	node.Left = makeNode(nums, left, mid - 1)
	node.Right = makeNode(nums, mid + 1, right)

	return node
}

func main() {

}
