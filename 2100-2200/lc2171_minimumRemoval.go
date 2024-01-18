package main

import "slices"

func minimumRemoval(beans []int) int64 {
	slices.Sort(beans)
	sum, mx, length := 0, 0, len(beans)
	for i, v := range beans {
		sum += v
		// 整体排序后, 意味着只需要扣除比当前多的袋子中, 额外多出的数量
		// 前边的就当清零了.
		// 那么, 相当于计算保留哪个数量时, 可以带来最多的豆子
		mx = max(mx, (length-i)*v)
	}
	// 使用总数减去这个保留数量之后(最大值), 就是应该扣除的数量
	return int64(sum - mx)
}
