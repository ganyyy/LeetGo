package main

import "sort"

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) != len(dictionary[j]) {
			return len(dictionary[i]) > len(dictionary[j])
		}
		return dictionary[i] < dictionary[j]
	})

	for _, d := range dictionary {
		// 双指针匹配啊...
		var l int
		for i := range s {
			if s[i] == d[l] {
				l++
				if l == len(d) {
					return d
				}
			}
		}
	}

	return ""
}
