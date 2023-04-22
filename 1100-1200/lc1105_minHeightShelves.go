//go:build ignore

package main

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	f := make([]int, n+1) // f[0]=0，翻译自 dfs(-1)=0
	for i := range books {
		f[i+1] = math.MaxInt
		maxH, leftW := 0, shelfWidth
		for j := i; j >= 0; j-- {
			leftW -= books[j][0]
			if leftW < 0 { // 空间不足，无法放书
				break
			}
			// 没啥好办法.. 就是先尝试放在同一行, 然后统计高度
			// 为啥要从后向前呢?
			maxH = max(maxH, books[j][1]) // 从 j 到 i 的最大高度
			f[i+1] = min(f[i+1], f[j]+maxH)
		}
	}
	return f[n] // 翻译自 dfs(n-1)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
