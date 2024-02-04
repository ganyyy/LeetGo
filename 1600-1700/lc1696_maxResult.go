package main

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n) // dp数组，存储到达每个位置的最大分数
	dp[0] = nums[0]      // 初始化dp的第一个元素

	// 单调递减栈, 队头最大, 可以直接选
	// 入队新元素, 需要从队尾开始清理小于当前和的其他数
	queue := []int{0} // 初始化双端队列，用于存储索引，将第一个元素的索引放入队列

	for i := 1; i < n; i++ {
		// 移除队列前端不在当前考虑的窗口[k]范围内的索引
		for queue[0] < i-k {
			queue = queue[1:]
		}
		// 计算dp[i]，即到达i位置的最大分数
		dp[i] = dp[queue[0]] + nums[i]

		// 移除队列后端所有dp值小于或等于当前dp[i]的索引，保持队列单调递减（这里就是单调栈）
		for len(queue) > 0 && dp[queue[len(queue)-1]] <= dp[i] {
			queue = queue[:len(queue)-1]
		}

		// 将当前索引加入队列后端
		queue = append(queue, i)
	}

	return dp[n-1]
}
