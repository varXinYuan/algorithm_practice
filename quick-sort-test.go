package main

import "fmt"

const SIZE = 100

func quickSort(seq *[SIZE]int, head int, end int) {
	if head >= end {
		return
	}

	// 根据基准值，调整顺序
	sortedIndex := changePos(seq, head, end)

	// 左边排序
	quickSort(seq, head, sortedIndex-1)

	// 右边排序
	quickSort(seq, sortedIndex+1, end)
}

func changePos(seq *[SIZE]int, head int, end int) int {
	baseVal := (*seq)[head]
	i := head + 1
	j := end

	// 根据基准值交换两侧元素
	for i < j {
		for (*seq)[j] > baseVal && j > i {
			j--
		}

		for (*seq)[i] < baseVal && i < j {
			i++
		}

		swap(seq, i, j)
	}

	// 基准值比所有值都小，说明已经就位，无须交换
	if baseVal < (*seq)[i] {
		return head
	}

	// 交换基准值与哨兵的位置
	swap(seq, head, i)

	return i
}

func swap(seq *[SIZE]int, m int, n int) {
	temp := (*seq)[m]
	(*seq)[m] = (*seq)[n]
	(*seq)[n] = temp
}

func main() {
	var seq = [SIZE]int{3, 4, 7, 8, 2, 5, 6, 1, 9, 10000}

	quickSort(&seq, 0, 8)

	for _, n := range seq {
		if n > 0 {
			fmt.Printf("%d\n", n)
		}
	}
}
