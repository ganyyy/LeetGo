package main

func countSegments(s string) int {
	var cnt int
	var check bool

	for i := range s {
		if s[i] == ' ' {
			check = false
			continue
		}
		if check {
			continue
		}
		cnt++
		check = true
	}
	return cnt
}
