package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NULL_NODE_VAL = -1

type MinHeal struct {
	cap   int
	count int
	main  []int
	emIdx int //当前空节点
}

func father(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return idx*2 + 1
}

func rightChild(idx int) int {
	return idx*2 + 2
}

func (h *MinHeal) HasLeft(idx int) bool {
	if idx >= h.cap || leftChild(idx) >= h.cap {
		return false
	}

	if h.main[leftChild(idx)] == NULL_NODE_VAL {
		return false
	}

	return true
}

func (h *MinHeal) HasRight(idx int) bool {
	if idx >= h.cap || rightChild(idx) >= h.cap {
		return false
	}

	if h.main[rightChild(idx)] == NULL_NODE_VAL {
		return false
	}

	return true
}

func (h *MinHeal) Insert(val int) {
	// 插入元素
	// 若堆已满，pop出元素
	if h.count >= h.cap {
		h.Pop()
	}

	// 插入元素
	h.main[h.emIdx] = val

	// 重平衡节点
	curIdx := h.emIdx
	for curIdx != 0 {
		if h.main[curIdx] < h.main[father(curIdx)] {
			h.Swap(curIdx, father(curIdx))
			curIdx = father(curIdx)
		} else {
			break
		}
	}

	if h.count < h.cap {
		h.count++
		h.emIdx = h.count
	}
}

func (h *MinHeal) Swap(idx1 int, idx2 int) {
	temp := h.main[idx1]
	h.main[idx1] = h.main[idx2]
	h.main[idx2] = temp
}

func (h *MinHeal) Pop() int {
	if h.count <= 0 {
		return NULL_NODE_VAL
	}

	curIdx := 0
	popVal := h.main[curIdx]

	for {
		if h.HasLeft(curIdx) && h.HasRight(curIdx) {
			if h.main[rightChild(curIdx)] < h.main[leftChild(curIdx)] {
				h.main[curIdx] = h.main[rightChild(curIdx)]
				curIdx = rightChild(curIdx)
			} else {
				h.main[curIdx] = h.main[leftChild(curIdx)]
				curIdx = leftChild(curIdx)
			}
		} else if h.HasLeft(curIdx) {
			h.main[curIdx] = h.main[leftChild(curIdx)]
			curIdx = leftChild(curIdx)
		} else if h.HasRight(curIdx) {
			h.main[curIdx] = h.main[rightChild(curIdx)]
			curIdx = rightChild(curIdx)
		} else {
			h.emIdx = curIdx
			h.main[curIdx] = NULL_NODE_VAL
			h.count--
			break
		}
	}

	return popVal
}

func NewHeal(cap int) *MinHeal {
	main := make([]int, cap, cap*2)
	for i := 0; i < cap; i++ {
		main[i] = NULL_NODE_VAL
	}

	return &MinHeal{
		cap,
		0,
		main,
		0,
	}
}

func main() {
	// 随机数测试
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	k := 10
	heal := NewHeal(k)

	for i := 0; i < 100; i++ {
		heal.Insert(r1.Intn(100))
	}

	//for _, val := range heal.main {
	//fmt.Println(val)
	//}
	//fmt.Println("------------")

	for heal.count > 0 {
		fmt.Println(heal.Pop())
	}
}
