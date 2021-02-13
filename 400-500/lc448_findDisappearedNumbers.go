package main

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		// 将正数作为下标, 对应位置的值置为负数
		nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
	}
	// 大于0的就是没出现过的
	var res []int
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
