package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	paths := make([]string, 0)
	traverse(root, &paths, "")
	return paths
}

func traverse(node *TreeNode, paths *[]string, p string) {
	if node == nil {
		return
	}

	if len(p) > 0 {
		p += fmt.Sprintf("->%v", node.Val)
	} else {
		p += fmt.Sprintf("%v", node.Val)
	}

	if node.Left != nil {
		traverse(node.Left, paths, p)
	}

	if node.Right != nil {
		traverse(node.Right, paths, p)
	}

	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, p)
	}
}

/*
  1
 /   \
2     3
 \
  5
*/
func main() {
	node1 := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node4 := &TreeNode{5, nil, nil}

	node1.Left = node2
	node1.Right = node3
	node2.Right = node4

	ret := binaryTreePaths(node1)
	fmt.Println(ret)
}
