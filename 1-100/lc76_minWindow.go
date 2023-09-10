package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	if ls == 0 {
		return ""
	}
	if lt == 0 {
		return ""
	}
	if ls < lt {
		return ""
	}
	minLen, start := math.MaxInt32, 0
	left, right := 0, 0
	count := make([]byte, 128)
	window := make([]byte, 128)
	var lc int // lc标识t中出现的不同字符的个数
	for i := 0; i < lt; i++ {
		if count[t[i]] == 0 {
			lc++
		}
		count[t[i]]++
	}
	var match int
	for right < ls {
		c := s[right]
		if count[c] != 0 { // 如果是一个存在于t中的字符串
			window[c]++
			if window[c] == count[c] { // 并且当前c的数量已经满足要求了
				match++
			}
		}
		right++
		for match == lc { // 全部字符的数量都满足了, 不断消减到出现第一个不满足的字符
			if minLen > right-left { // 计算最小的长度
				minLen = right - left
				start = left
			}
			c := s[left]
			if count[c] != 0 {
				window[c]--
				if window[c] < count[c] {
					match--
				}
			}
			left++ // 双指针缩减空间
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+minLen]
	}
}

// 明天过一遍窗口问题

func minWindow2(s, t string) string {
	// 这种解法和上边的主要区别在于: 不需要等待整个t串都被包含在进行左指针处理,
	// 而是在某一个字符满足条件就开始处理left
	if len(s) < len(t) {
		return ""
	}
	// 针对字符串类的问题, 往往可以通过定长数组来减少map hash计算的时间
	hash := make([]int, 128)
	// 记录各个值对应的个数
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}

	// l: 左指针
	// count: t串剩余待匹配的字符数量. 这个解法神奇就神奇在只有减少, 没有增加
	//		  从count第一次为0开始, 往后的count都是0!
	//		  因为此时不会再次发生hash[s[r]] >= 0的情况
	// minLen: 初始最大长度是原串
	// results: 返回的结果
	l, count, minLen, results := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		// 直接减
		hash[s[r]]--
		// 如果不是t串中的字符, 此时的hash[s[r]]会小于0. 此时不会减少count
		if hash[s[r]] >= 0 {
			// 说明[l,r]包含足够数量的t中的s[r]
			count--
		}
		// 如果左指针小于右指针 并且是s[l]剩余数量小于0
		// 说明左边可以前进一步, 同时增加s[l]的计数
		// 如果不是t串中的字符, 那么hash[s[l]]会一直小于等于0
		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}
		// 如果全部满足了, 并且新窗口比原来小, 就更新一下
		// max和结果的值
		if count == 0 && minLen > r-l+1 {
			minLen = r - l + 1
			results = s[l : r+1]
		}
	}

	return results
}

func minWindow3(s string, t string) string {
	var cnt, cur [64]int
	var num, lt int
	lt = len(t)
	if lt == 0 || lt > len(s) {
		return ""
	}
	var getIndex = func(b byte) int { return int(b) & 63 }
	for i := range t {
		cnt[getIndex(t[i])]++
	}

	var ret string

	var left int // 起始位置
	for i := range s {
		idx := getIndex(s[i])
		if cnt[idx] == 0 {
			if num == 0 {
				// 重置起点位置
				left = idx + 1
			}
			continue
		}
		if num == 0 {
			// 记录起点位置
			left = i
		}
		// 增加当前计数
		cur[idx]++
		if cur[idx] <= cnt[idx] {
			// 增加有效计数
			num++
		}

		// 移动左指针
		for num == lt {
			idx = getIndex(s[left])
			valid := s[left : i+1]
			if ret == "" || len(valid) < len(ret) {
				ret = valid
			}
			left++
			if cnt[idx] == 0 {
				continue
			}
			cur[idx]--
			if cur[idx] < cnt[idx] {
				num--
			}
		}
		for left < len(s) && cnt[getIndex(s[left])] == 0 {
			left++
		}
	}

	return ret
}

func findAnagrams(s string, p string) []int {
	ls, lp := len(s), len(p)
	var res []int
	if ls < lp {
		return res
	}
	hash := make([]int, 128)
	for i := 0; i < lp; i++ {
		hash[p[i]]++
	}
	count := lp
	for l, r := 0, 0; r < ls; {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}
		r++
		if r-l == lp {
			// 判断是不是异构
			if count == 0 {
				res = append(res, l)
				hash[s[l]]++
				count++
			} else {

			}
			l++
		}
	}
	return res
}

func main() {

	fmt.Println(findAnagrams("cbaebacbad",
		"abc"))
}
