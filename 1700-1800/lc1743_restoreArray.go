package main

import (
	"bytes"
)

const mx int = 1e5

func restoreArray(adjacentPairs [][]int) []int {

	g := [mx*2 + 1][2]int{} // 邻接的点. 每个点最多出现两次
	sz := [mx*2 + 1]byte{}  // 计数
	for _, p := range adjacentPairs {
		w := p[1] + mx
		v := p[0] + mx
		g[v][sz[v]] = w // 构建映射
		sz[v]++
		g[w][sz[w]] = v
		sz[w]++
	}

	n := len(adjacentPairs) + 1
	ans := make([]int, n)
	ans[0] = bytes.IndexByte(sz[:], 1) - mx // 找到第一个出现次数为1的点作为起点

	ans[1] = g[ans[0]+mx][0] - mx // 这里为啥取得是0呢? 因为ans[0]只有一个, 所以根据映射关系,
	// 这里也只有一个
	for i := 2; i < n; i++ {
		adj := g[ans[i-1]+mx]      // 基于前一个点的信息, 连起来
		if ans[i-2] == adj[0]-mx { // 这里需要看一下前边是不是出现了重复的数字.
			ans[i] = adj[1] - mx
		} else {
			ans[i] = adj[0] - mx
		}
	}
	return ans
}
