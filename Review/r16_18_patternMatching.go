package main

/**
你有两个字符串，即pattern和value。 pattern字符串由字母"a"和"b"组成，
用于描述字符串中的模式。
例如，字符串"catcatgocatgo"匹配模式"aabab"（其中"cat"是"a"，"go"是"b"），
该字符串也匹配像"a"、"ab"和"b"这样的模式。
但需注意"a"和"b"不能同时表示相同的字符串。编写一个方法判断value字符串是否匹配pattern字符串。


解题思路:
设 p 对应的长度为 lp,
设 s 对应的长度为 ls,
设 a 对应的字符串长为 la, 在匹配串中的数量为 ca
设 b 对应的字符串长为 lb, 在匹配串中的数量为 cb, 且满足 ca+cb = lp
那么有
ca * la + cb * lb = ls
其中 ca, cb, ls 是已知的, 需要求出 la 和 lb 的值

因为满足的解必须要是自然数,
所以针对a而言, 可以通过枚举获得所有可能的字符串长度[0, ls/ca],
同时还需要满足 lb 也为一个整数。
*/

// 复制粘贴一气呵成，是刷题中的豪杰
func patternMatching(pattern string, value string) bool {
	// 统计两者的数量
	countA, countB := 0, 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 'a' {
			countA++
		} else {
			countB++
		}
	}
	// 始终保持 A 是大的
	if countA < countB {
		countA, countB = countB, countA
		tmp := ""
		for i := 0; i < len(pattern); i++ {
			if pattern[i] == 'a' {
				tmp += "b"
			} else {
				tmp += "a"
			}
		}
		pattern = tmp
	}
	if len(value) == 0 {
		return countB == 0
	}
	if len(pattern) == 0 {
		return false
	}

	// 遍历可能的 a 的个数
	for lenA := 0; countA*lenA <= len(value); lenA++ {
		// 计算剩余的长度
		rest := len(value) - countA*lenA
		// 都为0 或者 满足整除
		if (countB == 0 && rest == 0) || (countB != 0 && rest%countB == 0) {
			var lenB int
			if countB == 0 {
				lenB = 0
			} else {
				lenB = rest / countB
			}
			pos, correct := 0, true
			var valueA, valueB string
			for i := 0; i < len(pattern); i++ {
				if pattern[i] == 'a' {
					sub := value[pos : pos+lenA]
					if len(valueA) == 0 {
						valueA = sub
					} else if valueA != sub {
						correct = false
						break
					}
					pos += lenA
				} else {
					sub := value[pos : pos+lenB]
					if len(valueB) == 0 {
						valueB = sub
					} else if valueB != sub {
						correct = false
						break
					}
					pos += lenB
				}
			}
			if correct && valueA != valueB {
				return true
			}
		}
	}
	return false
}
