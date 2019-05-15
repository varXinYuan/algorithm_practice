package main

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

var lastMinK int
var count int = 0
var curNode *Node

func PreOrderWithCount(head *Node) {
	if head == nil {
		return
	}

	PreOrderWithCount(head.Left)

	if count >= lastMinK {
		return
	}
	count++
	curNode = head

	PreOrderWithCount(head.Right)
}

func main() {
	// 测试数据
	n1 := &Node{nil, nil, 2}
	n9 := &Node{nil, nil, 32}

	n2 := &Node{n1, nil, 5}
	n4 := &Node{nil, nil, 12}
	n6 := &Node{nil, nil, 22}
	n8 := &Node{nil, n9, 27}

	n7 := &Node{n6, n8, 25}
	n3 := &Node{n2, n4, 9}

	n5 := &Node{n3, n7, 13}

	lastMinK = 8
	PreOrderWithCount(n5)
	fmt.Println(count, "  ", curNode.Val)
}
