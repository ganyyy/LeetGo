package main

import "sort"

func numMovesStonesII(s []int) []int {
	sort.Ints(s)
	n := len(s)
	// 最大的移动次数, 等同于 所有的空位数
	// 简单而言, 就是保证空位最多, 然后依次走一边
	/*
		    0     3   5   7       11
			X _ _ X _ X _ X _ _ _ X
			|-------------|
				  |---------------|
	*/
	e1 := s[n-2] - s[0] - n + 2
	e2 := s[n-1] - s[1] - n + 2 // 计算空位
	maxMove := max(e1, e2)
	if e1 == 0 || e2 == 0 { // 特殊情况：没有空位
		return []int{min(2, maxMove), maxMove}
	}
	// 最小的移动次数
	// 长度为n的窗口内, 最多的石子个数
	// 只要窗口内石子最多, 说明移动的次数越少

	/*
		    0     3   5   7       11
			X _ _ X _ X _ X _ _ _ X
		         |_ _ _ _ _|
	*/
	maxCnt, left := 0, 0
	for right, x := range s { // 滑动窗口：枚举右端点
		// x是右边界, x-n就是左边界
		for s[left] <= x-n { // 窗口大小大于 n
			left++
		}
		// right-left 就是 长度为 n 区间内的窗口的大小
		// 为啥是下标呢? 下标之间的差值就是 石子的个数, 然后 n-最大石子个数 = 最少的空位数
		maxCnt = max(maxCnt, right-left+1) // 维护窗口内的最大石子数
	}
	return []int{n - maxCnt, maxMove}
}
