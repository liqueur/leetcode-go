package sol_Same_Tree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p != nil && q == nil) || (p == nil && q != nil) || (p.Val != q.Val) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	node1 := & TreeNode{1, nil, nil}
	node2 := & TreeNode{2, nil, nil}
	node3 := & TreeNode{3, nil, nil}
	node1.Left = node2
	node1.Right = node3
}
