package main

func maxSumAfterPartitioning(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}
	// 滑动窗口
	var dp = make([]int, len(arr))
	var curMax = 0

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 小于k的情况
	for i := 0; i < k; i++ {
		if i >= len(arr) {
			break
		}
		curMax = max(curMax, arr[i])
		dp[i] = (i + 1) * curMax
	}
	// 对于后面的元素，它可以自己为一组，或者和前面几个数字为一组
	// ..., arr[i - 3], arr[i - 2], arr[i - 1], (arr[i])    =>  dp(i - 1) + max(arr[i-0:i+1]) * 1
	// ..., arr[i - 3], arr[i - 2], (arr[i - 1], arr[i])    =>  dp(i - 2) + max(arr[i-1:i+1]) * 2
	// ..., arr[i - 3], (arr[i - 2], arr[i - 1], arr[i])    =>  dp(i - 3) + max(arr[i-2:i+1]) * 3
	// ...,
	// ..., (arr[i - k + 1], arr[i - k + 2], ..., arr[i])   =>  dp(i - k) + max(arr[i-k+1:i+1]) * k
	for i := k; i < len(arr); i++ {
		var mm int // 当前位置的最大和
		var gm int // 各种分组的最大和
		for j := 1; j <= k; j++ {
			// 这个分组就是 arr[i-j+1: i+1]
			// 累计和就是 dp[i-j] + max(arr[i-j+1:i+1])*j
			gm = max(gm, arr[i-j+1])
			mm = max(mm, dp[i-j]+gm*j)
		}
		dp[i] = mm
	}
	return dp[len(arr)-1]

}
