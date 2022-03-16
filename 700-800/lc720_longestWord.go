package main

import "sort"

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	var set = make(map[string]bool, len(words))
	// for _, s := range words {
	//     set[s] = true
	// }

	var ret string

	// fmt.Println(words)

	for _, s := range words {
		if len(s) != 1 && !set[s[:len(s)-1]] {
			continue
		}
		set[s] = true
		if len(s) > len(ret) || (len(s) == len(ret) && s < ret) {
			ret = s
		}
	}

	return ret
}
