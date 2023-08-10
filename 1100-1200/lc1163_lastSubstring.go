//go:build ignore

package main

func lastSubstring(s string) string {
	i, j, n := 0, 1, len(s)
	// s[i:i+k+1] = ret
	// s[j:j+k+1] = 和ret进行比较的字符串
	for j < n {
		k := 0
		for j+k < n && s[i+k] == s[j+k] {
			// case1: s[i:i+k+1] == s[j:j+k+1], 二者字典序相同
			k++
		}
		if j+k < n && s[i+k] < s[j+k] {
			// case2: s[i+k] < s[j+k], 则 s[i:i+k+1] < s[j:j+k+1]
			//      此时, 需要将ret更新为 s[j:j+k+1], 即i=j
			//      j的取值, 取j+1和i+k+1的最大值, 可以跳过更多的可能性
			i, j = j, max(j+1, i+k+1)
		} else {
			// case3: s[i+k] > s[j+k], 则 s[i:i+k+1] > s[j:j+k+1]
			//      此时, ret 保留, j 直接跳过前半部分
			j = j + k + 1
		}
	}
	return s[i:]
}
