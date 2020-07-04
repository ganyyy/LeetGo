package main

import "fmt"

func isMatch(s string, p string) bool {
	if p == "*" {
		return true
	}
	lp := len(p)
	i, j := 0, 0      // 匹配位置
	start, k := -1, 0 // start: *出现的j; k: *出现时的i
	for i < len(s) {
		if j < lp && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
			continue
		}

		if j < lp && p[j] == '*' {
			// 如果出现通配, 记录当前位置
			start = j
			j++
			k = i
			continue
		}

		// 到这一步先看看有没有* , 有的话让*在多吞一个
		if start != -1 {
			k++
			i = k
			j = start + 1
			continue
		}

		// 到这一步只能是匹配失败
		return false
	}

	// 针对多个*的情况
	for j < lp && p[j] == '*' {
		j++
	}
	// 如果p串到头了就说明二者匹配
	return j == lp
}

func isMatch2(s, p string) bool {
	if p == "*" {
		return true
	}
	ls, lp := len(s), len(p)
	dp := make([][]bool, ls+1)
	for i := 0; i <= ls; i++ {
		dp[i] = make([]bool, lp+1)
	}
	// 0,0肯定是true
	dp[0][0] = true
	// 第一列表示为 p == "", 此时是不匹配的
	// 第一行表示为 s == "", 此时只有 p串是由 * 构成的才会匹配
	for i := 1; i <= lp; i++ {
		dp[0][i] = dp[0][i-1] && p[i-1] == '*'
	}

	// 构建dp
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				// 二者共进一步
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// dp[i][j-1]: 表示 * 没用上 如 ab ab*
				// dp[i-1][j]: 表示 * 用上了 如 abcd ab*
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[ls][lp]
}

func main() {
	fmt.Println(isMatch2("aab", "a*b"))
}
