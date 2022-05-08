package main

func diStringMatch(s string) []int {
	var ret = make([]int, len(s)+1)

	var bigger, little = len(s), 0

	for i := range s {
		if s[i] == 'I' {
			ret[i] = little
			little++
		} else {
			ret[i] = bigger
			bigger--
		}
	}
	ret[len(s)] = bigger

	return ret
}
