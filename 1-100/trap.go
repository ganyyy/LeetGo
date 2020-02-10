package main

import "fmt"

func trap(height []int) int {
	lh := len(height)
	maxLeft, maxRight := make([]int, lh), make([]int, lh)
	maxLeft[0] = height[0]
	maxRight[lh-1] = height[lh-1]
	// 构建左边最大和右边最大的数组
	for i := 1; i < lh; i++ {
		maxLeft[i] = min(height[i], maxLeft[i-1])
	}
	for i := lh - 2; i >= 0; i-- {
		maxRight[i] = max(height[i], maxRight[i+1])
	}
	fmt.Println(maxLeft, maxRight)
	res := 0
	// 找差值
	for i := 0; i < lh; i++ {
		if min(maxLeft[i], maxRight[i]) > height[i] {
			res += min(maxLeft[i], maxRight[i]) - height[i]
		}
	}
	return res
}

func trap2(height []int) int {
	lh := len(height)
	if lh < 1 {
		return 0
	}
	left, right := 0, lh-1
	maxLeft, maxRight := height[left], height[right]
	ans := 0

	for left < right {
		maxLeft = max(height[left], maxLeft)
		maxRight = max(height[right], maxRight)
		// 取最小边计算差值, 双指针向中间缩
		if maxLeft < maxRight {
			ans += maxLeft - height[left]
			left++
		} else {
			ans += maxRight - height[right]
			right--
		}
	}
	return ans
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
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
