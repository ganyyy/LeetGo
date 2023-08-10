package main

func trap(height []int) int {
	// 印象中可以通过双指针处理
	var left, right = 0, len(height) - 1
	var lMax, rMax int
	var res int
	for left < right {
		if height[left] < height[right] {
			if lMax > height[left] {
				res += lMax - height[left]
			} else {
				lMax = height[left]
			}
			left++
		} else {
			if rMax > height[right] {
				res += rMax - height[right]
			} else {
				rMax = height[right]
			}
			right--
		}
	}
	return res
}
