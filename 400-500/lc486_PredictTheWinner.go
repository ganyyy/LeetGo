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

func PredictTheWinnerWWW(nums []int) bool {
	if len(nums)&1 == 0 {
		// 偶数数组的情况下, 可以控制拿奇数还是偶数. 两部分的和一定存在 >= 的关系
		// 所以可以直接返回true
		return true
	}

	// dp多种方法, 关键是找准变量
	// 可以是数量, DP顺序, DP数量, DP的屁啊

	// dp[i][j] 表示在 nums[i:j+1]中, 玩家1的最优解(即1比2多获得的值的大小)
	// 有三种情况:
	//  i == j, 此时只有一种取法 所以 dp[i][j] = nums[i]
	//  j - i == 1, 此时有两种取法, 取最大值即可 dp[i][j] = abs(nums[i] - nums[j])
	//  j - i > 1, 此时需要根据前态进行判断.
	//       如果获取i, 那么对手能获得的最大收益为dp[i+1][j]
	//       如果获取j, 那么对手能获得的最大收益为dp[i][j-1]
	//       所以dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
	// 最终返回的结果就是 dp[0][n-1] >= 0

	var dp = make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}

	// 填充i==j的情况
	for i := range nums {
		dp[i][i] = nums[i]
	}
	// 填充j-i==1的情况
	// for i := 0; i < len(nums)-1; i++ {
	//     dp[i][i+1] = abs(nums[i]-nums[i+1])
	// }

	// 定义边界组的形式, 指定每一次DP的步长
	for k := 1; k < len(nums); k++ {
		// 计算区间的最优解
		for i, j := 0, k; j < len(nums); i, j = i+1, j+1 {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	return dp[0][len(nums)-1] >= 0
}
