package main

import "sort"

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	// 老生常谈: 如果返回true, 代表着当前值满足条件, 那么就需要往左偏; 否则就是往右偏

	// 上界为什么是 (price[len(price)-1]-price[0])/(k-1) ?
	/*
		假设我们把数组中的 k 个不同的糖果按照价格从小到大排序，记为 a1, a2, …, ak，那么它们之间的最小绝对差就是 min(a2-a1, a3-a2, …, ak-ak-1)。

		如果我们把这个最小绝对差记为 x，那么我们可以得到以下不等式：

		a2 - a1 >= x a3 - a2 >= x … ak - ak-1 >= x

		把这些不等式相加，我们可以得到：

		ak - a1 >= kx - x

		也就是说，x <= (ak - a1) / (k - 1)。
		如果甜蜜度超过了 (ak - a1) / (k - 1)，那么我们就无法找到 k 个不同的糖果，使得它们之间的最小绝对差不小于甜蜜度。
	*/

	// 针对甜蜜度进行值域二分, 下界是1, 上界是 (price[len(price)-1]-price[0])/(k-1)
	return sort.Search((price[len(price)-1]-price[0])/(k-1), func(d int) bool {
		d++ // 二分最小的 f(d+1) < k，从而知道最大的 f(d) >= k
		cnt, pre := 1, price[0]
		// 双指针, 找到所有间隔大于 d 的糖果并计数
		for _, p := range price[1:] {
			if p >= pre+d {
				cnt++
				pre = p
			}
		}
		// 如果数量少了, 就意味着 d 太大了, 需要往左偏(d减小); 否则就是往右偏(d增大)
		return cnt < k
	})
}
