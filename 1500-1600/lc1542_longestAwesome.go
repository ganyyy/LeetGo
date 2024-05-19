package main

func longestAwesome(s string) int {
	prefix := map[int]int{0: -1}
	ans := 0
	sequence := 0
	for j := 0; j < len(s); j++ {
		digit := int(s[j] - '0')
		sequence ^= (1 << digit)
		if prevIndex, ok := prefix[sequence]; ok {
			// 按照容斥原理, 如果前边出现了相同的,
			// 那么j-prevIndex这段区间肯定是拥有偶数个不同字符的字符串,、
			// 那么就一定是回文的
			// 注意: 这里不要更新prevIndex的位置,
			// 这样才能保证后续再出来相似的结果时, j-prevIndex是最长的
			ans = max(ans, j-prevIndex)
		} else {
			prefix[sequence] = j
		}
		for k := 0; k < 10; k++ {
			// 这里处理的是字符串内存在一个奇数字符的情况. 总共0-9 10个字符, 全都跑一边.
			if prevIndex, ok := prefix[sequence^(1<<k)]; ok {
				ans = max(ans, j-prevIndex)
			}
		}
	}
	return ans
}
