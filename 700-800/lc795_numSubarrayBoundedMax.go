package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	l := -1
	var tmp, ret int

	for r, v := range nums {
		if v > right {
			// 开启一个新的区间
			l = r
		}
		// 这一步就很精髓
		// 可以这么理解
		// [2,1,4,3] 2, 3
		// [2]满足, tmp = 1
		// [1]不满足, 但是 [2,1]满足, 所以此时应该再+1
		//
		if v >= left {
			// 新区间, 这个tmp是0
			tmp = r - l
		}
		ret += tmp
	}

	return ret
}

func numSubarrayBoundedMax2(nums []int, left int, right int) int {
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
		// [1]不满足, 但是 [2,1]满足, 此时提供的额外共享和[2]是等价的, 也就是再+1
		// 假设此时来了一个[3], 那么整体提供的贡献就是
		//  [2,1,3], [1,3], [3] = 3, 等同于数组的长度
		//
		if v >= left {
			// 注意: 上边 v > right 的情况会继续执行,
			// l = r, 所以这里的相当于将tmp置为0!
			tmp = r - l
		}
		ret += tmp
	}

	return ret
}
