//go:build ignore

package main

func minOperations(boxes string) []int {
	// 任意一个点, O(1)时间内获取左右的前缀和
	ln := len(boxes)
	if ln < 1 {
		return nil
	}
	var left = make([]int, ln)
	var right = make([]int, ln)

	v := func(i int) int {
		if i >= ln || i < 0 {
			return 0
		}
		return int(boxes[i] - '0')
	}

	for l := 1; l < ln; l++ {
		left[l] += v(l-1) + left[l-1]
	}
	for r := ln - 2; r >= 0; r-- {
		right[r] += v(r+1) + right[r+1]
	}
	for l := 1; l < ln; l++ {
		left[l] += left[l-1]
	}
	for r := ln - 2; r >= 0; r-- {
		right[r] += right[r+1]
		left[r] += right[r]
	}

	return left
}
