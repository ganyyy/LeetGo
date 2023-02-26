package main

func movesToMakeZigzag(nums []int) int {
	// 模拟, 然后分类讨论
	// 直接计算就行欸...
	help := func(pos int) int {
		res := 0
		for i := pos; i < len(nums); i += 2 {
			a := 0
			if i-1 >= 0 {
				a = max(a, nums[i]-nums[i-1]+1)
			}
			if i+1 < len(nums) {
				a = max(a, nums[i]-nums[i+1]+1)
			}
			// 将 i 变成 i-1, i+1 中的最大值, 所需要做的操作(?)
			// 因为只能减小 i-1, i+1 对应位置的值
			// 为啥不操作偶数位呢? 因为一旦调整了偶数位, 那么对奇数位的操作一定会更多(只能减少!)
			res += a
		}
		return res
	}

	return min(help(0), help(1))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
