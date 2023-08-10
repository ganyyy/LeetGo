package main

func maxAlternatingSum(nums []int) int64 {
	// 选 or 不选?

	// 作为奇数, 作为偶数, 作为开头(特殊的偶数)
	if len(nums) == 0 {
		return 0
	}

	var even int // 偶数子序列的最大和
	var odd int  // 奇数子序列的最大和
	var mx int
	even = nums[0]
	mx = even
	for _, v := range nums[1:] {
		ce := even - v      // 选取这个数字作为奇数位时
		co := max(v, odd+v) // 偶数位也可能是作为开头的情况, 同时也规避了 相加为0的情况(?)
		mx = max(mx, max(ce, co))
		// fmt.Println(even, odd, v, ce, co, mx)
		even = max(even, co) // 取最大值
		odd = max(odd, ce)   // 取最大值
	}
	return int64(mx)
}
