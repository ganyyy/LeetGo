package main

import "sort"

/**
给你一个整数数组 arr 和一个目标值 target ，请你返回一个整数 value ，
使得将数组中所有大于 value 的值变成 value 后，数组的和最接近  target （最接近表示两者之差的绝对值最小）。

如果有多种使得和最接近 target 的方案，请你返回这些整数中的最小值。

请注意，答案不一定是 arr 中的数字。

*/

func findBestValue(arr []int, target int) int {
	// 先排序
	sort.Ints(arr)

	var sum int
	var size = len(arr)
	for i, v := range arr {
		x := (target - sum) / (size - i)
		if x < v {
			var t = float64(target-sum) / float64(size-i)
			if t-float64(x) > 0.5 {
				return x + 1
			} else {
				return x
			}
		}
		sum += v
	}
	return arr[size-1]
}
