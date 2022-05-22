package main

func repeatedNTimes(nums []int) int {
	// 可以理解为, 超过一半的数中, 每三个数中一定会存在相同的值
	for gap := 1; gap <= 3; gap++ {
		for i, num := range nums[:len(nums)-gap] {
			if num == nums[i+gap] {
				return num
			}
		}
	}
	return -1 // 不可能的情况
}
