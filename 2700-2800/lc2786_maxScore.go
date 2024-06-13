package main

func maxScore(nums []int, x int) int64 {
	n := len(nums)
	var dp [2]int
	// +1 切换下标位置
	dp[nums[0]&1] = nums[0]
	dp[(nums[0]+1)&1] = -x
	for i := 1; i < n; i++ {
		pos := nums[i] & 1
		dp[pos] = max(dp[pos], dp[(pos+1)&1]-x) + nums[i]
	}
	return int64(max(dp[0], dp[1]))
}
