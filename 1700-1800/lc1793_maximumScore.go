package main

func maximumScore(nums []int, k int) int {
	n := len(nums)
	left, right := k-1, k+1
	ans := 0
	// minNum: 枚举的数组中的最小值
	for minNum := nums[k]; ; minNum-- {
		// 注意: minNum在迭代过程中可能不存在于数组中,
		// 但是不影响最终的结果, 因为满足最终结果的minNum一定存在于数组中
		// 因为只有当minNum存在于数组中时, 才会有可能更新left和right!
		for left >= 0 && nums[left] >= minNum {
			left--
		}
		for right < n && nums[right] >= minNum {
			right++
		}
		ans = max(ans, (right-left-1)*minNum)
		if left == -1 && right == n {
			break
		}
	}
	return ans
}
