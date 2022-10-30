package main

func partitionDisjoint(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lm, m := nums[0], nums[0]
	var index int
	// 就是贪心, 找到最大值就去更新对应的索引位置
	// 不过, 还需要维护一个全局的最大值, 才可以保证更新索引时不会覆盖掉
	for i := 1; i < len(nums); i++ {
		m = max(m, nums[i])
		// 很巧妙的做法
		if lm > nums[i] {
			lm = m
			index = i
		}
	}
	return index + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
