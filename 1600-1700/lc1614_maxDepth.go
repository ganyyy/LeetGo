package main

func maxDepth(s string) int {

	var ret int
	var cnt int
	for i := range s {
		switch s[i] {
		case '(':
			cnt++
		case ')':
			ret = max(ret, cnt)
			cnt--
		}
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
