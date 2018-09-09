package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	if root != nil {
		queue = append(queue, root)

		for len(queue) > 0 {
			level := make([]int, 0)
			size := len(queue)

			for i := 0; i < size; i++ {
				node := queue[0]
				queue = queue[1:]
				level = append(level, node.Val)
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
			res = append(res, level)
		}
	}

	for i := 0; i < len(res) / 2; i++ {
		res[i], res[len(res) - i - 1] = res[len(res) - i - 1], res[i]
	}

	return res
}

/*
   3
   / \
  9  20
    /  \
   15   7
 */
func main() {
	node1 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{9, nil, nil}
	node3 := &TreeNode{20, nil, nil}
	node4 := &TreeNode{15, nil, nil}
	node5 := &TreeNode{7, nil, nil}

	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5

	levelOrder(node1)
}
