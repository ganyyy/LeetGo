package main

import "unsafe"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var ret []byte

	var neg bool
	if num < 0 {
		neg = true
		num = -num
	}

	for num > 0 {
		ret = append(ret, '0'+byte(num%7))
		num /= 7
	}
	if neg {
		ret = append(ret, '-')
	}
	for l, r := 0, len(ret)-1; l < r; l, r = l+1, r-1 {
		ret[l], ret[r] = ret[r], ret[l]
	}
	return *(*string)(unsafe.Pointer(&ret))
}
