package main

import (
	"math"
)

func find132pattern(nums []int) bool {
	// 需要使用一个栈处理嘛?
	// 需要记忆化搜索了.
	// 暴力解法, 过了一次
	// 可以理解为, 去掉当前字符后, 前半段和后半段是否存在一对数, 满足 前 < 后 < 当前值

	// 第一趟, 确定左边的最小值
	var tmp = make([]int, len(nums))
	tmp[0] = math.MaxInt32
	for i := 1; i < len(nums)-1; i++ {
		tmp[i] = min(tmp[i-1], nums[i-1])
	}
	// 根据获取的最小数组, 从右向左依次判断是否存在满足条件的值

	// 右半部分需要找到小于当前值的最大值...?
	for i := len(nums) - 2; i >= 1; i-- {
		var minMax = math.MinInt32
		// 这一块, 怎么优化更合适呢?
		for j := i + 1; j < len(nums); j++ {
			if nums[j] >= nums[i] {
				continue
			}
			minMax = max(minMax, nums[j])
		}
		if minMax > tmp[i] {
			return true
		}
	}

	return false
}

func find132patternStack(nums []int) bool {
	//
	// 栈顶保存的是3
	var stack []int
	// minVal保存的是2
	var minVal = math.MinInt32

	for i := len(nums) - 1; i >= 0; i-- {
		// nums[i]保存的是1
		if minVal > nums[i] {
			return true
		}
		// 这个栈是一个单调递增, 栈顶元素最小. 所以所有的出栈的元素都可以作为2的备选
		for len(stack) != 0 && stack[len(stack)-1] < nums[i] {
			minVal = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	println(find132pattern([]int{-2, 1, 2, -2, 1, 2}))
}
