package main

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	resPos := 0

	for left < right-1 {
		mid = (left + right) / 2

		if nums[mid] < target {
			left = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			resPos = mid
			break
		}
	}

	if left >= right-1 {
		resPos = -1
	}

	return resPos
}

func main() {
	n := [...]int{5}

	fmt.printf("%d", search(n, 5))
}
