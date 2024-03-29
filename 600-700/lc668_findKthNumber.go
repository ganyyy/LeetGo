package main

import "sort"

func findKthNumber(m, n, k int) int {
	// 第K小说明小于等于x的元素有K-1个
	// M行, N列, 每行N个元素

	// 值二分法
	return sort.Search(m*n, func(x int) bool {
		// 查找比X小的数的个数
		// 完全小于x的行的数量
		count := x / n * n
		// 从部分小于x的行开始, 计算总共小于x的值的数量
		for i := x/n + 1; i <= m; i++ {
			count += x / i
		}
		// 总的小于X的个数,
		// 如果数量大于等于K, 说明第K小的元素在左边
		// 如果数量小于K, 说明第K小的元素在右边
		// 这里为啥取 = 号呢? 是为了求取左边界
		// 3 x 3的乘法表, 不存在5,7,8, 此时
		//	5和4是等价的
		//	6,7,8是等价的
		// 当取得5时, 就需要继续向左靠直到左边界
		// 1, 2, 3
		// 2, 4, 6
		// 3, 6, 9
		// 1,2,2,3,3,4,6,6,9
		// N(<=4) == N(<=5) = 6
		// N(<=6) == N(<=7) == N(<=8) = 8
		// 反过来说, 求第6小的数, 得到的应该是4
		// 求第7小的数/第8小的数, 得到的应该是6
		return count >= k
	})
}
