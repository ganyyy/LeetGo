package main

import "sort"

func combinationSum4(nums []int, target int) int {
	// dfs 搞一下. 十有八九会溢出吧

	// 不进行sort也没问题. 但是sort之后可以方便的进行提前退出循环.
	sort.Ints(nums)

	// 如果出现了负数, 那么依旧可以通过统一加入最小值的绝对值将其转换为整数数组进行处理
	// 需要注意的点为: 相加不能超过 int的最大表示范围

	// dp 表示 到达每一个数字的可能次数
	var dp = make([]int, target+1)
	dp[0] = 1
	// 关于目标和选项谁在外的问题, 是个大学问
	// 如果是组合([1,2]等同于[2,1])问题, 那么选项在外边
	// 如果是排列([1,2]不同于[2,1])问题, 那么目标在外边
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i == nums[j] {
				dp[i] += 1
				break
			} else if i > nums[j] {
				dp[i] += dp[i-nums[j]]
			} else {
				break
			}
		}
	}
	return dp[target]
}
