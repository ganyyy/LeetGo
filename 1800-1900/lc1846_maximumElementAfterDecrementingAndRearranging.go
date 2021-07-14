package main

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)

	// 必须是1开头
	if arr[0] != 1 {
		arr[0] = 1
	}
	// 两两遍历, 找出差值大于1的数
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i]-arr[i+1]) > 1 {
			arr[i+1] = arr[i] + 1
		}
	}

	return arr[len(arr)-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
