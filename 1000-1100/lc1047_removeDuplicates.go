package main

import "unsafe"

func removeDuplicates(S string) string {
	var bs = make([]byte, 0, len(S))

	for i := 0; i < len(S); i++ {
		if len(bs) > 0 && bs[len(bs)-1] == S[i] {
			bs = bs[:len(bs)-1]
		} else {
			bs = append(bs, S[i])
		}
	}

	return *(*string)(unsafe.Pointer(&bs))
}
