package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}

	var result *ListNode
	if l1.Val > l2.Val {
		result = l2
		result.Next = mergeTwoLists(l1, l2.Next)
	} else {
		result = l1
		result.Next = mergeTwoLists(l1.Next, l2)
	}
	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	L := len(lists)
	if 0 == L {
		return nil
	}
	if 1 == L {
		return lists[0]
	}

	mid := L / 2
	return mergeTwoLists(mergeKLists(lists[0:mid]), mergeKLists(lists[mid:]))
}

func main() {
	l13 := &ListNode{5, nil}
	l12 := &ListNode{4, l13}
	l11 := &ListNode{1, l12}

	l23 := &ListNode{4, nil}
	l22 := &ListNode{3, l23}
	l21 := &ListNode{1, l22}

	l32 := &ListNode{1, nil}
	l31 := &ListNode{1, l32}

	lists := []*ListNode{l11, l21, l31}
	resHead := mergeKLists(lists)

	for resHead != nil {
		fmt.Println(resHead.Val)
		resHead = resHead.Next
	}
}
