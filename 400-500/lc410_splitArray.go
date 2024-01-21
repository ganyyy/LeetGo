package main

// 值二分法的应用
func splitArray(nums []int, m int) int {
	// 咸鱼到 连DP 都不会

	var max, sum int
	// 1. 先看长度和区间相等的情况
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if len(nums) == m {
		return max
	}
	// 剩下的基于2分进行处理

	// max 表示 一个数组一组时的最大值, sum表示全部数组在一个组时的最大值
	for max < sum {
		mid := max + (sum-max)>>1

		// temp 表示子数组的和, cnt表示子数组的个数
		temp, cnt := 0, 1
		for _, n := range nums {
			temp += n
			if temp > mid { // 当前子数组的和超过了 允许的最大值, 要放到下一个子数组中
				temp = n
				cnt++
			}
		}

		// 如果子数组的数量比 给定的数组的数量要大,
		// 说明 数组和的最小值要大于 mid
		// 否则 数组和最大值 应该要比 mid 小
		if cnt > m {
			max = mid + 1
		} else {
			sum = mid
		}
	}
	return max
}

func splitArray2(nums []int, k int) int {
	var maxElement int
	var total int
	for _, num := range nums {
		maxElement = max(num, maxElement)
		total += num
	}

	if len(nums) <= k {
		// fast path: 直接返回最大值. 每个组只有一个数字, 那么最大值就是数组整体的最大值
		return maxElement
	}

	// 左边界: 每个数组只有一个数字
	var left = maxElement
	// 右边界: 所有的数字在一个数组中
	var right = total

	for left < right {
		maxSubElement := left + (right-left)/2
		var curGroupSum int
		var groupCnt int
		groupCnt++
		for _, num := range nums {
			curGroupSum += num
			if curGroupSum > maxSubElement {
				curGroupSum = num
				groupCnt++
			}
		}
		// 判断子数组的个数
		if groupCnt > k {
			// 最大值偏小, 增大左边界
			left = maxSubElement + 1
		} else {
			// 最大值偏大, 缩减右边界
			right = maxSubElement
		}
	}
	return left
}

func main() {

}
