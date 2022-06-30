package main

import "strconv"

func diffWaysToCompute(expression string) []int {
	var nums []int
	var length = len(expression)

	for i := 0; i < length; i++ {
		var c = expression[i]
		if c == '+' || c == '-' || c == '*' {
			subA := diffWaysToCompute(expression[:i])
			subB := diffWaysToCompute(expression[i+1:])
			for _, va := range subA {
				for _, vb := range subB {
					if c == '+' {
						nums = append(nums, va+vb)
					} else if c == '-' {
						nums = append(nums, va-vb)
					} else {
						nums = append(nums, va*vb)
					}
				}
			}
		}
	}
	if len(nums) == 0 {
		var v, _ = strconv.Atoi(expression)
		nums = append(nums, v)
	}
	return nums
}
