package sol_Remove_Duplicates_from_Sorted_List

import (
	"fmt"
)

/*
Input: 1->1->2
Output: 1->2

Input: 1->1->2->3->3
Output: 1->2->3
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	pre := head
	cur := head.Next
	for cur != nil {
		if pre.Val != cur.Val {
			pre.Next = cur
			pre = cur
		}
		cur = cur.Next
	}
	pre.Next = nil
	return head
}

func main() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{1, nil}
	node3 := &ListNode{2, nil}
	node4 := &ListNode{3, nil}
	node5 := &ListNode{3, nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	deleteDuplicates(node1)
	for node1 != nil {
		fmt.Println(node1.Val)
		node1 = node1.Next
	}
}
