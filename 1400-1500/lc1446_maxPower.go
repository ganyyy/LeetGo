package main

func maxPower(s string) int {
	var l, r int
	var ret = 1
	for ; r < len(s); r++ {
		if s[l] != s[r] {
			if r-l > ret {
				ret = r - l
			}
			l = r
		}
	}

	if r-l > ret {
		ret = r - l
	}

	return ret
}
