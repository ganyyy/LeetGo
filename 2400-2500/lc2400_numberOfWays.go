package main

func numberOfWays(startPos, endPos, k int) int {

	const MOD = 1e9 + 7

	type pair struct{ x, y int }
	dp := map[pair]int{}
	var f func(int, int) int
	// 当前的位置, 剩余的步数
	f = func(x, left int) int {
		// 不管向左还是向右, 如果差距超过了剩余的步数, 直接枝减即可
		if abs(x-endPos) > left {
			return 0
		}
		if left == 0 {
			return 1 // 成功到达终点，算一种方案
		}
		// 到达某点的几种可能
		p := pair{x, left}
		if v, ok := dp[p]; ok {
			return v
		}
		// 向左 or 向右. 剩余的步数都需要-1
		res := (f(x-1, left-1) + f(x+1, left-1)) % MOD
		// 记忆一下当前组合对应的步数
		dp[p] = res
		return res
	}
	return f(startPos, k)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
