package main

import "fmt"

func calculate(s string) int {
	// 加减保存
	var numStack []int
	var charStack []byte

	const (
		Empty = ' '
		Add   = '+'
		Sub   = '-'
		Mul   = '*'
		Div   = '/'
	)

	// 乘除计算
	var curNum int
	var isNum bool

	var calc = func() {
		if len(numStack) < 2 {
			return
		}
		var b = numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		switch charStack[len(charStack)-1] {
		case Add:
			numStack[len(numStack)-1] += b
		case Sub:
			numStack[len(numStack)-1] -= b
		case Mul:
			numStack[len(numStack)-1] *= b
		case Div:
			numStack[len(numStack)-1] /= b
		}
		charStack = charStack[:len(charStack)-1]
	}

	for i := 0; i < len(s); i++ {
		var c = s[i]
		if c == Empty {
			continue
		}
		if c >= '0' && c <= '9' {
			curNum = curNum*10 + int(c-'0')
			isNum = true
			continue
		}
		// 到这里说明不是数字了, 肯定是某个负号之一
		numStack = append(numStack, curNum)
		curNum = 0
		isNum = false

		// 根据优先级, 进行提前计算
		for len(charStack) != 0 {
			var tc = charStack[len(charStack)-1]
			if tc == Mul || tc == Div || c == Add || c == Sub {
				calc()
			} else {
				break
			}
		}
		charStack = append(charStack, c)
	}

	if isNum {
		numStack = append(numStack, curNum)
	}

	// 最后清空栈
	for len(numStack) >= 2 {
		calc()
	}

	return numStack[0]
	// 明天就要出带括号的计算器了吧...
}

func calculateSimply(s string) int {
	const (
		Empty = ' '
		Add   = '+'
		Sub   = '-'
		Mul   = '*'
		Div   = '/'
	)
	// 默认是正数
	var sign byte = Add
	// 数栈
	var numStack []int
	var cur int
	var c byte
	for i := 0; i < len(s); i++ {
		// 符号, 空格的ASCII都小于 '0'
		c = s[i]
		if c >= '0' {
			cur = cur*10 + int(c-'0')
		}
		if (c < '0' && c != Empty) || i == len(s)-1 {
			switch sign {
			case Add:
				numStack = append(numStack, cur)
			case Sub:
				numStack = append(numStack, -cur)
			case Mul:
				numStack[len(numStack)-1] *= cur
			case Div:
				numStack[len(numStack)-1] /= cur
			}
			sign = c
			cur = 0
		}
	}

	for i := len(numStack) - 1; i >= 0; i-- {
		cur += numStack[i]
	}
	return cur
}

func main() {

	var cases = []string{
		"3+2*2",
		"3+2*21",
		"3+2/2",
		"3*2*2",
		"3-       2         *2        ",
		"1-1+1",
		"1*2-3/4+5*6-7*8+9/10",
	}

	for _, s := range cases {
		fmt.Println(calculate(s))
	}
}
