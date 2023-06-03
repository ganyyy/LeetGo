//go:build ignore

package main

func maxRepOpt1(text string) int {
	// 计数
	var cnt [26]int
	var add = func(b byte) {
		cnt[b-'a']++
	}
	var get = func(b byte) int {
		return cnt[b-'a']
	}

	for _, c := range text {
		add(byte(c))
	}

	var ret int
	var ln = len(text)
	for i := 0; i < ln; {
		// 获取j, 使得[i:j]都是同一个字符
		ib := text[i]
		j := i + 1
		for ; j < ln && text[j] == text[i]; j++ {
		}
		// 判断长度是否还有剩余(意味着可以获取一个额外的字符进行填充)
		total := get(ib)
		oj := j
		// 前后都可以填充欸(!)
		if (j < ln || i > 0) && j-i < total {
			j++ // 如果填充了后边, 相当于跳过一个.
			// 二次获取 [i:j_old] + [j_old] + [j_old:j_new]
			for ; j < ln && text[j] == text[i] && j-i < total; j++ {
			}
		}
		ret = max(ret, j-i)
		i = oj
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
