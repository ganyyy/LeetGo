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

func check(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
}
