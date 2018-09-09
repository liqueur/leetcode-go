package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head
	half := make([]int, 0)

	for fast != nil && fast.Next != nil {
		half = append(half, slow.Val)
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	fmt.Println(half)

	for i := 0; i < len(half); i++ {
		if slow.Val != half[len(half) - 1 - i] {
			return false
		} else {
			slow = slow.Next
		}
	}

	return true
}

func main() {

}
