//go:build ignore

package main

func maxSubarraySumCircular(nums []int) int {
	total, maxSum, minSum, currMax, currMin := nums[0], nums[0], nums[0], nums[0], nums[0]
	// 单向最大值(curMax)
	// 循环最大值 = total - 单项最小值
	for i := 1; i < len(nums); i++ {
		total += nums[i]
		currMax = max(currMax+nums[i], nums[i])
		maxSum = max(maxSum, currMax)
		currMin = min(currMin+nums[i], nums[i])
		minSum = min(minSum, currMin)
	}

	// 要么最大和子数组成环, 要么最小和子数组成环

	if total == minSum {
		// 相当于全部小于0的情况(!)
		// 此时的maxSum等同于数组中的最大值
		return maxSum
	} else {
		return max(maxSum, total-minSum)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxSubarraySumCircular2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	var total, maxSum, minSum, curMax, curMin int
	val := nums[0]
	start := val
	total, maxSum, minSum, curMax, curMin =
		start, start, start, start, start

	for _, val = range nums[1:] {
		total += val
		curMax = max(curMax+val, val)
		maxSum = max(curMax, maxSum) // 连续的最大子数组和
		curMin = min(curMin+val, val)
		minSum = min(curMin, minSum) // 连续的最小子数组和
	}

	if total == minSum {
		// 特殊情况: 整个数组的和就是最小子数组的和
		// 此时最大的子数组和一定在中间!
		return maxSum
	}
	// 将整体拆成两段:
	// 如果最大子数组和在中间, 那么就是sum
	// 如果最大子数组在两边, 那么就是total-minSum
	return max(maxSum, total-minSum)
}
