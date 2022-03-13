package main

import "unsafe"

// 为啥最高是36进制呢? 因为 26个英文字母+10个数字 = 36

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var ret []byte

	// 注意一下负数的情况
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
