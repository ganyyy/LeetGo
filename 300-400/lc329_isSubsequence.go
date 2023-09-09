package main

func isSubsequence(s string, t string) bool {
	if s == t {
		return true
	}
	var sl, tl int
	lenS, lenT := len(s), len(t)
	for tl < lenT {
		if sl < lenS && s[sl] == t[tl] {
			sl++
		}
		if sl == lenS {
			return true
		}
		tl++
	}
	return false
}

func isSubsequence2(s string, t string) bool {
	sl := len(s)
	if sl > len(t) {
		return false
	}
	if sl == 0 || s == t {
		return true
	}

	var si int
	for ti := range t {
		if t[ti] == s[si] {
			si++
			if si >= sl {
				return true
			}
		}
	}
	return false
}
