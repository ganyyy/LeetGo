package main

import "strings"

var vowel = [128]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func toGoatLatin(sentence string) string {
	var words = strings.Split(sentence, " ")

	var sb strings.Builder
	var count int
	for _, word := range words {
		sb.WriteByte(' ')
		if vowel[word[0]] {
			sb.WriteString(word)
		} else {
			sb.WriteString(word[1:])
			sb.WriteByte(word[0])
		}
		count++
		sb.WriteString("ma")
		sb.WriteString(strings.Repeat("a", count))
	}

	return sb.String()[1:]
}
