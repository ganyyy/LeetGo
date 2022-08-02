package main

import "sort"

func orderlyQueue(s string, k int) string {
	if k == 1 {
		ans := s
		for i := 1; i < len(s); i++ {
			// 每次都把最前边的字符移动到末尾, 计算出一个最字典序结果
			s = s[1:] + s[:1]
			if s < ans {
				ans = s
			}
		}
		return ans
	}
	// 完全按照字典序排序即可
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	return string(t)
}
