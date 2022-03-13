package main

func countHighestScoreNodes(parents []int) (ans int) {
	n := len(parents)
	children := make([][]int, n)
	// 首先构建所有的边关系
	// 这个理论上支持多叉树
	for node := 1; node < n; node++ {
		p := parents[node]
		children[p] = append(children[p], node)
	}

	maxScore := 0
	// 先理解一下这个 dfs 是要干啥的

	// 以当前节点为根节点的子树的节点的个数
	var dfs func(int) int
	dfs = func(node int) int {
		// size的起始值为n-1, 表示该子树至少有一个节点
		score, size := 1, n-1
		for _, ch := range children[node] {
			// 子节点越多, size越小, 那么最终的结果就会越大
			sz := dfs(ch)
			// score是顺便计算的
			score *= sz
			size -= sz
		}
		if node > 0 {
			// 如果不是根节点的话, 还需要计算其父节点对应的得分
			score *= size
		}
		// 更新得分统计
		if score == maxScore {
			ans++
		} else if score > maxScore {
			maxScore = score
			ans = 1
		}
		// 返回删除该节点获取到的子树的分数
		return n - size
	}
	dfs(0)
	return
}
