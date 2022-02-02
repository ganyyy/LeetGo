package main

import (
	"math/bits"
	"unicode"
)

func ToLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return b
}

type Set struct {
	Cnt int
	S   [128]bool
}

func (s *Set) Add(b byte) {
	if s.S[b] {
		return
	}
	s.S[b] = true
	s.Cnt++
}

func Check(s string) bool {
	var a, b Set
	for i := range s {
		a.Add(s[i])
		b.Add(ToLower(s[i]))
	}

	return a.Cnt == b.Cnt*2
}

//longestNiceSubstringBad 这是一个坏方法, 时间复杂度为O(n^2)
func longestNiceSubstringBad(s string) string {
	var ret string
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			var str = s[i : j+1]
			if Check(str) && len(str) > len(ret) {
				ret = str
			}
		}
	}
	return ret
}

func longestNiceSubstring(s string) (ans string) {
	mask := uint(0)
	// 统计所有字符出现的数量
	for _, ch := range s {
		mask |= 1 << (unicode.ToLower(ch) - 'a')
	}
	maxTypeNum := bits.OnesCount(mask)

	// 以窗口内字符的种类为窗口限制条件,
	for typeNum := 1; typeNum <= maxTypeNum; typeNum++ {
		// 统计窗口内的大小写字符的数量
		var lowerCnt, upperCnt [26]int

		// l,r表示窗口的左右边界
		// total表示当前窗口中字符种类的个数
		// cnt表示成对出现的字符的数量. 如果和typeName的数量相等, 就说明这是一个完美字符串
		var total, cnt, l int
		for r, ch := range s {
			// 扩展右边界
			idx := unicode.ToLower(ch) - 'a'
			if unicode.IsLower(ch) {
				lowerCnt[idx]++
				if lowerCnt[idx] == 1 && upperCnt[idx] > 0 {
					// 首次出现的情况, 增加匹配的字符的对数
					cnt++
				}
			} else {
				upperCnt[idx]++
				if upperCnt[idx] == 1 && lowerCnt[idx] > 0 {
					// 首次出现的情况, 增加匹配的字符的对数
					cnt++
				}
			}
			if lowerCnt[idx]+upperCnt[idx] == 1 {
				total++
			}

			// 收缩左边界
			for total > typeNum {
				idx := unicode.ToLower(rune(s[l])) - 'a'
				// 消减总的计数
				if lowerCnt[idx]+upperCnt[idx] == 1 {
					total--
				}
				if unicode.IsLower(rune(s[l])) {
					lowerCnt[idx]--
					if lowerCnt[idx] == 0 && upperCnt[idx] > 0 {
						cnt--
					}
				} else {
					upperCnt[idx]--
					if upperCnt[idx] == 0 && lowerCnt[idx] > 0 {
						cnt--
					}
				}
				l++
			}

			if cnt == typeNum && r-l+1 > len(ans) {
				ans = s[l : r+1]
			}
		}
	}
	return
}
