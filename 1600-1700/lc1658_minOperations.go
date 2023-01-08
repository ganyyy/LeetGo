package main

func minOperations(nums []int, x int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < x {
		return -1
	}

	// [left|...|right]
	// 初始状态下, left = -1, right = 0
	// lSum = 0, rSum = sum

	right := 0
	lSum := 0
	rSum := sum
	ans := n + 1

	for left := -1; left < n; left++ {
		// 左指针++, lSum++
		if left != -1 {
			lSum += nums[left]
		}
		// 右指针++, rSum--
		for right < n && lSum+rSum > x {
			rSum -= nums[right]
			right++
		}
		if lSum+rSum == x {
			ans = min(ans, (left+1)+(n-right))
		}
	}
	if ans > n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
