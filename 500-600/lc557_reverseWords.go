package main

func reverseWords(s string) string {
	var left, right int
	tmp := []byte(s)
	tmp = append(tmp, ' ')
	for right < len(tmp) {
		if tmp[right] == ' ' {
			right++
			reverse(tmp[left : right-1])
			left = right
		} else {
			right++
		}
	}
	return string(tmp[:len(tmp)-1])
}

func reverse(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
