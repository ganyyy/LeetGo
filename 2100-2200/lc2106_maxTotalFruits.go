package main

import "sort"

func maxTotalFruits(fruits [][]int, startPos, k int) int {
	n := len(fruits)
	// 向左最远能到 fruits[left][0]
	left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos-k })
	right, s := left, 0
	for ; right < n && fruits[right][0] <= startPos; right++ {
		s += fruits[right][1] // 从 fruits[left][0] 到 startPos 的水果数
	}
	ans := s
	for ; right < n && fruits[right][0] <= startPos+k; right++ {
		s += fruits[right][1] // 枚举最右位置为 fruits[right][0]
		// 利益最大化, 肯定就得单向
		// 先向右, 再向左: (right - startPos) + (right - left) <= k
		// 先向左, 再向右: (startPos - left) + (right - left) <= k
		for fruits[right][0]*2-fruits[left][0]-startPos > k &&
			fruits[right][0]-fruits[left][0]*2+startPos > k {
			s -= fruits[left][1] // fruits[left][0] 无法到达
			left++
		}
		ans = max(ans, s) // 更新答案最大值
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
