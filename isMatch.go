package main

import "fmt"

func isMatch(s string, p string) bool {
	// 解析匹配串
	// 找出 * 的位置
	m := make(map[int]uint8)
	for i, c := range p {
		if c == '*' {
			if i-1 >= 0 {
				m[i-1] = p[i-1]
			}
		}
	}
	pLen := len(p)
	// 没有*的情况
	if len(m) == 0 {
		if pLen != len(s) {
			return false
		}
	}
	pStart := 0
	sSplit := -1 // 记录出现.*匹配时的s位置
	for i := 0; i < len(s); {
		c := s[i]
		if pc, ok := m[pStart]; !ok {
			if pStart >= pLen {
				return false
			}
			cc := p[pStart]
			if cc != '.' && c != cc {
				return false
			}
			pStart++
			i++
		} else {
			// 跳过并尝试匹配下一个字符
			if pc != '.' && pc != c {
				pStart = pStart + 2
			} else {
				if sSplit == -1 {
					sSplit = i
				}
				i++
			}
		}
	}
	if pStart < pLen-2 {
		// 看看后边还有没有, 这是预防 XXX.*XXX的情况, 挨个剔除前边的数
		if sSplit != -1 {
			s = s[sSplit:]
			sLen := len(s)
			p = p[pStart+2:]
			res := false
			for i := 0; i < sLen; i++ {
				res = res || isMatch(s, p)
				if res {
					return true
				}
				s = s[1:]
			}
		}
		return false
	} else {
		return true
	}
}

func isMatch2(s, p string) bool {
	pArr := make([]uint8, 0, len(p))
	pStat := make([]uint8, 0, len(p))
	index := 0
	for i, v := range p {
		if v == '*' {
			pStat[index-1] = 1
		} else {
			pArr = append(pArr, p[i])
			pStat = append(pStat, 0)
			index++
		}
	}
	/*
		pArray
		pa = 0 正常元素
		pa = 1 带*元素
		pa = 2 *
	*/
	start := -1 // 回溯索引
	sStart := -1 // s 回溯
	pIndex := 0 // p串标志位索引, 对应p索引
	i := 0
	for i < len(s) { // 遍历s串
		sv := s[i]         // s串字符
		if pIndex >= len(pArr) {
			if start != -1 {
				pIndex = start
				start = -1
				i = sStart + 1
			} else {
				return false
			}
			continue
		}
		pa := pStat[pIndex] // p标志位

		if sv == pArr[pIndex] || pArr[pIndex] == '.' {
			if pa == 1 {
				start = pIndex
				sStart = i
			} else {
				i++
			}
			pIndex++
		} else {
			if pa == 1 {
				pIndex++
			} else {
				if start != -1 {
					pIndex = start
					start = -1
					i = sStart + 1
				} else {
					return false
				}
			}
		}
 	}

	for i := pIndex; i < len(pArr); i++ {
		if pStat[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isMatch("aabcbcbcaccbcaabc", ".*a*aa*.*b*.c*.*a*"))
}
