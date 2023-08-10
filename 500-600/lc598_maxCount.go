package main

func maxCount(m int, n int, ops [][]int) int {
	// 简单而言, 就是求交集罢了..

	var x, y = m, n

	for _, op := range ops {
		x, y = min(op[0], x), min(op[1], y)
	}

	return x * y
}
