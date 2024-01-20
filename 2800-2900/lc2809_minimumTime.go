package main

import "slices"

func minimumTime(nums1, nums2 []int, x int) int {
	s1, s2, n := 0, 0, len(nums1)
	id := make([]int, n)
	for i := range id {
		id[i] = i
		s1 += nums1[i]
		s2 += nums2[i]
	}
	// 对下标数组排序，避免破坏 nums1 和 nums2 的对应关系
	slices.SortFunc(id, func(i, j int) int { return nums2[i] - nums2[j] })

	// 0-n, 可以理解为最少操作0次, 最多操作n次. 即每个岛屿操作一次(?)
	// 优先重置较小的增量对应的数, 这样可以保证整体的增量更小
	// 每个位置上, 至多重置1次
	// f[i]保存的是到达第i秒时, 整体减小的最大值
	f := make([]int, n+1)
	for i, p := range id {
		// 0/1背包. 选或者不选的区别
		// 因为是二维的压缩, 这里的后态基于前态, 所以倒序迭代
		a, b := nums1[p], nums2[p]
		for j := i + 1; j > 0; j-- {
			// 如果第j次不重置对应的数, 则相当于没有发生减少
			// 如果第j次  重置对应的数, 则相当于在f[j-1]的基础上, 当前减少额外增加了 nums1[i]+nums2[i]*j(相当于前j s的增量被清理了)
			f[j] = max(f[j], f[j-1]+a+b*j)
		}
	}

	for t, v := range f {
		// t是操作的秒数
		// v是最大的减少量
		// s1+s2*t表示的是到达指定秒数时, 没发生任何减少时的原始值
		// 那么减去v之后, 相当于最少的剩余值. 找到首个满足条件的数, 返回即可(因为它是最近的!)
		if s1+s2*t-v <= x {
			return t
		}
	}
	// 找不到返回-1
	return -1
}
