package main

func maximalNetworkRank(n int, roads [][]int) int {
	// TODO: 获取度数最大、次大的两个点的组合

	// connect 表示 i,j 之间是否存在链接
	connect := make([][]int, n)
	for i := range connect {
		connect[i] = make([]int, n)
	}
	// degree 表示 i 对应的边的数量
	degree := make([]int, n)
	for _, v := range roads {
		connect[v[0]][v[1]] = 1
		connect[v[1]][v[0]] = 1
		degree[v[0]]++
		degree[v[1]]++
	}

	maxRank := 0
	// 两点不能重复
	// 枚举所有可能的(i, j)组合
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 如果二者相交, 就需要减去一条共同边
			rank := degree[i] + degree[j] - connect[i][j]
			maxRank = max(maxRank, rank)
		}
	}
	return maxRank
}
