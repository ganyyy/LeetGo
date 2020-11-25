package main

func maximumGap(nums []int) int {
	// 桶排序

	// O(n)的时间复杂度的排序方案只有 桶排序, 对应的空间复杂度也为 O(n)

	var n = len(nums)
	if n < 2 {
		return 0
	}

	var max, min = math.MinInt32, math.MaxInt32

	// step1: 先找出最大值和最小值
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	if max == min {
		return 0
	}

	// step2: 构建排序桶
	// 设置n+1个桶的目的是为了保证 间距最大的两个数被分配到 不同的桶中.
	// 最大间距x >= (max-min)/(n-1) (恰好每个数之间相差为1的情况时相等)
	// 所以 x 严格大于 (max-min)/n
	// 使用 n+1个桶可以保证 每个桶的最大差值为(max-min)/n, 即同一个桶中的最大和最小值之间的差一定小于x

	// 每个桶的最大值, 最小值, 是否有值
	var mins, maxs, has = make([]int, n+1), make([]int, n+1), make([]bool, n+1)

	var idx int
	for _, v := range nums {
		// 计算当前值属于哪个桶
		idx = (v - min) * n / (max - min)
		// 维护每个桶最大值, 最小值, 是否存在
		if !has[idx] {
			mins[idx] = v
			maxs[idx] = v
			has[idx] = true
		} else {
			mins[idx] = getMin(mins[idx], v)
			maxs[idx] = getMax(maxs[idx], v)
		}
	}

	// step3: 寻找最大间距.
	// 因为每个桶之间的差值(maxs[i]-mins[i]) 一定是小于等于 (max-min)/n的
	// 所以 最大间距 一定在前一个桶的最大值和后一个桶的最小值的差值中.
	var res int
	max = maxs[0] // maxs[0]一定存在
	for i := 1; i <= n; i++ {
		if has[i] {
			if l := mins[i] - max; l > res {
				res = l
			}
			max = maxs[i]
		}
	}
	return res
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
