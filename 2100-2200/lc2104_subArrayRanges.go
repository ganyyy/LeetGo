package main

func subArrayRanges(nums []int) int64 {
	// 先来个o(n^2)

	var ret int
	// 想想怎么优化吧...
	for i := 0; i < len(nums); i++ {
		var max, min = nums[i], nums[i]
		for j := i + 1; j < len(nums); j++ {
			if max < nums[j] {
				max = nums[j]
			}
			if min > nums[j] {
				min = nums[j]
			}
			ret += max - min
		}

	}
	return int64(ret)
}

func subArrayRangesStack(nums []int) int64 {
	n := len(nums)

	// 核心就是4个栈, 分别表示
	// 左边最大值对应的索引, 左边最小值对应的索引
	// 右边最大值对应的索引, 右边最小值对应的索引
	// 说得好, 我写不出来!

	minLeft := make([]int, n)
	maxLeft := make([]int, n)
	var minStk, maxStk []int
	for i, num := range nums {
		// minStack: 左边最小值
		for len(minStk) > 0 && nums[minStk[len(minStk)-1]] > num {
			minStk = minStk[:len(minStk)-1]
		}
		if len(minStk) > 0 {
			minLeft[i] = minStk[len(minStk)-1]
		} else {
			minLeft[i] = -1
		}
		minStk = append(minStk, i)

		// 如果 nums[maxStk[len(maxStk)-1]] == num, 那么根据定义，
		// nums[maxStk[len(maxStk)-1]] 逻辑上小于 num，因为 maxStk[len(maxStk)-1] < i
		for len(maxStk) > 0 && nums[maxStk[len(maxStk)-1]] <= num {
			maxStk = maxStk[:len(maxStk)-1]
		}
		if len(maxStk) > 0 {
			maxLeft[i] = maxStk[len(maxStk)-1]
		} else {
			maxLeft[i] = -1
		}
		maxStk = append(maxStk, i)
	}

	minRight := make([]int, n)
	maxRight := make([]int, n)
	minStk = minStk[:0]
	maxStk = maxStk[:0]
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		// 如果 nums[minStk[len(minStk)-1]] == num, 那么根据定义，
		// nums[minStk[len(minStk)-1]] 逻辑上大于 num，因为 minStk[len(minStk)-1] > i
		for len(minStk) > 0 && nums[minStk[len(minStk)-1]] >= num {
			minStk = minStk[:len(minStk)-1]
		}
		if len(minStk) > 0 {
			minRight[i] = minStk[len(minStk)-1]
		} else {
			minRight[i] = n
		}
		minStk = append(minStk, i)

		for len(maxStk) > 0 && nums[maxStk[len(maxStk)-1]] < num {
			maxStk = maxStk[:len(maxStk)-1]
		}
		if len(maxStk) > 0 {
			maxRight[i] = maxStk[len(maxStk)-1]
		} else {
			maxRight[i] = n
		}
		maxStk = append(maxStk, i)
	}

	var sumMax, sumMin int64
	for i, num := range nums {
		sumMax += int64(maxRight[i]-i) * int64(i-maxLeft[i]) * int64(num)
		sumMin += int64(minRight[i]-i) * int64(i-minLeft[i]) * int64(num)
	}
	return sumMax - sumMin
}
