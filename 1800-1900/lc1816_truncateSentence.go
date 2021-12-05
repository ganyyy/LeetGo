package main

func truncateSentence(s string, k int) string {
	var cnt int
	var idx int
	for cnt < k && idx < len(s) {
		if s[idx] == ' ' {
			cnt++
		}
		idx++
	}
	if idx == len(s) {
		return s
	}
	return s[:idx-1]
}
