package main

func distributeCandies(n int, limit int) int {

	var c2 = func(n int) int {
		if n < 2 {
			return 0
		}
		return n * (n - 1) / 2
	}

	// 容斥原理
	// n个球, 两个挡板, 那么挡板可以放置的位置为 n+2
	//      这个咋理解呢? 简单来说: 把挡板也看成是一个特殊的球!
	//      然后从n+2个位置中选出两个位置作为挡板的位置, 就是所有分组的方案了!
	// 一个人超过limit的场景, 相当于预分配给他 limit+1个糖果, 然后从后续的糖果中再选两个位置: S1 = c2(n-(limit+1)+2)
	//      三个人都有可能, 所以是 3*S1,
	// 但是, 这里包含了有两个人超过limit的场景: S2 = c2(n-2*(limit+1)+2). 两两组合, 这个也得 * 3
	// 同时, 还包括了三个人都超过limit的情况, 此时 S3 = c2(n-3*(limit+1)+2). 只有一种组合

	return c2(n+2) - 3*c2(n-limit+1) + 3*c2(n-2*limit) - c2(n-3*limit-1)
}
