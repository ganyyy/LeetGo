package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	// e 后面可以跟负数, 不能跟小数
	// -后边不能跟-
	// .后边不能跟
	isE := false     // 指数
	isSmall := false // 小数
	isFlag := false  // 负数
	isDig := false   // 整数
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'e' {
			if !isDig || isE {
				return false
			}
			isE = true
			isSmall = true
			isFlag = false
			isDig = false
		} else if c == '.' {
			if isSmall || isE {
				return false
			}
			isSmall = true
		} else if c == '-' || c == '+' {
			// 前边是数字或者前边是负数或者
			// 不是e开头的情况现是小数
			if isDig || (isSmall && !isE) || isFlag {
				return false
			}
			isFlag = true
		} else if c < '0' || c > '9' {
			return false
		} else {
			isDig = true
			// 只要出现一个数字, 就不能是负数
			isFlag = true
		}
	}
	if (isE || isSmall || isFlag) && !isDig {
		return false
	}
	return true
}

func isNumber2(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	isE := false     // 指数
	isSmall := false // 小数
	isFlag := false  // 负数
	isDig := false   // 整数
	// e 后面可以跟负数, 不能跟小数
	// -后边不能跟-
	// .后边不能跟
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'e', 'E':
			// e前边需要有数字, e后边不能跟小数
			if !isDig || isE {
				return false
			}
			isE = true
			isSmall = true
			// e后边可以有符号和数字
			isFlag = false
			isDig = false
		case '.':
			// 小数点只能在e前边, 并且只能出现一次
			if isSmall || isE {
				return false
			}
			isSmall = true
		case '+', '-':
			// 前边是数字或者前边是负数或者
			// 不是e开头的小数
			if isDig || (isSmall && !isE) || isFlag {
				return false
			}
			isFlag = true
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			isDig = true
			// 只要出现一个数字, 就不能是负数
			isFlag = true
		default:
			// 不在合法范围内的字符
			return false
		}
	}
	if (isE || isSmall || isFlag) && !isDig {
		return false
	}
	return true
}

func main() {
	ss := []string{
		//"0",
		//" 0.1 ",
		//"abc",
		//"1 a",
		//"2e10",
		//" -90e3   ",
		//" 1e",
		//"e3",
		//" 6e-1",
		//" 99e2.5 ",
		//"53.5e93",
		//" --6 ",
		//"-+3",
		//"95a54e53",
		//"+3",
		//".1",
		//".",
		//"+",
		//"-",
		//"e",
		" ",
	}
	for _, s := range ss {
		fmt.Println(s, isNumber(s))
	}
}
