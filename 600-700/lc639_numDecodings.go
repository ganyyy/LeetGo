package main

func check1digit(ch byte) int {
	if ch == '*' {
		// 1-9
		return 9
	}
	if ch == '0' {
		// 0不是有效的编码
		return 0
	}
	return 1
}

func check2digits(c0, c1 byte) int {
	//11-19, 21-26
	if c0 == '*' && c1 == '*' {
		return 15
	}
	if c0 == '*' {
		if c1 <= '6' {
			//1, 2
			return 2
		}
		//1
		return 1
	}
	if c1 == '*' {
		if c0 == '1' {
			// 11-19
			return 9
		}
		if c0 == '2' {
			// 21-26
			return 6
		}
		// 不匹配
		return 0
	}
	// 无前导0, 且必须是一个合法的数字. 直接组合
	if c0 != '0' && (c0-'0')*10+(c1-'0') <= 26 {
		return 1
	}
	return 0
}

func numDecodings(s string) int {
	const mod int = 1e9 + 7
	// 具体做法和91题差不多. 当前态只和前态/前前态相关
	a, b, c := 0, 1, 0
	for i := range s {
		// 至于为啥是乘法? 因为可以自由组合捏😊

		// 取一位的情况
		c = b * check1digit(s[i]) % mod
		if i > 0 {
			// 取两位的情况
			c = (c + a*check2digits(s[i-1], s[i])) % mod
		}
		a, b = b, c
	}
	return c
}
