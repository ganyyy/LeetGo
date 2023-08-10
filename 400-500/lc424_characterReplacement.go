package main

func characterReplacement(s string, k int) int {
	// 替换, 替换个毛啊
	// 学什么客户端, 学得会吗?

	// 贪心? DP?

	// 暴力解法: 挨个替换, 然后求最大值.... 程序这么写要爆炸

	// 开始滑动窗口了

	// 当前窗口内, 每个字符的数量
	var cnt [26]int
	// 前后指针
	var left, right int

	// res存储结果, size表示当前窗口内最多的字符的数量
	var res, size int

	for ; right < len(s); right++ {
		var c = s[right]
		cnt[c-'A']++
		// 当前窗口中的 数量最多的字符的数量
		// 这里的size 保存的字符 甚至都有可能不在当前的窗口中
		// size始终都是满足条件的, 数量最多的字符的数量.
		size = max(size, cnt[c-'A'])

		// 当前窗口的字符总数量和 size+k做对比
		// 如果 <=, 那么说明最多替换k个字符, 可以使整个窗口中的字符都是重复的
		// 否则就需要左边界++
		if right-left+1 > size+k {
			// 如果消减的正好属于size的一部分, 该如何处理?
			// 不需要处理, 因为除非存在比当最大数量更多的字符, 否则就没有更新的必要
			cnt[s[left]-'A']--
			left++
		}

		// 更新一下可能出现的最大值
		res = max(res, right-left+1)
	}

	// 理论上, 结果的最大值不会超过 size+k, 但是可能存在k不会全部用掉的情况, 所以需要维护一个最大值
	return res
}

func characterReplacement2(s string, k int) int {
	// 没必要每次进行最大值判断

	// 这种情况没必要处理了, 全部替换好了
	if len(s) <= k {
		return len(s)
	}
	// 贪心? DP?

	// 暴力解法: 挨个替换, 然后求最大值.... 程序这么写要爆炸

	// 开始滑动窗口了

	// 当前窗口内, 每个字符的数量
	var cnt [26]int
	// 前后指针
	var left, right int

	// res存储结果, size表示当前窗口内最多的字符的数量
	var size int

	for ; right < len(s); right++ {
		var c = s[right]
		cnt[c-'A']++
		// 当前窗口中的 数量最多的字符的数量
		// 这里的size 保存的字符 甚至都有可能不在当前的窗口中
		// size始终都是满足条件的, 数量最多的字符的数量.
		size = max(size, cnt[c-'A'])

		// 当前窗口的字符总数量和 size+k做对比
		// 如果 <=, 那么说明最多替换k个字符, 可以使整个窗口中的字符都是重复的
		// 否则就需要左边界++
		if right-left+1 > size+k {
			// 如果消减的正好属于size的一部分, 该如何处理?
			// 不需要处理, 因为除非存在比当最大数量更多的字符, 否则就没有更新的必要
			cnt[s[left]-'A']--
			left++
		}
	}

	// 理论上, 结果的最大值不会超过 size+k, 但是可能存在k不会全部用掉的情况
	if size+k > len(s) {
		return len(s)
	}
	return size + k
}
