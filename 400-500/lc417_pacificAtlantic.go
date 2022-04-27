package main

var dir = [4][2]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func pacificAtlantic(heights [][]int) [][]int {
	// 先确定一下, 整体的最小值

	var m, n = len(heights), len(heights[0])
	var flag = make([][]uint8, m)
	for i := range flag {
		flag[i] = make([]uint8, n)
	}
	// 两边逆流而上, 找对边

	var checkA = func(i, j int) bool {
		return i >= m || j >= n
	}
	var checkB = func(i, j int) bool {
		return i < 0 || j < 0
	}

	var dfs func(i, j int)
	var shift = 0

	dfs = func(i, j int) {
		flag[i][j] |= 1 << shift
		for _, d := range dir {
			var na, nb = i + d[0], j + d[1]
			if checkA(na, nb) || checkB(na, nb) || flag[na][nb]&(1<<shift) != 0 {
				continue
			}
			if heights[na][nb] < heights[i][j] {
				continue
			}
			dfs(na, nb)
		}
	}

	// 首尾两行
	for i := 0; i < n; i++ {
		shift = 0
		dfs(0, i)
		shift = 1
		dfs(m-1, i)
	}

	// 左右两端
	for i := 0; i < m; i++ {
		shift = 0
		dfs(i, 0)
		shift = 1
		dfs(i, n-1)
	}

	var res [][]int
	for i := 0; i < m; i++ {
		// fmt.Println(flag[i])
		for j := 0; j < n; j++ {
			if flag[i][j] != 3 {
				continue
			}
			res = append(res, []int{int(i), int(j)})
		}
	}

	return res
}
