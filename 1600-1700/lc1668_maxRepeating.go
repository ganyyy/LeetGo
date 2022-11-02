package main

import "strings"

func maxRepeating(sequence string, word string) int {
	var cnt int
	for {
		if strings.Index(sequence, strings.Repeat(word, cnt+1)) != -1 {
			cnt++
		} else {
			break
		}
	}
	return cnt
}
