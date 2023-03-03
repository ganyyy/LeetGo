package main

func countTriplets(nums []int) int {
	var cnt [1 << 16]int
	for i := range nums {
		for j := range nums {
			cnt[nums[i]&nums[j]]++
		}
	}
	res := 0
	for _, v := range nums {
		// 要找到所有 v & ? == 0的数
		// v (0b1010_0101)
		// ^v(0b0101_1010)
		// 那么 ? 一定是 ^v 的子集!
		x := v ^ 0xffff
		// 计算子集的方法:
		// x  :  0b0101_1010
		// sub1: 0b0101_1000
		// sub2: 0b0101_0000
		// sub3: 0b0100_0000
		// sub4: 0b0000_0000
		for sub := x; sub > 0; sub = (sub - 1) & x {
			res += cnt[sub]
		}
		res += cnt[0]
	}
	return res
}
