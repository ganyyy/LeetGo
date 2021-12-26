package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	var ret []string
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			ret = append(ret, words[i])
		}
	}
	return ret
}

func findOcurrencesBad(text string, first string, second string) []string {
	var match = first + " " + second
	var ln = len(match)
	var lt = len(text)
	var ret []string
	for i := 0; i < lt-ln; {
		if text[i] == ' ' {
			i++
			continue
		}
		if text[i:i+ln] != match {
			for i < lt && text[i] != ' ' {
				i++
			}
			continue
		}
		i = i + ln
		if text[i] != ' ' {
			for i < lt && text[i] != ' ' {
				i++
			}
			continue
		}
		i++
		var end = i
		for end < lt && text[end] != ' ' {
			end++
		}

		ret = append(ret, text[i:end])

		// 前置推导一下
		var t int
		for t = i - 2; t > 0 && text[t] != ' '; t-- {
		}

		if text[t+1:i-1] == first {
			i = t + 1
		}
	}
	return ret
}
