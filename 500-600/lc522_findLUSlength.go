package main

func isSubseq(s, t string) bool {
	ptS := 0
	for ptT := range t {
		if s[ptS] == t[ptT] {
			if ptS++; ptS == len(s) {
				return true
			}
		}
	}
	return false
}

func findLUSlength(strs []string) int {
	ans := -1
	// 没啥好说的, 纯暴力解法
next:
	for i, s := range strs {
		for j, t := range strs {
			// 找到不是其他序列子集的最长序列
			if i != j && isSubseq(s, t) {
				continue next
			}
		}
		if len(s) > ans {
			ans = len(s)
		}
	}
	return ans
}
