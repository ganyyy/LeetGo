package main

import (
	"strconv"
	"strings"
)

func solveEquation(equation string) string {

	// 分割常量和系数
	var parse = func(s string) ([]int, []int) {
		var minus bool
		var n, x []int
		for i := 0; i < len(s); {
			// 计算符号位
			if s[i] == '-' || s[i] == '+' {
				minus = s[i] == '-'
				i++
				continue
			}
			// 计算数字位
			var cur int
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				cur = 10*cur + int(s[i]-'0')
				i++
			}
			if minus {
				cur = -cur
			}
			// 获取对应的数字
			if i < len(s) && s[i] == 'x' {
				if i == 0 || s[i-1] == '+' || s[i-1] == '-' {
					// 特殊处理x不存在前缀数字的情况
					cur = 1
					if minus {
						cur = -1
					}
				}
				i++
				x = append(x, cur)
			} else {
				n = append(n, cur)
			}
		}
		return n, x
	}

	var sp = strings.Split(equation, "=")
	var n1, x1 = parse(sp[0])
	var n2, x2 = parse(sp[1])

	var n, x int
	for _, v := range n1 {
		n += v
	}
	for _, v := range n2 {
		n -= v
	}
	for _, v := range x1 {
		x -= v
	}
	for _, v := range x2 {
		x += v
	}

	if x == 0 && n == 0 {
		// 行如 x = x
		return "Infinite solutions"
	} else if x == 0 && n != 0 {
		// 行如 x = x+1
		return "No solution"
	}
	// 正常计算
	return "x=" + strconv.Itoa(n/x)
}
