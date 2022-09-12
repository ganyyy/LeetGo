package main

import "fmt"

func sumSubarrayMins(arr []int) int {
	// arr[i]中的数字, 最左边可以到达的位置, 以及最右边可以到达的位置

	var stack []int
	ln := len(arr)
	l, r := make([]int, ln), make([]int, ln)
	// 假设每个位置, 都可以被使用n次
	// l[i] = -1, ∵  i∈[0,n), (i-l[i])∈[1,n]
	for i := 0; i < ln; i++ {
		l[i] = -1
		r[i] = ln
	}

	// 边界值怎么考虑的呢? 其实和切片的原理相同, 左右边界取任意一个就行. 否则会出现少算/多算的情况

	// 正向算r, 计算i可达的最右边位置
	for i := 0; i < ln; i++ {
		for len(stack) != 0 && arr[stack[len(stack)-1]] > arr[i] {
			r[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	// 逆向算l, 计算i可达的最左边位置
	for i := ln - 1; i >= 0; i-- {
		for len(stack) != 0 && arr[stack[len(stack)-1]] >= arr[i] {
			l[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	var ret int
	const MOD int = 1e9 + 7

	for i := 0; i < ln; i++ {
		a, b := l[i], r[i]
		ret = (ret + (i-a)*(b-i)*arr[i]%MOD) % MOD
	}

	fmt.Println(l, r)

	return ret
}

func main() {
	println(sumSubarrayMins([]int{4, 3, 2, 1}))
}
