package main

import "fmt"

type stack struct {
	vals []int
}

func (s *stack) push(b int) {
	s.vals = append(s.vals, b)
}

func (s *stack) pop() (b int) {
	b, s.vals = s.vals[s.len()-1], s.vals[:s.len()-1]
	return
}

func (s *stack) len() int {
	return len(s.vals)
}

func (s *stack) top() int {
	if s.len() == 0 {
		return 0
	}
	return s.vals[s.len()-1]
}

func newStack() *stack {
	return &stack{}
}

func calculate3(s string) int {
	const (
		Add   = '+'
		Sub   = '-'
		Left  = '('
		Right = ')'
		Empty = ' '
		Minus = '-'
	)

	var numStack, charStack = newStack(), newStack()

	var calcVal = func() {
		for charStack.len() != 0 {
			var a, b = numStack.pop(), numStack.pop()
			switch charStack.pop() {
			case Add:
				numStack.push(a + b)
			case Sub:
				numStack.push(b - a)
			}
		}
	}
	var curNum int
	var isNum bool
	var c byte
	for i := 0; i < len(s); i++ {
		// 还是要看括号的
		// 可以忽略掉, 看到一个 - 号直接补一个 0 即可
		c = s[i]
		switch c {
		case Right, Left, Empty:
			continue
		}
		if c >= '0' && c <= '9' {
			curNum *= 10
			curNum += int(c - '0')
			isNum = true
			continue
		}
		if isNum {
			numStack.push(curNum)
			curNum = 0
			isNum = false
		}
		if numStack.len() >= 2 {
			calcVal()
		}
		if c == Minus {
			if numStack.len() != 0 {
				charStack.push(Add)
			}
			numStack.push(0)
		}
		charStack.push(int(c))
	}

	calcVal()

	return numStack.top()
}

func calculate2(s string) int {
	var ss []int

	var res, sign, cur int
	sign = 1
	var c byte
	var ln = len(s)
	for i := 0; i < ln; i++ {
		c = s[i]
		if c >= '0' && c <= '9' {
			cur = int(c - '0')
			for i+1 < ln && s[i+1] >= '0' && s[i+1] <= '9' {
				i++
				c = s[i]
				cur = cur*10 + int(c-'0')
			}
			// 即时计算结果
			res += sign * cur
			continue
		}
		switch c {
		case '+':
			sign = 1
		case '-':
			sign = -1
		case ' ':
		case '(':
			// 出现括号了, 就将之前计算的值/符号入栈, 重新计算结果
			ss = append(ss, res)
			ss = append(ss, sign)
			res = 0
			sign = 1
		case ')':
			if len(ss) >= 2 {
				// 当前值*对应的符号+前一个值对应的符号
				// 相当于最多只会出依次栈
				res = ss[len(ss)-1]*res + ss[len(ss)-2]
				ss = ss[:len(ss)-2]
			}
		}
	}
	return res
}

func calculate5(s string) (ans int) {
	// 一种很新颖的做法: 用栈来保存符号
	// 没有乘除法, 只有加减法所以可以直接计算
	// 这种做法相当于直接把括号拆开了, 将外部的符号直接作用到括号内部
	var ops = []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
			fmt.Println(ops, sign, num, ans)
		}
	}
	return
}

func main() {
	println(calculate5("-2-3-(4-5-6)-1+10"))
}
