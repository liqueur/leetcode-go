package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	head = reverseList(p.Next)
	p.Next.Next = p
	p.Next = nil

	return head
}

func main() {

}
