package main

func longestSubstring(s string, k int) int {

	var n = len(s)

	if n < k {
		return 0
	}

	var cnt [26]int
	for i := range s {
		cnt[s[i]-'a']++
	}

	var l, r, res int

	for ; r < n; r++ {
		// 如果字符 s[r] 出现的次数小于 k, 说明这个字符串一定不会包含在结果中
		// 直接进行切割比较即可
		if cnt[s[r]-'a'] < k {
			res = max(res, longestSubstring(s[l:r], k))
			l = r + 1
		}
	}

	// 如果 l == 0, 说明所有的字符都满足大于k, 直接返回即可
	if l == 0 {
		return n
	}
	// 否则就要从 res 和 l:n(r) 中计算一下最大值
	return max(res, longestSubstring(s[l:n], k))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubstringWindow(s string, k int) (ans int) {
	for t := 1; t <= 26; t++ {
		// 按照字符的数量进行统计
		cnt := [26]int{}
		total := 0 // 当前窗口中, 字符类的个数
		lessK := 0 // 当前窗口中, 不满足数量>=k的个数
		l := 0
		for r, ch := range s {
			ch -= 'a'
			if cnt[ch] == 0 {
				total++ // 首次出现的字符, 需要统一增加计数
				lessK++
			}
			cnt[ch]++
			if cnt[ch] == k {
				lessK-- // 满足条件后, 减少目标计数
			}

			// 缩减左边界, 保证集合内符合要求的字符的数量
			for total > t {
				ch := s[l] - 'a'
				if cnt[ch] == k {
					lessK++
				}
				cnt[ch]--
				if cnt[ch] == 0 {
					total--
					lessK--
				}
				l++
			}
			if lessK == 0 {
				// 满足要求的情况下, 更新答案
				ans = max(ans, r-l+1)
			}
		}
	}
	return ans
}
