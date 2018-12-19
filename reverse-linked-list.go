package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode = nil
	curr := head

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}
	n5 := &ListNode{5, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	res := reverseList(n1)

	curr := res
	for curr != nil {
		fmt.Printf("%d\n", curr.Val)

		curr = curr.Next
	}
}
