package main

import "strconv"

func numDecodings(s string) int {
	// 去掉前导0
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	// len(s) >= 2
	var a, b int
	a = 1 // a一定为1

	// 检查两位有几种可能
	if checkValid(s[:2]) {
		// 只有当两个都合法, 并且第二位也合法时, 才为2
		if checkValid(s[1:2]) {
			b = 2
		} else {
			b = 1
		}
	} else {
		// 如果能合并拼成一个数字, 也算合法
		if checkValid(s[:1]) && checkValid(s[1:2]) {
			b = 1
		}
	}

	for i := 2; i < len(s); i++ {
		var tmp int
		if checkValid(s[i-1 : i+1]) {
			tmp += a
		}
		if checkValid(s[i : i+1]) {
			tmp += b
		}
		a, b = b, tmp
	}
	return b
}

func numDecodings2(s string) int {
	// 去掉前导0
	if s[0] == '0' {
		return 0
	}

	// dfs处理, 计算
	var res int
	var dfs func(s string)

	dfs = func(s string) {
		if len(s) == 0 {
			res++
		} else if len(s) == 1 {
			if checkValid(s) {
				res++
			}
		} else if len(s) == 2 {
			if checkValid(s) {
				res++
				if checkValid(s[1:]) {
					res++
				}
			} else {
				if checkValid(s[:1]) && checkValid(s[1:]) {
					res++
				}
			}
		} else {
			if checkValid(s[:1]) {
				dfs(s[1:])
			}
			if checkValid(s[:2]) {
				dfs(s[2:])
			}
		}
	}

	dfs(s)

	return res
}

func checkValid(s string) bool {
	// s只能长度为1或者2
	var ret, _ = strconv.Atoi(s)
	if ret == 0 {
		return false
	}
	if ret > 26 {
		return false
	}
	if ret < 10 && len(s) == 2 {
		return false
	}
	return true
}
