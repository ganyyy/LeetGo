package main

func minArray(numbers []int) int {
	// 当然可以遍历进行处理

	// 为啥不用二分呢?
	left, right := 0, len(numbers)-1

	for left < right {
		mid := left + (right-left)>>1
		if numbers[mid] < numbers[right] {
			// 中间小于右边, 则[mid,right]是有序的, 最小值一定在 [left, mid]之间
			right = mid
		} else if numbers[right] < numbers[mid] {
			// 中间大于右边, 则[mid, right]中间有拐点
			left = mid + 1
		} else {
			// 相等的话, 不能确定是左边还是右边, -- 去重
			right--
		}
	}
	return numbers[left]
}

func minArray2(numbers []int) int {
	var left, right = 0, len(numbers) - 1

	for left < right {
		var mid = left + (right-left)/2
		var _, r, m = numbers[left], numbers[right], numbers[mid]

		// 确定递增的区间
		if r > m {
			right = mid
		} else if r < m {
			left = mid + 1
		} else {
			right--
		}
	}

	return numbers[left]
}

func main() {

}
