package main

import "math"

func maxProduct(nums []int) int {
	max := math.MinInt32
	// 分别表示当前最大和最小
	iMax, iMin := 1, 1
	for _, v := range nums {
		if v < 0 {
			// 小于0, 最大变最小
			iMin, iMax = iMax, iMin
		}
		// 保证iMax始终都是正数
		iMax = getMax(iMax*v, v)
		// iMin只要数组中存在负数就一定是负数
		iMin = getMin(iMin*v, v)
		// 获取最大的正数
		max = getMax(max, iMax)
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

}
