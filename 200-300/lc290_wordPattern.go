package main

import "strings"

func wordPattern(pattern string, s string) bool {
	var ps = strings.Split(s, " ")
	if len(pattern) != len(ps) {
		return false
	}

	var set [26]string
	var sets = make(map[string]byte, 26)

	for i := 0; i < len(pattern); i++ {
		var idx = pattern[i] - 'a'
		if set[idx] == "" {
			set[idx] = ps[i]
			if sets[ps[i]] != 0 {
				return false
			}
			sets[ps[i]] = pattern[i]
		} else if set[idx] != ps[i] {
			return false
		}
	}

	return true
}
