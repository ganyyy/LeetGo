package main

func partitionDisjoint(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lm, m := nums[0], nums[0]
	var index int
	for i := 1; i < len(nums); i++ {
		m = max(m, nums[i])
		// 很巧妙的做法
		if lm > nums[i] {
			lm = m
			index = i
		}
	}
	return index + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
