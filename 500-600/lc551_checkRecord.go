package main

func checkRecord(s string) bool {
	var acnt int
	var sl, msl int

	var calc = func() bool {
		msl = max(msl, sl)
		sl = 0
		return msl >= 3
	}

	for i := range s {
		if s[i] == 'A' {
			acnt++
			if acnt > 1 {
				return false
			}
			if calc() {
				return false
			}
		} else if s[i] == 'L' {
			sl++
		} else {
			if calc() {
				return false
			}
		}
	}
	calc()
	return acnt < 2 && msl < 3
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
