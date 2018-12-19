package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertNodeAfter(currNode *ListNode, newNode *ListNode) *ListNode {
	if newNode == nil {
		return currNode
	}

	newNode.Next = currNode.Next
	currNode.Next = newNode

	return newNode
}

func oddEvenList(head *ListNode) *ListNode {
	oddListHead := head
	evenListHead := head.Next

	// initialize current nodes
	currNode := head.Next.Next
	currOddNode := oddListHead
	currEvenNode := evenListHead

	// iterate linked list
	for currNode != nil {
		oddNode := currNode
		var evenNode *ListNode = nil
		if currNode.Next != nil {
			evenNode = currNode.Next
		}

		var nextNode *ListNode = nil
		if currNode.Next != nil && currNode.Next.Next != nil {
			nextNode = currNode.Next.Next
		}

		currOddNode = insertNodeAfter(currOddNode, oddNode)
		currEvenNode = insertNodeAfter(currEvenNode, evenNode)

		currNode = nextNode
	}

	currEvenNode.Next = nil

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

	res := oddEvenList(n1)

	curr := res
	for curr != nil {
		fmt.Printf("%d\n", curr.Val)

		curr = curr.Next
	}
}
