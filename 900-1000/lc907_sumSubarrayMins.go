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

func sumSubarrayMins2(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	// 单调栈, 👉
	var monoStack []int
	// [1,7,5,2,4,3] 从右向左看, 每当最小值发生变化时, 对应的区间 [1,2,2,2,3,3]
	// [1,7,5,2,4,3] -> [1]
	// [7,5,2,4,3] 	 -> [2]
	// [5,2,4,3]	 -> [2]
	// [2,4,3]	 	 -> [2]
	// [4,3]	 	 -> [3]
	// [3]	 	 	 -> [3]

	// dp[i] arr[:i+1]的子数组的最小值之和
	dp := make([]int, n)
	for i, x := range arr {
		for len(monoStack) > 0 && arr[monoStack[len(monoStack)-1]] > x {
			monoStack = monoStack[:len(monoStack)-1]
		}
		// k代表着以x为最小值的子数组所能延伸的最远区间
		k := i + 1
		if len(monoStack) > 0 {
			// 如果存在前置元素, 那么出现的次数就会相应的减少
			k = i - monoStack[len(monoStack)-1]
		}
		dp[i] = k * x
		if len(monoStack) > 0 {
			dp[i] += dp[i-k]
		}
		ans = (ans + dp[i]) % mod
		monoStack = append(monoStack, i)
	}
	return
}

func sumSubarrayMins3(arr []int) int {
	const MOD = 1e9 + 7

	// 哨兵值, 保证末尾的计算
	arr = append(arr, -1)
	var ret int

	var stack = []int{-1}

	for right, val := range arr {
		for last := len(stack) - 1; last >= 1 && arr[stack[last]] >= val; last-- {
			idx := stack[last]
			stack = stack[:last]
			// arr[idx]左边有idx-stack[last-1]个元素, 右边有right-idx个元素
			// 在这个范围内, arr[idx]是最小值
			ret = (ret + arr[idx]*(idx-stack[last-1])*(right-idx)) % MOD
		}
		stack = append(stack, right)
	}
	return ret
}

func main() {
	println(sumSubarrayMins([]int{4, 3, 2, 1}))
}
