package main

func maxScore(s string) int {
	// dp?

	var c1 int
	var max int

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			c1++
		}
	}
	var c0 int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			c0++
		} else {
			c1--
		}
		if c0+c1 > max {
			max = c0 + c1
		}
	}
	return max
}
