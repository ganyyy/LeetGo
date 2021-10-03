package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	// 长除法?

	var a, b = numerator, denominator

	if a%b == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	var sb strings.Builder

	if a < 0 && b > 0 || a > 0 && b < 0 {
		sb.WriteString("-")
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	sb.WriteString(strconv.Itoa(a/b) + ".")

	a %= b

	var m = make(map[int]int)
	// 长除法就是不停的x10 取余
	for a != 0 {
		m[a] = sb.Len()
		a *= 10
		sb.WriteString(strconv.Itoa(a / b))

		// 重新计算a整出后的余数
		a %= b
		// 如果出现过, 这就是一个循环节
		if idx, ok := m[a]; ok {
			// 出现了循环节
			return fmt.Sprintf("%s(%s)", sb.String()[:idx], sb.String()[idx:])
		}
	}

	return sb.String()
}
