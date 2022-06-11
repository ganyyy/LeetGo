package main

func minFlipsMonoIncr(s string) int {
	var length = len(s)
	// dp[i]可以理解为, s[:i]中存在的1的个数
	var dp = make([]int, length+1)
	for i := 1; i <= length; i++ {
		dp[i] = dp[i-1] + int(s[i-1]-'0')
	}
	var ret = length
	for i := 1; i <= length; i++ {
		var zero = dp[i-1]                            // 从头到尾替换成0的消耗
		var one = (length - i) - (dp[length] - dp[i]) // 从尾到头替换成1的消耗
		if t := zero + one; t < ret {
			ret = t
		}
	}
	return ret
}
