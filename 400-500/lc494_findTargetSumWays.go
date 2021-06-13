package main

func findTargetSumWays(nums []int, target int) int {
	// 还是背包问题啊
	// 每一个值可选可不选
	// 目标总和为 target
	// 这是一个排列问题还是一个组合问题?
	// 组合问题要从备选项入手
	// 所以状态转移方程是?
	// 这里面所有的数值都必须用上, 所以外围应该是基于 target 进行处理
	// DP我是真的不行. 草
	// 所有 nums的最大值是 sum(nums)

	var total int
	for _, v := range nums {
		total += v
	}

	// 假定当前值为 v, 选取该值使其满足target的可能为 dp[target-v]+1; 不选取该值使其满足 target的数量为dp[target+v]+1
	// 假设target = 3, 备选的数字为 [1,1,1,1,1], 实际范围在[-5, 5]之间
	// var dp = make([]int, total*2)
	// dp[0] = 1
	// 值为1有多少种组合?
	// 行不通. 不能简单的计数

	/***
	真正的解法:
		假定存在正子集P, 负子集N
		sum(P) - sum(N) = target
		=> sum(P) - sum(N) + sum(P) + sum(N) = target + sum(P) + sum(N) (两边同时增加sum(P)+sum(N))
		=> 2*sum(P) = target + sum(nums)
		=> sum(P) = (target+sum(nums))/2
	*/

	// 提前返回的条件
	if total < target || (target+total)&1 != 0 {
		return 0
	}

	// 此时相当于从该集合内选取任意数字使其满足 sum(P) == (target+total)>>1
	var s = (target + total) >> 1
	var dp = make([]int, s+1)
	// 此时每个数要么取, 要么不取. 需要总和为 s
	// dp[i] 表示 总和为 i的取法, dp[0] 意味着一个都不取, 只有一种方法
	dp[0] = 1
	// 本质上还是一个0/1背包问题
	// 如果选项是固定的, 那么就需要选项在外边

	// 背包问题啊, 怎么搞呢?

	for _, n := range nums {
		// 外围是选择数字, 因为这是一个组合问题.
		for i := s; i >= n; i-- {
			// i < n的情况不需要处理. 因为不存在负数
			// 内围需要从target开始选取. 这里其实做了一个DP压缩, 由二维压缩到了一维
			// 因为只受前值的影响, 所以要后续遍历
			dp[i] += dp[i-n]
		}
	}
	return dp[s]
}
