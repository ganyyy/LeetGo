package main

import "unsafe"

func sortString(s string) string {
	if len(s) < 2 {
		return s
	}
	var res = make([]byte, 0, len(s))

	// 统计各个字符的数量
	var t = [26]int{}
	for i := 0; i < len(s); i++ {
		t[s[i]-'a']++
	}

	// 很多无效的遍历, 可以进行优化

	// 填充字符
	for len(res) < len(s) {
		// 正序
		for i := 0; i < 26; i++ {
			if t[i] != 0 {
				t[i]--
				res = append(res, byte(i)+'a')
			}
		}
		// 倒序
		for i := 25; i >= 0; i-- {
			if t[i] != 0 {
				t[i]--
				res = append(res, byte(i)+'a')
			}
		}
	}

	return toString(res)
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
