package main

import "unsafe"

func isAlienSorted(words []string, order string) bool {
	var dic [26]byte
	for i := 0; i < 26; i++ {
		dic[order[i]-'a'] = byte(i)
	}

	var translate = func(b string) string {
		var bs = make([]byte, len(b))
		for i := range b {
			bs[i] = dic[b[i]-'a'] + 'a'
		}
		return *(*string)(unsafe.Pointer(&bs))
	}

	for i := 0; i < len(words)-1; i++ {
		if translate(words[i]) > translate(words[i+1]) {
			return false
		}
	}
	return true
}
