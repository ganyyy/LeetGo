package main

import "unsafe"

func reverseStr(s string, k int) string {
	var bs = []byte(s)

	var reverse = func(bb []byte) {
		for l, r := 0, len(bb)-1; l < r; l, r = l+1, r-1 {
			bb[l], bb[r] = bb[r], bb[l]
		}
	}

	for i := 0; i < len(s); i += k * 2 {
		var tmp = k
		if i+k > len(s) {
			tmp = len(s) - i
		}
		reverse(bs[i : i+tmp])
	}

	return *(*string)(unsafe.Pointer(&bs))
}
