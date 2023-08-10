package main

func longestSubstring(s string, k int) int {
	// 分治法
	var n = len(s)
	if n < k {
		return 0
	}

	// 统计各个字符的数量
	var cnt [26]int
	for i := range s {
		cnt[s[i]-'a']++
	}

	var l, res int

	for r := 0; r < n; r++ {
		if cnt[s[r]-'a'] >= k {
			continue
		}
		// 如果该字符不满足条件, 就需要适当的收缩窗口大小
		res = max(res, longestSubstring(s[l:r], k))
		l = r + 1
	}

	if l == 0 {
		// 所有的字符都满足要求, 直接返回即可
		return n
	}
	// 为啥还要再算一遍呢? 因为后半段可能不符合要求, 但是总体是满足要求的的
	// 比如 (aabcabb, 3). a,b 是满足需求的, 但是经过c分割后的 abb中的a,b是不满足需求的
	return max(res, longestSubstring(s[l:], k))
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
