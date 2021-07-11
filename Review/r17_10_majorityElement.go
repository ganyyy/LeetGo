package main

func majorityElement(nums []int) int {
	// 摩尔计数法

	// 1. 当当前计数为0时, 候选者就是当前数
	// 2. 当当前数和候选者相同时, 计数+1, 否则-1

	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	// 这里需要进行第二波验证. 因为可能会出现恰好数量不足的情况
	// 如果不进行验证, 可能会出现 [1,1,2,2,3,4]算出来的结果为2
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	if count*2 > len(nums) {
		return candidate
	}
	return -1
}
