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
	var lc int
	for i := 0; i < lt; i++ {
		if count[t[i]] == 0 {
			lc++
		}
		count[t[i]]++
	}
	var match int
	for right < ls {
		c := s[right]
		if count[c] != 0 {
			window[c]++
			if window[c] == count[c] {
				match++
			}
		}
		right++
		for match == lc {
			if minLen > right-left {
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
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+minLen]
	}
}

func minWindow2(s, t string) string {
	// 这种解法和上边的主要区别在于: 不需要等待整个t串都被包含在进行左指针处理,
	// 而是在某一个字符满足条件就开始处理left
	if len(s) < len(t) {
		return ""
	}
	// 针对字符串类的问题, 往往可以通过定长数组来减少map hash计算的时间
	hash := make([]int, 256)
	// 记录各个值对应的个数
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}

	// l: 左指针
	// count: t串的长度
	// max: 初始最大长度是原串
	// results: 返回的结果
	l, count, max, results := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		// 直接减
		hash[s[r]]--
		// 如果减后还满足>=0说明这是在t串中的数, 满足直接减去即可,
		// 说明[l,r]包含足够数量的t中的s[r]
		if hash[s[r]] >= 0 {
			count--
		}
		// 如果左指针小于右指针 并且是s[l]剩余数量小于0
		// 说明左边可以前进一步, 同时增加s[l]的计数
		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}
		// 如果全部满足了, 并且新窗口比原来小, 就更新一下
		// max和结果的值
		if count == 0 && max > r-l+1 {
			max = r - l + 1
			results = s[l : r+1]
		}
	}

	return results
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
