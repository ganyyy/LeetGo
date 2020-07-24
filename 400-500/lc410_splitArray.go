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

func main() {

}
