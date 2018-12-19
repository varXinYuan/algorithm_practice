package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m >= n {
		return head
	}

	// get and store pointer to elem (m-1)
	pos := 1
	curr := head
	var before *ListNode = &ListNode{0, head}
	for ; pos < m; pos++ {
		before = curr
		curr = curr.Next
	}

	// iterate linked list, reverse
	var prev *ListNode = nil
	for ; pos <= n; pos++ {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	// get and store pointer to elem (n+1)
	listNewEnd := before.Next
	listNewStart := prev
	after := curr
	if m == 1 {
		head = listNewStart
	}

	// splice before, reverse list, after
	before.Next = listNewStart
	listNewEnd.Next = after

	return head
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

	res := reverseBetween(n1, 3, 4)

	curr := res
	for curr != nil {
		fmt.Printf("%d\n", curr.Val)

		curr = curr.Next
	}
}
