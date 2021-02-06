package main

import "fmt"

func checkPossibility(nums []int) bool {
	var ln = len(nums)
	if ln <= 2 {
		return true
	}

	var r int
	var flag bool

	flag = nums[0] > nums[1]
	// 好吧, 能修改数就好办了
	for r = 1; r < ln-1; r++ {
		if nums[r] > nums[r+1] {
			if flag {
				return false
			}
			// 看一下要更换r还是r+1
			if nums[r-1] > nums[r+1] {
				nums[r+1] = nums[r]
			} else {
				nums[r] = nums[r-1]
			}
			flag = true
		}
	}

	return true
}

func main() {
	var testCases = [][]int{
		{5, 7, 1, 8},
		{5, 7, 8, 1},
		{5, 4, 8, 1},
		{3, 4, 2, 3},
		{4, 2, 3},
		{4, 2, 1},
	}
	for _, v := range testCases {
		fmt.Println(v, checkPossibility(v))
	}
}
