package main

func arrayNesting(nums []int) int {
	// 枝减?
	var ret int
	for _, v := range nums {
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
			cur, nums[cur] = nums[cur], -1
		}
		if curMax > ret {
			ret = curMax
		}
	}
	return ret
}
