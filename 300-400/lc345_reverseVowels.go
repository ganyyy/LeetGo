package main

import "unsafe"

func reverseVowels(s string) string {
	var bs = []byte(s)

	var l, r = 0, len(bs) - 1

	var check = func(i byte) bool {
		return i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' || i == 'A' || i == 'E' || i == 'I' || i == 'O' || i == 'U'
	}

	for l < r {
		for l < r && !check(bs[l]) {
			l++
		}
		for l < r && !check(bs[r]) {
			r--
		}
		bs[l], bs[r] = bs[r], bs[l]
		l++
		r--
	}

	return *(*string)(unsafe.Pointer(&bs))
}
