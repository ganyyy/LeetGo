package main

func maxConsecutiveAnswers(answerKey string, k int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var maxLen = func(want byte) int {
		var ret int
		var start, end int
		var cur = k
		for end = 0; end < len(answerKey); end++ {
			if answerKey[end] == want {
				continue
			}
			// 先让可以替换的字符串最长
			if cur > 0 {
				cur--
				continue
			}
			ret = max(ret, end-start)
			// 再次计算到第一个不满足的字符
			for answerKey[start] == want {
				start++
			}
			// 跳过这个字符, 再看再次计算最长的结果
			start++
		}
		return max(ret, end-start)
	}

	// fmt.Println(a, b)
	return max(maxLen('T'), maxLen('F'))
}
