package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertAfter(curNode *ListNode, newNode *ListNode) *ListNode {
	newNode.Next = curNode.Next
	curNode.Next = newNode

	return newNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	listsNum := len(lists)
	loopFlag := listsNum
	for _, list := range lists {
		if list == nil {
			loopFlag--
		}
	}

	sortedLinkedCur := &ListNode{0, nil}
	sortedLinkedHead := sortedLinkedCur

	for loopFlag > 0 {
		// 遍历找出最小值的下标
		minIdx := 0
		for i := range lists {
			if lists[i] == nil {
				continue
			}

			// 初始值
			if lists[minIdx] == nil {
				minIdx = i
				continue
			}

			// compare
			if lists[i].Val < lists[minIdx].Val {
				minIdx = i
			}
		}

		// 将最小值对应的链表指针后移
		// 后移的时候检查当前链表是否已到尾，循环因子减一
		min := lists[minIdx]
		lists[minIdx] = lists[minIdx].Next
		if lists[minIdx] == nil {
			loopFlag--
		}

		// 将找到的最小值插到新的结果链表
		sortedLinkedCur = insertAfter(sortedLinkedCur, min)
	}

	return sortedLinkedHead.Next
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
