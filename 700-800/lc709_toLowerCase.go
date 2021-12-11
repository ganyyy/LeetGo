package main

import "unsafe"

func toLowerCase(s string) string {
	var ret = make([]byte, len(s))
	for i := range s {
		var c = s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		ret[i] = c
	}
	return *(*string)(unsafe.Pointer(&ret))
}
