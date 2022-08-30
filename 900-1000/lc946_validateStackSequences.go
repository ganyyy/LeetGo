package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	var stack = make([]int, 0, len(pushed))

	var start int

	var show = func() {
		fmt.Printf("stack:%v, start:%v, pushed:%v, popped:%v\n", stack, start, pushed[start:], popped)
	}

	for {
		// 要么一直入栈
		for ; start < len(pushed) && pushed[start] != popped[0]; start++ {
			stack = append(stack, pushed[start])
		}
		// 此时入栈栈顶和出栈头相同, 这两个直接出栈(不入栈就等同于出栈)
		start++
		popped = popped[1:]
		// 要么弹出栈顶
		for len(stack) > 0 && len(popped) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		}
		if start >= len(pushed) || len(popped) == 0 {
			break
		}
	}

	_ = show
	return len(stack) == 0 && start == len(pushed) && len(popped) == 0
}
