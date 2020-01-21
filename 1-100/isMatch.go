package main

import "fmt"

func isMatch3(s string, p string) bool {
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
	start := -1  // 回溯索引
	sStart := -1 // s 回溯
	pIndex := 0  // p串标志位索引, 对应p索引
	i := 0
	for i < len(s) { // 遍历s串
		sv := s[i] // s串字符
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

const (
	TRUE  = 1
	FALSE = 2
)

// dp[i][j] == true 表示 s [:i+1] 和 p[:j+1] 是匹配的
var res [][]int

func isMatch(s, p string) bool {
	// 初始化, len(s)行 len(p)列
	res = make([][]int, len(s)+1, len(s)+1)
	for i := range res {
		res[i] = make([]int, len(p)+1, len(p)+1)
	}

	return dp(0, 0, s, p)
}

func dp(i, j int, s, p string) bool {
	// 有结果返回结果
	if res[i][j] != 0 {
		return res[i][j] == TRUE
	}
	// 匹配结果
	var ans bool
	if j == len(p) {
		// 如果相等就更新标志位
		ans = i == len(s)
	} else {
		// 匹配当前字符是否相等
		match := i < len(s) && (s[i] == p[j] || p[j] == '.')
		// 看看匹配串下一个是不是*
		if j+1 < len(p) && p[j+1] == '*' {
			// 存在且是 * , 则去匹配 i 和p的下一个常规字符(跳过*相当于+2)  或者满足i==j后 匹配 i+1和 j
			// 疑问：如何保证j+2 < len(p) ?
			ans = (dp(i, j+2, s, p)) || match && dp(i+1, j, s, p)
		} else {
			// 如果j+1不是*, 则匹配i+1,j+1
			ans = match && dp(i+1, j+1, s, p)
		}
	}
	if ans {
		res[i][j] = TRUE
	} else {
		res[i][j] = FALSE
	}
	return ans
}

func main() {
	fmt.Println(isMatch("aaab", "a*ab"))
}
