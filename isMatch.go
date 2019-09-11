package main

import "fmt"

func isMatch(s string, p string) bool {
	// 解析匹配串
	// 找出 * 的位置
	m := make(map[int]uint8)
	for i, c := range p {
		if c == '*' {
			if i - 1 >= 0 {
				m[i - 1] = p[i-1]
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
	for i, c := range s {
		if pc, ok := m[pStart]; !ok {
			if pStart >= pLen {
				return false
			}
			cc := p[pStart]
			if cc != '.' && c != int32(cc) {
				return false
			}
			pStart ++
		} else {
			if pc == '.' {
				if sSplit == -1 {
					sSplit = i
				}
				continue
			}
			// 跳过并尝试匹配下一个字符
			if int32(pc) != c {
				pStart = pStart + 2
			}
		}
	}
	// 看看后边还有没有, 这是预防 XXX.*XXX的情况, 挨个剔除前边的数
	if pStart < pLen-2 {
		s = s[sSplit:]
		sLen := len(s)
		p = p[pStart+2:]
		res := false
		for i := 0; i < sLen; i ++ {
			res = res || isMatch(s, p)
			if res {
				return true
			}
			s = s[1:]
		}
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(isMatch("aaaab", ".*aa."))
}
