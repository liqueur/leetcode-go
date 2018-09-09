package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	pre := &ListNode{0, head}
	root := pre

	for pre.Next != nil {
		if pre.Next.Val == val {
			t := pre.Next
			pre.Next = t.Next
		} else {
			pre = pre.Next
		}
	}

	return root.Next
}

func main() {
	removeElements(nil, 1)
}
