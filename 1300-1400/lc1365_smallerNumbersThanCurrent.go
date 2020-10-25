package main

func smallerNumbersThanCurrent(nums []int) []int {
	// 无脑走一遍当然可以, 此时时间复杂度是 O(n^2)
	// 肯定存在 O(n) 的方案.

	// 桶的思想解决问题

	// [0,100]
	var bucket [101]int

	// 统计每个数字的数量
	for _, v := range nums {
		bucket[v]++
	}

	// 计算小于当前值的数的个数
	for i := 1; i < len(bucket); i++ {
		bucket[i] += bucket[i-1]
	}

	// 计算结果
	for i, v := range nums {
		if v != 0 {
			nums[i] = bucket[v-1]
		}
	}

	return nums
}
