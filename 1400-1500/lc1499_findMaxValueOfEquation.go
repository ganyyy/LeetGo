//go:build ignore

package main

import "math"

func findMaxValueOfEquation(points [][]int, k int) int {
	// y-x 的最大值
	var stack []int // 递减队列, y-x的最大值
	var ret = math.MinInt32
	for idx, point := range points {
		// 去除头部
		for ; len(stack) > 0 && point[0]-points[stack[0]][0] > k; stack = stack[1:] {
		}
		// 维护队列
		val := calcVal(point)
		for len(stack) > 0 {
			ln := len(stack)
			last := points[stack[ln-1]]
			if calcVal(last) <= val {
				ret = max(ret, calcRet(last, point))
				stack = stack[:ln-1]
			} else {
				break
			}
		}
		// 再次计算结果
		if len(stack) > 0 {
			ret = max(ret, calcRet(points[stack[0]], point))
		}
		stack = append(stack, idx)
	}
	return ret
}

func calcVal(a []int) int {
	return a[1] - a[0]
}

func calcRet(a, b []int) int {
	return a[1] + b[1] + abs(a[0]-b[0])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
