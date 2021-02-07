package main

func maxTurbulenceSize(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	// 解法1: 摆动序列
	var up, down, res = 1, 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			// 上升趋势, 加上前一个下降的值
			up, down = down+1, 1
		} else if arr[i] < arr[i-1] {
			// 下降趋势, 加上前一个上升的值
			down, up = up+1, 1
		} else {
			// 相等, 直接重置两个值
			up, down = 1, 1
		}
		// 统计一下最大大小
		// 如果直接在循环中遍历最大值, 可以不用考虑边界问题
		res = max(res, max(up, down))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
