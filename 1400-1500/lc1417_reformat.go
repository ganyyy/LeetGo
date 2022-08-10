package main

import "unsafe"

func reformat(s string) string {
	if len(s) <= 1 {
		return s
	}
	var aa []byte
	var ia []byte
	for idx := 0; idx < len(s); idx++ {
		if s[idx] >= '0' && s[idx] <= '9' {
			ia = append(ia, s[idx])
		} else {
			aa = append(aa, s[idx])
		}
	}
	if abs(len(aa)-len(ia)) > 1 {
		return ""
	}
	var ret = make([]byte, 0, len(s))
	// 交替整合
	if len(aa) < len(ia) {
		aa, ia = ia, aa
	}
	var i1 int
	for ; i1 < len(ia); i1++ {
		ret = append(ret, aa[i1], ia[i1])
	}
	if i1 < len(aa) {
		ret = append(ret, aa[i1])
	}
	return *(*string)(unsafe.Pointer(&ret))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
