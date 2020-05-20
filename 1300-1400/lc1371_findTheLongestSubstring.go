package main

func findTheLongestSubstring(s string) int {
	// 先记录所有的元音的位置
	e := make([]byte, 128)
	e['a'] = 1
	e['e'] = 2
	e['i'] = 3
	e['o'] = 4
	e['u'] = 5

	var state, maxSize int
	m := make(map[int]int)
	m[0] = -1
	for i := 0; i < len(s); i++ {
		// 是元音
		if t := e[s[i]]; t != 0 {
			// 这一步 是 将对应的元音位置的比特位置1
			// 如果某一个原因出现了偶数次, 那么state对应位一定是0，否则就是1
			state ^= (1 << (5 - t))
		}
		// 如果出现两个相等的状态, 就说明中间一定是偶数个元音, 获取最大差值即可
		if v, ok := m[state]; ok {
			if i-v > maxSize {
				maxSize = i - v
			}
		} else {
			// 只需要记录一次即可, 以后出现相同的值都以第一次出现为准
			// 这样才能取得最大值
			m[state] = i
		}
	}

	return maxSize
}

func main() {
	findTheLongestSubstring("eleetminicoworoep")
}
