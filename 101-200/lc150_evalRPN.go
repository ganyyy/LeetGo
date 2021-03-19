package main

import "strconv"

func evalRPN(tokens []string) int {
	// 需要把整体当成一个栈进行处理
	// 可以用于表达式求值
	// 这个有点意思. 需要记录一下

	var stack = make([]int, 0, len(tokens)>>1)

	for _, v := range tokens {
		switch v {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
		default:
			val, _ := strconv.Atoi(v)
			stack = append(stack, val)
			continue
		}
		stack = stack[:len(stack)-1]
	}
	return stack[0]
}
