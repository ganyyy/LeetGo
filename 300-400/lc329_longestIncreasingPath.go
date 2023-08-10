package main

import "sort"

var Dir = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

// DP思路
func longestIncreasingPath(matrix [][]int) int {
	ln := len(matrix)
	if ln == 0 {
		return 0
	}
	lm := len(matrix[0])
	if lm == 0 {
		return 0
	}
	if lm == 1 && ln == 1 {
		return 1
	}

	// 构建dp数组 dp[i][j]表示 以这个位置为结尾 的最大递增数组的长度
	dp := make([][]int, ln)

	// 先把矩阵中的每一个数组按照大小进行排序
	numLst := make([][3]int, 0, ln*lm)
	for i := 0; i < ln; i++ {
		dp[i] = make([]int, lm)
		for j := 0; j < lm; j++ {
			// dp的每一个位置的默认值都是1, 表示长度为1的递增序列
			dp[i][j] = 1
			numLst = append(numLst, [3]int{matrix[i][j], i, j})
		}
	}
	sort.Slice(numLst, func(i, j int) bool {
		return numLst[i][0] < numLst[j][0]
	})

	var max int
	for _, v := range numLst {
		i, j := v[1], v[2]
		// 从头到尾遍历4个方向找最大值
		for _, dir := range Dir {
			// 坐标不越界
			if x, y := i+dir[0], j+dir[1]; x < ln && x >= 0 && y < lm && y >= 0 {
				// 如果周围存在比它小的值, 那么就意味着可以通过其他位置走到这里,
				// 尝试更新一下最长路径
				if matrix[i][j] > matrix[x][y] {
					if n := 1 + dp[x][y]; n > dp[i][j] {
						dp[i][j] = n
						if n > max {
							max = n
						}
					}
				}
			}
		}
	}
	return max
}

// DFS思路
func longestIncreasingPathDFS(matrix [][]int) int {
	ln := len(matrix)
	if ln == 0 {
		return 0
	}
	lm := len(matrix[0])
	if lm == 0 {
		return 0
	}
	if lm == 1 && ln == 1 {
		return 1
	}

	cache := make([][]int, ln)
	for i := 0; i < ln; i++ {
		cache[i] = make([]int, lm)
	}

	var helper func(i, j int) int

	helper = func(i, j int) int {
		// 四个方向依次查找, 如果有缓存就直接返回, 否则就递归到下一层
		if cache[i][j] != 0 {
			return cache[i][j]
		}
		m := 1
		if i >= 1 && matrix[i][j] < matrix[i-1][j] {
			m = max(m, 1+helper(i-1, j))
		}
		if j >= 1 && matrix[i][j] < matrix[i][j-1] {
			m = max(m, 1+helper(i, j-1))
		}
		if i < ln-1 && matrix[i][j] < matrix[i+1][j] {
			m = max(m, 1+helper(i+1, j))
		}
		if j < lm-1 && matrix[i][j] < matrix[i][j+1] {
			m = max(m, 1+helper(i, j+1))
		}
		cache[i][j] = m
		return m
	}

	m := 1
	// 从头到尾找
	for i := 0; i < ln; i++ {
		for j := 0; j < lm; j++ {
			m = max(m, helper(i, j))
		}
	}
	return m
}

func main() {

}
