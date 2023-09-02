package main

import "math"

func maxProduct(nums []int) int {
	ret := math.MinInt32
	// 分别表示当前最大和最小
	iMax, iMin := 1, 1
	for _, v := range nums {
		if v < 0 {
			// 小于0, 最大变最小
			iMin, iMax = iMax, iMin
		}
		// 保证iMax始终都是正数
		iMax = max(iMax*v, v)
		// iMin只要数组中存在负数就一定是负数
		iMin = min(iMin*v, v)
		// 获取最大的正数
		ret = max(ret, iMax)
	}
	return ret
}

func main() {

}
