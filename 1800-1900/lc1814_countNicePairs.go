package main

func countNicePairs(nums []int) int {
	var count = make(map[int]int)
	var ret int
	// nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
	// 等同于 nums[i] - rev(nums[i]) = nums[j] - rev(nums[j])
	// 所以统计 nums[i] - rev(nums[i]) 的数量即可
	for _, num := range nums {
		sub := num - rev(num)
		old := count[sub]
		ret = (ret + old) % (1e9 + 7)
		count[sub] = old + 1
	}
	return ret
}

func rev(num int) (ret int) {
	for num != 0 {
		ret = ret*10 + num%10
		num /= 10
	}
	return
}
