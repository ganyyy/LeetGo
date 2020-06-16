package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	index := 0
	char := uint8(0)
	flag := false
	for {
		s := 0
		e := len(strs) - 1
		for s < e {
			ss := strs[s]
			es := strs[e]
			s++
			e--
			if index >= len(ss) || index >= len(es) || ss[index] != es[index] {
				flag = true
				break
			} else {
				char = ss[index]
			}

			if s == e {
				mid := strs[s]
				if index >= len(mid) || mid[index] != char {
					flag = true
					break
				}
			}
		}
		if flag {
			break
		} else {
			index++
		}
	}
	return strs[0][:index]
}

func longestCommonPrefixN(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	res := strs[0]
	for _, v := range strs[1:] {
		var i int
		for ; i < len(v) && i < len(res) && v[i] == res[i]; i++ {
		}
		if i < len(res) {
			res = res[:i]
		}
	}

	return res
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"aaa", "aa", "aaa"}))
}
