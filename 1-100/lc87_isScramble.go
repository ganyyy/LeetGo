package main

func isScramble3(s1 string, s2 string) bool {
	// 区间dp
	// dp[i][j][l]表示 s1[i:i+l]和s2[j:j+l]是相互转换的
	// dp[i][j][k] = k看成长度, w看成在k范围内的交叉点, 即如果是扰乱数,
	// 就必须保证 中间存在一个拐点 w
	// 使得前半段等于前半段:s1[i:i+w]=s2[j:i+w] s1[i+w:i+k]=s2[j+w:j+k]
	// s1 = "ab cde" s2="ba ced"
	// (s11=s21) dp[i][j][w] && (s12=s22)dp[i+w][j+w][k-w]
	// 或者 前半段等于后半段: s1[i:i+w]=s2[j+k-w:j+k] s1[i+w:i+k]=s2[j:j+k-w]
	// s1 = "ab cde" s2="ced ba"
	// (s11=s22) dp[i][j+k-w][w] && (s12=s21)dp[i+w][j][k-w]

	// 日了, 真他妈人才, 这都想得到?
	// 写了都看不懂, 所以写了有啥用?
	// 这玩意是背的?
	ln := len(s1)
	if ln == 1 {
		return s1 == s2
	}
	if ln != len(s2) {
		return false
	}
	dp := make([][][]bool, ln)
	for i := 0; i < ln; i++ {
		dp[i] = make([][]bool, ln)
		for j := 0; j < ln; j++ {
			dp[i][j] = make([]bool, ln+1)
			// 初始化一个字符的情况, 直接对比是否相等即可
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	// k的最小长度从2开始
	for l := 2; l <= ln; l++ {
		// s1的起点位置
		for i := 0; i <= ln-l; i++ {
			// s2的起点位置
			for j := 0; j <= ln-l; j++ {
				// w的位置
				for w := 1; w <= l-1; w++ {
					// 第一种情况
					if dp[i][j][w] && dp[i+w][j+w][l-w] {
						dp[i][j][l] = true
						break
					}
					// 第二种情况
					if dp[i][j+l-w][w] && dp[i+w][j][l-w] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][ln]
}

func isScramble(s, t string) bool {
	if !randomEqual(s, t) {
		// 先看看两者是否相等
		return false
	}
	var check = func(a, b string) bool {
		n := len(s)
		if n == 1 {
			return a == b
		}
		for i := 1; i < n; i++ {
			// 如果相等, 就尝试进行匹配, 这种情况是 a1=b1 a2=b2
			if randomEqual(a[0:i], b[0:i]) {
				if isScramble(a[0:i], b[0:i]) && isScramble(a[i:], b[i:]) {
					return true
				}
			}
			// 这种情况是 a1=b2 a2=b1
			if randomEqual(a[0:i], b[n-i:n]) {
				if isScramble(a[0:i], b[n-i:]) && isScramble(a[i:], b[:n-i]) {
					return true
				}
			}
		}
		return false
	}

	return check(s, t)
}

func randomEqual(a, b string) bool {
	// 检测两个字段是否相等
	if len(a) != len(b) {
		return false
	}
	var count [26]byte

	for i := 0; i < len(a); i++ {
		count[a[i]-'a']++
		count[b[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}

func main() {

}
