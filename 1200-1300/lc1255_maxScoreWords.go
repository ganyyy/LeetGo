package main

func maxScoreWords(words []string, letters []byte, score []int) int {
	// 位运算的状态枚举
	var ret int
	var cnt [26]int

	n := len(words)

	for _, c := range letters {
		cnt[c-'a']++
	}

next:
	for bit := 1; bit <= (1 << n); bit++ {
		// 统计每个单词所组成的集合
		var wordCount [26]int
		for k, word := range words {
			// 对于当前的bit合集而言, 只统计对应位置为1的单词
			if bit&(1<<k) == 0 {
				continue
			}

			for _, ch := range word {
				idx := ch - 'a'
				wordCount[idx]++
				// 提前终止不必要的判断
				if cnt[idx] < wordCount[idx] {
					continue next
				}
			}
		}

		// 判断子集是否合法
		var cur int
		for idx, c := range wordCount {
			cur += score[idx] * c
		}
		if cur > ret {
			ret = cur
		}
	}
	return ret

}
