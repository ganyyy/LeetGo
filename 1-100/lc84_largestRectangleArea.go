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

// 堆栈法
func largestRectangleArea2(heights []int) int {
	// 栈中存放的是比当前索引高度要小的 索引, 如果比当前高度高, 那就选择出栈
	stack := make([]int, 0, len(heights)+1)
	// 保证添加的首个索引比后序的都小, 那就永远不会出栈
	stack = append(stack, -1)
	var mArea int

	// Mark

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

func largestRectangleAreaNew(heights []int) int {
	// 单调递增栈处理
	var stack = make([]int, 1, len(heights)+1)
	stack[0] = -1

	var res int
	for i, v := range heights {
		// 比如 [2, 1, 5, 6, 2, 3]
		// 栈中元素对应的数值如下:
		// [2]
		// [1].       出栈2*1
		// [1, 5]
		// [1, 5, 6]
		// [1, 2]. 	  出栈6*1, 再出栈5*2
		// [1, 2, 3]. 出栈3*1, 2*2, 1*5
		for t := len(stack) - 1; t != 0 && heights[stack[t]] >= v; t-- {
			// 长度就是栈中的最后一个数值
			// 宽度就是 i 距离 栈顶前一个元素的距离
			// 可以这么理解: 矩形的高度相当于是 heights[stack[t]]和 heights[i-1]中的较小值
			// 但是在i之前, stack是单调递增的, 所以 heights[stack[t]] < heights[i-1], 因此可以直接使用 heights[stack[t]]
			// 宽度的起始值是 stack[t-1]+1, 因为 heights[stack[t-1]]是比 heights[stack[t]]小的, 所以宽度至少是1
			res = max(res, heights[stack[t]]*(i-(stack[t-1])+1))
			stack = stack[:t]
		}
		stack = append(stack, i)
	}

	// 计算剩余的递增栈
	for t, h := len(stack)-1, len(heights); t != 0; t-- {
		res = max(res, heights[stack[t]]*(h-stack[t-1]-1))
	}

	return res
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
