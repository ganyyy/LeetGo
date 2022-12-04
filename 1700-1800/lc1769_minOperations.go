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

	// 从左到右的前缀和和累加和
	var sum int
	for l := 1; l < ln; l++ {
		sum += v(l - 1)
		left[l] += sum + left[l-1]
	}

	// 从右到左的前缀和和累加和
	sum = 0
	for r := ln - 2; r >= 0; r-- {
		sum += v(r + 1)
		right[r] += sum + right[r+1]
		// 同时更新结果
		left[r] += right[r]
	}

	return left
}
