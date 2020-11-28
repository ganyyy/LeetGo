package main

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)

	// 最大组合一定是相邻的
	for i := len(A) - 1; i >= 2; i-- {
		if A[i] < A[i-1]+A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}
