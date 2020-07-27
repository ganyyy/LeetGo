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
