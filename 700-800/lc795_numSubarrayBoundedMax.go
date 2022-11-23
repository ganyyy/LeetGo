package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	l := -1
	var tmp, ret int

	for r, v := range nums {
		if v > right {
			l = r
		}
		// 这一步就很精髓
		// 可以这么理解
		// [2,1,4,3] 2, 3
		// [2]满足, tmp = 1
		// [1]不满足, 但是 [2,1]满足, 所以此时应该再+1
		//
		if v >= left {
			tmp = r - l
		}
		ret += tmp
	}

	return ret
}
