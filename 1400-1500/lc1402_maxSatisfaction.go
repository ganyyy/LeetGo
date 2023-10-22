package main

import "sort"

func maxSatisfaction(satisfaction []int) int {
	// [left][right] 区间内的最大价值
	// 对于 [mid] 这道菜, 可以选择做或者不做(?)
	// 或则, 先排个序(?)
	// 首先可以确定的是: 尽可能地将 s 值最大的放到最后, s 值较小的选择做或者不做
	sort.Ints(satisfaction)

	var ret, cur, sum int
	for i := len(satisfaction) - 1; i >= 0; i-- {
		cur += satisfaction[i] + sum
		if cur <= 0 || cur < ret {
			break
		}
		ret = cur
		sum += satisfaction[i]
	}

	return ret
}

func maxSatisfactionDP(satisfaction []int) int {
	// 如果是dp的话(?)

	length := len(satisfaction)
	dp := make([]int, length+1)
	sort.Ints(satisfaction)
	res := 0
	// i代表总的菜的数量, i∈[1,length]
	// 有点类似于01背包啊..

	/*
		当将二维dp压缩到一维时, 各个位置对应如下
		leftTop  top
		left     cur

		leftTop: dp[i-1][j-1], pre dp[j]
		top: dp[i-1][j], dp[j]
		left: dp[i][j-1], dp[j-1]
	*/

	// i代表是否选择第i道菜
	for i := 1; i <= length; i++ {
		// j代标选择的菜的数量, j∈[1,i]

		// 取值是左上&上, 所以可以进行压缩

		// 左上表示当前选择这道菜
		var leftTop = 0
		for j := 1; j <= i; j++ {
			// 上表示当前不选择这道菜
			top := dp[j]

			// dp执行的是判断是否要选取第i道菜
			// 第i道菜带来的收益和第几手选择相关

			dp[j] = leftTop + satisfaction[i-1]*j
			if j < i {
				// 可以不选择这道菜等价于从
				// [i-1]道菜中选取j道菜, 那么很明显: j应该小于i!
				// 因为i-1道菜中选取i道菜是不可能的
				dp[j] = max(dp[j], top)
			}
			leftTop = top
			res = max(res, dp[j])
		}
	}
	return res
}

/*
 */
