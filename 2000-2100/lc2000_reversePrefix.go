package main

import (
	"strings"
	"unsafe"
)

func reversePrefix(word string, ch byte) string {
	var idx = strings.Index(word, string([]byte{ch}))
	if idx == -1 {
		return word
	}
	return reverse(word[:idx+1]) + word[idx+1:]
}

func reverse(s string) string {
	var ret = []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return *(*string)(unsafe.Pointer(&ret))
}
