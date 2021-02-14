package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	var l, r int
	var res int

	for ; r < len(nums); r++ {
		if nums[r] != 1 {
			res = max(res, r-l)
			l = r + 1
		}
	}

	res = max(res, r-l)

	return res
}

func findMaxConsecutiveOnes2(nums []int) int {
	var cnt int
	var res int

	for r := 0; r < len(nums); r++ {
		if nums[r] != 1 {
			res = max(res, cnt)
			cnt = 0
		} else {
			cnt++
		}
	}

	res = max(res, cnt)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1}))
}
