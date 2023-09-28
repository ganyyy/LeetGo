package main

func trailingZeroes(n int) int {
	var cnt int

	// 2 * 5 出一个0
	// 统计质因数为2的个数, 以及质因数为5的个数
	// 但是num(5) << num(2)
	// 所以相当于计算有多少个5
	for n != 0 {
		// 这种方式相当于分别计算5,25,125,625... 的数量
		cnt += n / 5
		n /= 5
	}
	return cnt
}
