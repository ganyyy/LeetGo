package main

import "sort"

func minimizeArrayValue(nums []int) int {
	// 最大最小值问题啊
	mx := 0
	for _, x := range nums {
		mx = max(mx, x)
	}
	return sort.Search(mx, func(maxNum int) bool {
		// 关键是怎么把这个限制给应用到数组上的
		// 这个maxNum相当于是设定了数组中的最大值
		// 比这个大的要减(nums[i]-maxNum)次并加到nums[i-1]上
		// 同理, 这个拥有传递性. 中断的前提是 在某个位置上, nums[i]+extra <= maxNum
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			// 这里相当于获取所有大于maxNum的值, 然后减去maxNum, 然后不停的累加到前面去.
			// 相当于比maxNum大的值都被减去了, 然后将这个extra分配给前面比maxNum小的值
			extra = max(nums[i]+extra-maxNum, 0)
		}
		// 最后所有的累加值都到了nums[0]上
		// 如果为true, 说明这个maxNum是可行的, 然后尝试更小的值
		// 如果为false, 说明这个maxNum是不可行的, 需要尝试更大的值
		return nums[0]+extra <= maxNum
	})
}
