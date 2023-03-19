package main

func findLexSmallestString(s string, a int, b int) string {
	// n 一定是偶数
	// b 一定 < n
	n := len(s)
	ans := s
	for range s {
		// 轮转
		s = s[n-b:] + s[:n-b]
		cs := []byte(s)
		// 累加
		for j := 0; j < 10; j++ {
			// 奇数位累加
			for k := 1; k < n; k += 2 {
				cs[k] = byte((int(cs[k]-'0')+a)%10 + '0')
			}
			if b&1 == 1 {
				// 偶数位累加
				for p := 0; p < 10; p++ {
					for k := 0; k < n; k += 2 {
						cs[k] = byte((int(cs[k]-'0')+a)%10 + '0')
					}
					s = string(cs)
					if ans > s {
						ans = s
					}
				}
			} else {
				s = string(cs)
				if ans > s {
					ans = s
				}
			}
		}
	}
	return ans
}
