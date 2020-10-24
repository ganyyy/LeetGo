package main

import (
	"fmt"
	"math"
)

func longestMountain(A []int) int {
	if len(A) < 3 {
		return 0
	}
	// 先是升序的,
	var add = true
	// 起点, 计数点
	var start, i, res int
	// 处理边界值问题, 直接加上一个无法超过的值, 保证在结尾会进行一次判断
	A = append(A, math.MaxInt32)
	for i = 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			start = i
			break
		}
	}
	for ; i < len(A)-1; i++ {
		// 相等的直接跳过, 并重置起点
		if A[i] == A[i+1] {
			// 特殊处理, 防止遗漏
			if i > 0 && A[i-1] > A[i] {
				res = max(res, i-start+1)
			}
			start = math.MaxInt32
			// 如果接下来的数比当前大, 那么起点就是i,
			// 如果接下来的数比当前小, 那么会一直找到对应的拐点

			add = false
			continue
		}
		if A[i] < A[i+1] != add {
			if add {
				// 升序->降序
				add = false
			} else {
				// 降序->升序, 计算最大值
				add = true
				res = max(res, i-start+1)
				start = i
			}
		}
	}
	// 如果正好是一个完整的回文串, 该如何处理?
	if res < 3 {
		return 0
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var cases = [][]int{
		{1, 2, 3, 2, 1, 2, 3, 4, 5, 3, 2, 1},
		{1, 2, 2, 2, 2, 3},
		{1, 2, 2, 2, 2, 3, 2},
		{0, 2, 2},
		{2, 1, 4, 7, 3, 2, 5},
		{2, 2, 2},
		{2, 3, 3, 2, 0, 2},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{0, 0, 1, 0, 0, 1, 1, 1},
	}
	for _, c := range cases {
		fmt.Println(longestMountain(c))
	}
}
