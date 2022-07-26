package main

import (
	"fmt"
	"unicode"
)

func fractionAddition(expression string) string {
	denominator, numerator := 0, 1 // 分子，分母
	// 不断地读取分子/分母, 而后累计计算公倍数
	for i, n := 0, len(expression); i < n; {
		// 读取分子
		denominator1, sign := 0, 1
		if expression[i] == '-' || expression[i] == '+' {
			if expression[i] == '-' {
				sign = -1
			}
			i++
		}
		for i < n && unicode.IsDigit(rune(expression[i])) {
			denominator1 = denominator1*10 + int(expression[i]-'0')
			i++
		}
		denominator1 = sign * denominator1
		i++

		// 读取分母
		numerator1 := 0
		for i < n && unicode.IsDigit(rune(expression[i])) {
			numerator1 = numerator1*10 + int(expression[i]-'0')
			i++
		}

		// 分子交叉相乘再相加
		denominator = denominator*numerator1 + denominator1*numerator
		// 分母直接乘
		numerator *= numerator1
	}
	if denominator == 0 {
		return "0/1"
	}
	// 计算最大公约数, 对分子/分母进行相约
	g := gcd(abs(denominator), numerator)
	return fmt.Sprintf("%d/%d", denominator/g, numerator/g)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
