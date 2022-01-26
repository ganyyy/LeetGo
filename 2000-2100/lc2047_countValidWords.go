package main

import "strings"

func countValidWords(sentence string) int {
	var words = strings.Split(sentence, " ")

	var validCnt int

	var check = func(b byte) bool {
		return b >= 'a' && b <= 'z'
	}

next:
	for _, w := range words {
		if w == "" {
			continue
		}

		var foundSplit bool
		var foundChar bool
		for i := range w {
			var c = w[i]
			// 连字符
			if c == '-' {
				if foundSplit || i == 0 || i == len(w)-1 {
					continue next
				}
				if !check(w[i-1]) || !check(w[i+1]) {
					continue next
				}

				foundSplit = true
			}
			// 数字
			if c >= '0' && c <= '9' {
				continue next
			}
			// 符号
			if c == '!' || c == '.' || c == ',' {
				if foundChar || i != len(w)-1 {
					continue next
				}
				foundChar = true
			}
		}

		// fmt.Println(w)

		validCnt++
	}

	return validCnt
}
