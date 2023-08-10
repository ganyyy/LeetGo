package main

func respace(dictionary []string, sentence string) int {
	// 到第i个字母时, 最小的改动
	ln := len(sentence)
	dp := make([]int, ln+1)
	// 起始是 0
	dp[0] = 0

	for i := 1; i <= ln; i++ {
		// 默认值是都不想等
		dp[i] = dp[i-1] + 1
		// 遍历字典, 如果存在相等的子串就进行比较
		for _, v := range dictionary {
			if lv := len(v); lv <= i {
				if v == sentence[i-lv:i] {
					// 更新最小值
					dp[i] = min(dp[i], dp[i-lv])
				}
			}
		}
	}
	return dp[ln]
}
