package main

func specialPerm(nums []int) (ans int) {
	n := len(nums)
	u := 1<<n - 1
	// [i][x]: 在已选数字集合为i的情况下, 本次尝试选取x拥有的特殊组合数
	memo := make([][]int, u)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	// s 代表剩余可选的数字, 如果对应位置是1则表示可选
	// i表示的是当前选取的数字, 用来判断后续是否选择(%=0)
	dfs = func(s, i int) (res int) {
		if s == 0 {
			return 1 // 找到一个特别排列
		}
		p := &memo[s][i]
		if *p != -1 { // 之前计算过
			return *p
		}
		for j, x := range nums {
			if s>>j&1 > 0 && (nums[i]%x == 0 || x%nums[i] == 0) {
				res += dfs(s^(1<<j), j)
			}
		}
		*p = res // 记忆化
		return
	}
	for i := range nums {
		ans += dfs(u^(1<<i), i)
	}
	return ans % 1_000_000_007
}
