package main

/**
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

这里采用的方法是以空间换时间.
*/

func longestConsecutive(nums []int) int {
	// 每个端点值对应的连续区间长度
	m := make(map[int]int, len(nums))
	var res int
	for _, v := range nums {
		// 如果已存在, 直接略过
		if _, ok := m[v]; ok {
			continue
		}
		// 取左右的长度
		l, r := m[v-1], m[v+1]
		// 更新最大长度
		var now = l + r + 1
		if now > res {
			res = now
		}
		// 更新该连续区间左右端点的值
		m[v] = now
		m[v-l] = now
		m[v+r] = now

	}

	return res
}

func main() {

}
