package main

func canIWin(maxChoosableInteger, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	// 选择一定是有答案的, 要么赢, 要么输

	// 位运算压缩
	dp := make([]int8, 1<<maxChoosableInteger)
	// -1表示没有选择
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(int, int) int8
	dfs = func(usedNum, curTot int) (res int8) {
		dv := &dp[usedNum]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		// 迭代所有可以选择的数字
		for i := 0; i < maxChoosableInteger; i++ {
			// 如果当前数字没选,
			// 并且 (选完之后超过了预期, 或者下一轮对手输了), 我就会必赢
			// 找到一种赢得可能即可
			if usedNum>>i&1 == 0 && (curTot+i+1 >= desiredTotal || dfs(usedNum|1<<i, curTot+i+1) == 0) {
				return 1
			}
		}
		return
	}
	return dfs(0, 0) == 1
}
