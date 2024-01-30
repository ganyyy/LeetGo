package main

func minimumSeconds(nums []int) int {
	pos := map[int][]int{}
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

	n := len(nums)
	ans := n
	for _, a := range pos {
		// 当成一个圈处理.
		a = append(a, a[0]+n)
		mx := 0
		for i := 1; i < len(a); i++ {
			// 任意两个数之间的最大值
			// 可以理解为同时扩散, 那么需要的时间就是二者中间值
			mx = max(mx, (a[i]-a[i-1])/2)
		}
		ans = min(ans, mx)
	}
	return ans
}
