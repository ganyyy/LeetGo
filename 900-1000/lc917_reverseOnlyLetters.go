package main

import "fmt"

func reverseOnlyLetters(S string) string {
	res := []byte(S)
	left, right := 0, len(S)-1
	for left <= right {
		lc, rc := check(res[left]), check(res[right])
		if lc && rc {
			res[left], res[right] = res[right], res[left]
			left++
			right--
		} else {
			if !lc {
				left++
			}
			if !rc {
				right--
			}
		}
	}
	return string(res)
}

func reverseOnlyLetters2(s string) string {
	var ret = []byte(s)
	for l, r := 0, len(s)-1; l < r; {
		if !check(s[l]) {
			l++
			continue
		}
		if !check(s[r]) {
			r--
			continue
		}
		ret[l], ret[r] = ret[r], ret[l]
		l++
		r--
	}
	return string(ret)
}

func check(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
}
