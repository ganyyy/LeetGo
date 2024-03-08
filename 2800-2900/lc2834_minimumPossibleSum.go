package main

func minimumPossibleSum(n, target int) int {
	// 对于小于 m = ⌊target/2⌋ 的数字来说, 可以直接添加
	// 如果 m < n, 那么添加的就是target, target+1.... 一系列的数字
	// 整体来看, 就是两个等差数列(公差为1)的求和
	// ∑(1,m) m + ∑(target, target+n-m-1) n-m
	m := min(target/2, n)
	return (m*(m+1) + (target*2+n-m-1)*(n-m)) / 2 % 1_000_000_007
}
