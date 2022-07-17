package main

func arrayNesting(nums []int) int {
	// 枝减?
	var ret int
	for _, v := range nums {
		// 每个位置所形成的链表肯定是唯一的
		// 每个位置最多就查一次
		if v == -1 {
			continue
		}
		if ret > len(nums)/2 {
			return ret
		}
		curMax := 1
		cur := nums[v]
		nums[v] = -1
		for v != cur {
			curMax++
			// 所有走过的位置全部置成-1
			// 这个相当于最后一步设置的头节点
			cur, nums[cur] = nums[cur], -1
		}
		if curMax > ret {
			ret = curMax
		}
	}
	return ret
}
