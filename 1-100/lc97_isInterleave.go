package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}
	if l1 == 0 {
		return s2 == s3
	}
	if l2 == 0 {
		return s1 == s3
	}

	// dp[i][j] 表示 s3[:i+j] 可以由 s1[:i]+s2[:j]交替组成
	dp := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]bool, l2+1)
	}
	// 长度都为0时 自然是相等的
	dp[0][0] = true

	// 初始化第一列, 即当 l2 == 0时 s1 能匹配的最大长度
	for i := 1; i <= l1; i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = true
		} else {
			break
		}
	}
	// 初始化第一行, 即当 l1 == 0时 s2 能匹配的最大长度
	for i := 1; i <= l2; i++ {
		if s2[i-1] == s3[i-1] {
			dp[0][i] = true
		} else {
			break
		}
	}

	// 开始遍历
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			c := s3[i+j-1]
			// 分别和两个字符串的当前位置进行对比, 如果有一个符合要求的, 就可以理解为当前位置是符合要求的
			if c == s1[i-1] {
				if dp[i-1][j] {
					dp[i][j] = true
				}
			}
			if c == s2[j-1] {
				if dp[i][j-1] {
					dp[i][j] = true
				}
			}
		}
	}

	return dp[l1][l2]
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	l1, l2 := len(s1), len(s2)
	l3 := len(s3)
	if l1+l2 != l3 {
		return false
	}
	if l1 < l2 {
		l1, l2, s1, s2 = l2, l1, s2, s1
	}
	dp := make([]bool, l2+1)
	dp[0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i > 0 {
				// 上方
				dp[j] = dp[j] && s1[i-1] == s3[i+j-1]
			}
			if !dp[j] && j > 0 {
				// 左方
				dp[j] = dp[j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return dp[l2]
}

func main() {
	isInterleave("aabcc",
		"dbbca",
		"aadbbcbcac")
}
