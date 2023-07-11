package main

func countBits(num int) []int {

	var res = make([]int, num+1)

	// 如何快速的计算bit个数?
	// 可以通过前缀和实现吗?

	// 如果是奇数, 那么就是前边的偶数的比特位+1
	// 如果是偶数, 那么就和 >> 1 得到的 奇数/偶数 的比特位相同

	// 初始值肯定是0
	res[0] = 0
	for i := 1; i <= num; i++ {
		if (i & 1) == 1 {
			// 奇数
			res[i] = res[i-1] + 1
		} else {
			// 偶数
			res[i] = res[i>>1]
		}

		// 或许可以再进一步的优化一下

		// res[i] = res[i>>1] + (i & 1)
	}

	return res
}

func countBits2(num int) []int {

	var res = make([]int, num+1)

	// 先来几个初始值
	res[0] = 0
	// 好一个dp, 学到了
	for i := 1; i <= num; i++ {
		res[i] = res[i>>1] + (i & 1)
	}

	return res
}
