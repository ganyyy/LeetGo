package main

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minChanges(nums []int, k int) int {
	// 这个数学推论直接看官方题解..
	// 数字以k为周期, 分为k组, 且组内所有元素的值必须相等
	// 目的是在k组数字中, 找到改变最小的数字的个数
	// 等同于最多能有多少个数字不发生变化

	// 每个区间数字的计数
	fs := make([]map[int]int, k)
	for i := 0; i < k; i++ {
		fs[i] = map[int]int{}
	}
	for i, n := range nums {
		fs[i%k][n]++
	}

	// mas 是每个区间中, 出现次数最多的数字的个数
	// mass 是mas后缀和(mass[0] = sum(mas[:]), mass[k-1]=mas[k-1])
	mas, mass := make([]int, k), make([]int, k)

	// 统计每个区间出现最多的数字的个数
	// 在该组内, 这个数字就是不需要变化的
	for i, f := range fs {
		for _, c := range f {
			mas[i] = max(mas[i], c)
		}
	}

	// 计算后缀和, 每个位置上的值表示后续不可变的数字的个数
	mass[k-1] = mas[k-1]
	for i := k - 2; i >= 0; i-- {
		mass[i] = mass[i+1] + mas[i]
	}

	// 选出各个区间中, 出现次数最多的数字个数中的最小值
	m := mas[0]
	for i := 1; i < k; i++ {
		m = min(m, mas[i])
	}

	// 如果要求某组数字全变, 肯定选择出现次数最少的那一组
	// o表示当前剩余的不需要变的数字的个数
	o := mass[0] - m

	var cal func(int, int, int)
	cal = func(i, s, c int) {
		// i: 当前是第几组
		// s: 上一组异或的结果?
		// c: 保留的数字的个数
		if i == k {
			// 到第k组时, 需要保证异或的结果为0.
			if s == 0 {
				o = max(o, c)
			}
		} else if c+mass[i] > o {
			// 这也算是一种枝减. 如果保留的数字的数量小于已经计算过的最大值
			// 就不需要处理该组了.
			// 计算的前提是 改组数字可能会出现大于最大值的情况
			for j, f := range fs[i] {
				// 当前组每个 数字和其出现的次数
				cal(i+1, s^j, c+f)
			}
		}
	}

	cal(0, 0, 0)
	// 最终结果就是数组的长度-不可变数字的个数
	return len(nums) - o
}
