package main

import "fmt"

func movingCount(m int, n int, k int) int {
	check := make([][]bool, m)
	for i := 0; i < m; i++ {
		check[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, check)
}

func dfs(i, j, m, n, k int, check [][]bool) int {
	// 如果越界或者已经走过或者超过上限, 就返回
	if i >= m || j >= n || sum(i)+sum(j) > k || check[i][j] {
		return 0
	}
	// 当前节点标记为可以走
	check[i][j] = true
	// 总数 = 1 + 向右走 + 向下走
	return 1 + dfs(i+1, j, m, n, k, check) + dfs(i, j+1, m, n, k, check)
}

// 求和
func sum(a int) int {
	var res int
	for a != 0 {
		res += a % 10
		a /= 10
	}
	return res
}

func main() {
	fmt.Println(sum(35))
}
