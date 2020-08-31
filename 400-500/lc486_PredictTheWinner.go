package main

func PredictTheWinner(nums []int) bool {

	// 如果是奇数, 那么可以按照这个模式走
	// 如果是偶数, 玩家1一定获胜. 因为不管A, B谁能赢, 都能互换
	ln := len(nums)
	if ln&1 == 0 {
		return true
	}
	// dp[i][j] 表示从[i:j+1]能拿到的最优解
	var dp [][]int
	dp = make([][]int, ln)
	for i := 0; i < ln; i++ {
		dp[i] = make([]int, ln)
	}
	// 一共的和, 最后通过计算 sum - dp[0][ln-1]和 dp[0][ln-1]的值判断 A是否可以获胜
	var sum int

	// 首先是只有一个数字的情况, 先手拿肯定是最好的
	for i := 0; i < ln; i++ {
		dp[i][i] = nums[i]
		sum += nums[i]
	}
	// 其次是只有两个数字的情况, 想要赢需要先手拿大的
	for j := 1; j < ln; j++ {
		dp[j-1][j] = max(nums[j-1], nums[j])
	}

	// 再看常规的情况, 从对角线开始递推
	for k := 2; k < ln; k++ {
		for i := 0; i+k < ln; i++ {
			j := i + k
			dp[i][j] = max(
				// 先手取i, 后手会取 nums[j]或者nums[i+1], 此时 玩家能取得的最大值为 nums[i]+去掉后手取得值之后的最优解的最小值
				nums[i]+min(dp[i+1][j-1], dp[i+2][j]),
				// 先手取j, 后手会取 nums[j-1]或者nums[i], 此时 玩家能取得的最大值为 nums[j]+去掉后手取得值之后的最优解的最小值
				nums[j]+min(dp[i][j-2], dp[i+1][j-1]),
			)
		}
	}
	return dp[0][ln-1] >= sum-dp[0][ln-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
