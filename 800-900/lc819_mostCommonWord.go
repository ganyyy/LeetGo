package main

import "strings"

var symbol = func() (arr [128]bool) {
	var str = "!?',;. "
	for i := range str {
		arr[str[i]] = true
	}
	return
}()

func mostCommonWord(paragraph string, banned []string) string {
	var checkIsBanned = func(word string) bool {
		// fmt.Println(word)
		for _, ban := range banned {
			if ban == word {
				return true
			}
		}
		return false
	}
	var cnt = make(map[string]int)
	var addWord = func(word string) {
		word = strings.ToLower(word)
		if checkIsBanned(word) {
			return
		}
		cnt[word]++
	}

	var start int
	for i := range paragraph {
		var ch = paragraph[i]
		if !symbol[ch] {
			continue
		}
		if start == i {
			start++
			continue
		}
		addWord(paragraph[start:i])
		start = i + 1
	}

	if start != len(paragraph)-1 {
		addWord(paragraph[start:])
	}

	var ret string
	var max int
	for word, count := range cnt {
		if count > max {
			ret = word
			max = count
		}
	}
	return ret
}
