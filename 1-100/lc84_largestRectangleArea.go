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
	// 通过push 0 进入heights可以保证一个循环内走完
	//heights = append(heights, 0)
	stack := make([]int, 0, len(heights)+1)
	stack = append(stack, -1)
	mArea := 0
	for i := 0; i < len(heights); i++ {
		// 如果栈顶不为空(-1), 并且栈顶对应的值值比当前值要大, 那么就计算一波最大值
		for t := len(stack) - 1; stack[t] != -1 && heights[stack[t]] >= heights[i]; t = len(stack) - 1 {
			// 长度为栈顶, 宽度为当前位置到栈顶索引的位置-1
			// 栈是[-1, 0, 1], 当前位置是2, 长度为 2-0-1=1
			// 出栈后变成[-1, 0], 当前位置依旧是2, 长度为2-(-1)-1=2
			// 继续出栈, 栈变成[-1]或者height[stack.top] < v时终止循环, 保证栈内元素对应的高是递增的
			mArea = max(mArea, heights[stack[t]]*(i-stack[t-1]-1))
			// 栈顶出栈
			stack = stack[:t]
		}
		// 当前索引入栈
		stack = append(stack, i)
	}
	// 最后还需要判断一下栈是否为空, 因为站内元素对应的高一定是递增的, 所以可以进行
	for t := len(stack) - 1; stack[t] != -1; t = len(stack) - 1 {
		// 栈内剩下的元素一定是递增关系, 并且当前相当于走完了全部的循环, 相当于终止索引为height.length
		mArea = max(mArea, heights[stack[t]]*(len(heights)-stack[t-1]-1))
		// 计算完成后出栈
		stack = stack[:t]
	}
	return mArea
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
