package main

import (
	"math"
)

// 失败了. 想想也是, 不超时就邪门了
func minCutOverTime(s string) int {
	// 貌似可以直接套用上一题的模板?

	type State uint8
	const (
		NotSearch State = iota // 未搜索
		True                   // 是回文串
		False                  // 不是回文串
	)

	var n = len(s)

	// 初始化 dp数组
	var dp = make([][]State, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]State, n)
	}

	// 递归确认函数
	var isPalindrome func(i, j int) State
	isPalindrome = func(i, j int) State {
		// 大于等于的不用管了
		if i >= j {
			return True
		}

		// 搜索过的也不用管了
		if dp[i][j] != NotSearch {
			return dp[i][j]
		}

		// 默认为不合法的
		dp[i][j] = False
		if s[i] == s[j] {
			// 当前是不是回文串取决于前边是不是回文串
			dp[i][j] = isPalindrome(i+1, j-1)
		}
		return dp[i][j]
	}

	var res = math.MaxInt32
	var cnt int
	// 递归遍历函数
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			if cnt < res {
				res = cnt
			}
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(i, j) == True {
				cnt++
				dfs(j + 1)
				cnt--
			}
		}
	}
	dfs(0)

	return res - 1
}

func minCut(s string) int {
	// 应该会用到贪心吧, 就是从每个位置出发, 寻找最大的递归字串

	// 直接构建DP

	// 转变一下思路

	// 寻找以每个字符为中心的最长回文字串, 并记录需要分割的最小次数

	var n = len(s)
	if n <= 1 {
		return 0
	}

	// dp[i] 表示s[:i+1]分割后每个字符串都是回文串所需要的最小分割次数
	var dp = make([]int, n)
	for i := 0; i < n; i++ {
		// 默认情况下, 将每一个字符单独分割.
		// 所以初始情况下
		// 每个节点都需要分割 n-1 次
		dp[i] = n - 1
	}

	// i, j 分别为查找的起点
	// i == j || i == j-1
	// 对应单数回文串和双数回文串的情况
	var helper = func(i, j int) {
		// 如果满足回文串的条件
		for i >= 0 && j < n && s[i] == s[j] {
			// 可以获得的最长回文串为 [i, j]
			if i == 0 {
				// 如果到头了, 说明不需要分割. 这就是一个很大的回文串
				dp[j] = 0
				break
			} else {
				// 否则, 将 [:j+1]分割为两段,
				// [:i], [i:j+1]
				// 此时dp[j] 的值相当于在 [:i]的基础上加一次分割
				// 所以取 min(dp[j], dp[i-1]+1)
				dp[j] = min(dp[j], dp[i-1]+1)
			}
			i--
			j++
		}
	}

	for i := 0; i < n; i++ {
		// 单数
		helper(i, i)
		// 双数
		helper(i, i+1)
	}

	return dp[n-1]

}
