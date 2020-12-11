package main

func wiggleMaxLength2(nums []int) int {
	// dp
	if len(nums) <= 1 {
		return len(nums)
	}
	// 这波理解做的不是很好
	// 第一个数即是up也是down
	var up, down = 1, 1
	for i := 1; i < len(nums); i++ {
		// 如果当前是上升趋势, 最大的摇摆长度就是 down+1
		// 反之亦然
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}

	if up > down {
		return up
	} else {
		return down
	}
}

func wiggleMaxLength(nums []int) int {
	// 贪心
	n := len(nums)
	if n < 2 {
		return n
	}
	// 统计峰和谷的数量
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		// 当前是一个峰前边是一个谷
		// 或者
		// 当前是一个谷前边是一个峰
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}
