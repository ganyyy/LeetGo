package main

func longestDecomposition(text string) int {
	if len(text) == 1 {
		return 1
	}
	k := 0
	str := text
	for len(str) > 0 {
		// 每次都从头开始取子串
		l := 1
		r := len(str) - 1
		for l <= r {
			// 尽可能的从两端获取相同的最短的部分
			if str[:l] == str[r:] {
				k += 2
				str = str[l:r]
				break
			}
			l++
			r--
		}
		if l > r && len(str) > 0 {
			// 最后剩下的单独是一段
			k++
			break
		}
	}
	return k
}
