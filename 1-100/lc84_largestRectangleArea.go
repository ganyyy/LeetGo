package main

import (
	"fmt"
)

// 分治法
func largestRectangleArea(heights []int) int {
	var search func(left, right int) int
	search = func(left, right int) int {
		if left > right {
			// 没有宽度, 面积返回0
			return 0
		}
		mIndex := left
		for i := left + 1; i <= right; i++ {
			if heights[i] < heights[mIndex] {
				mIndex = i
			}
		}
		// 左右中找最小值
		// 1 包含自己在内的
		// 2 左边计算的结果
		// 3 右边计算的结果
		return max((right-left+1)*heights[mIndex], max(search(left, mIndex-1), search(mIndex+1, right)))
	}
	return search(0, len(heights)-1)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 堆栈法
func largestRectangleArea2(heights []int) int {
	// 栈中存放的是比当前索引高度要小的 索引, 如果比当前高度高, 那就选择出栈
	stack := make([]int, 0, len(heights)+1)
	// 保证添加的首个索引比后序的都小, 那就永远不会出栈
	stack = append(stack, -1)
	var mArea int

	//   [2, 1, 5, 6, 2, 3] 的矩形 数组 想象成
	// [-1, 2, 1, 5, 6, 2, 3]  几条边构成 的 图形
	// 比如当前循环到了 i=4, v=2
	// 此时栈中的元素为 [-1, 1, 2, 3]
	// heights[3] >= 2, area = heights[3] * (4- [2] -1) = 6*1, 3 出栈
	// heights[2] >= 2, area = heights[2] * (4- [1] -1) = 5*2, 2 出栈
	// heights[1] <  2, break

	for i, v := range heights {
		// 保证栈内元素都是一直递增的, 如果出现比当前更高的, 就出栈
		for t := len(stack) - 1; t != 0 && heights[stack[t]] >= v; t-- {
			// 计算矩阵中的最大高度
			mArea = max(mArea, heights[stack[t]]*(i-stack[t-1]-1))
			stack = stack[:t]
		}
		stack = append(stack, i)
	}

	// 预防 一直递增的问题, 此时的宽度就是 整个数组的宽度
	for t, h := len(stack)-1, len(heights); t != 0; t-- {
		mArea = max(mArea, heights[stack[t]]*(h-stack[t-1]-1))
	}
	return mArea
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
