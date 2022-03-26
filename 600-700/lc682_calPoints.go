package main

import "strconv"

func calPoints(ops []string) int {
	var stack []int

	for _, op := range ops {
		switch op {
		case "+":
			if len(stack) < 2 {
				continue
			}
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			if len(stack) < 1 {
				continue
			}
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			if len(stack) < 1 {
				continue
			}
			stack = stack[:len(stack)-1]
		default:
			var val, err = strconv.Atoi(op)
			if err != nil {
				continue
			}
			stack = append(stack, val)
		}
	}
	var ret int
	for _, v := range stack {
		ret += v
	}
	return ret
}
