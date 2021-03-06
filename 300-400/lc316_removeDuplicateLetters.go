package main

func removeDuplicateLetters(s string) string {
	// 保留最终的结果
	var stack = make([]byte, 0, len(s))
	// 统计stack中已存在的字符
	var set = [26]bool{}
	// 统计每个字符的个数
	var cnt = [26]int{}

	for i := range s {
		cnt[s[i]-'a']++
	}

	// 依次移除,
	for i := range s {
		// 如果是一个不存在的字符
		var c = s[i] - 'a'
		if !set[c] {
			// 栈不为空, 且栈顶元素当前的存量大于0, 且栈顶元素大于当前元素
			for len(stack) != 0 && cnt[stack[len(stack)-1]-'a'] > 0 && stack[len(stack)-1] > s[i] {
				// 删除栈顶并出栈
				set[stack[len(stack)-1]-'a'] = false
				stack = stack[:len(stack)-1]
			}
			// 加进去,
			set[c] = true
			stack = append(stack, c+'a')
		}
		// 减少计数
		cnt[c]--
	}

	return string(stack)
}

func removeDuplicateLettersNew(s string) string {
	// 最终结果
	var stack = make([]byte, 0, len(s))
	// 临时计数
	var cnt [26]int
	// 存在标记
	var set [26]bool

	// 统计每个元素的个数
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	var top int
	for i := 0; i < len(s); i++ {
		var idx = s[i] - 'a'
		if set[idx] {
			cnt[idx]--
			continue
		}
		for top = len(stack) - 1; top >= 0 && cnt[stack[top]-'a'] > 0 && stack[top] > s[i]; top = len(stack) - 1 {
			// 出栈
			set[stack[top]-'a'] = false
			stack = stack[:top]
		}
		set[idx] = true
		stack = append(stack, s[i])
		cnt[idx]--
	}
	return string(stack)
}
