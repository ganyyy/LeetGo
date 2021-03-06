package main

func partition(s string) [][]string {
	// 怎么分割回文串?
	// 这需要写个递归?
	var res [][]string

	// dp预计算所有回文串
	var n = len(s)
	// dp[i][j] 表示 s[i:j+1]为回文串
	// i >= j, dp[i][j] = true
	// i < j, dp[i][j] = s[i]==s[j] && dp[i+1][j-1]
	var dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	var dfs func(i int)
	var tmp []string
	dfs = func(i int) {
		if i == n {
			var t = make([]string, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				tmp = append(tmp, s[i:j+1])
				dfs(j + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0)
	return res
}

func partition2(s string) [][]string {
	// 上记忆化搜索
	// 所谓的记忆化搜索, 就是带状态的DP

	type State uint8

	const (
		NotSearch State = iota // 未搜索
		True                   // 是回文串
		False                  // 不是回文串
	)

	var n = len(s)

	// 初始化dp数组
	var dp = make([][]State, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]State, n)
	}

	// i, j 表示s[i:j+1]的搜索状态
	var isPalindrome func(i, j int) State

	isPalindrome = func(i, j int) State {
		// 大于等于直接返回即可
		if i >= j {
			return True
		}
		// 如果已经搜索过, 直接返回结果
		if dp[i][j] != NotSearch {
			return dp[i][j]
		}
		// 如果没有搜索过, 就判断是否是回文串
		dp[i][j] = False
		if s[i] == s[j] {
			dp[i][j] = isPalindrome(i+1, j-1)
		}
		return dp[i][j]
	}

	var tmp []string
	var res [][]string
	var dfs func(i int)

	dfs = func(i int) {
		if i == n {
			var t = make([]string, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(i, j) == True {
				tmp = append(tmp, s[i:j+1])
				dfs(j + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0)
	return res
}
