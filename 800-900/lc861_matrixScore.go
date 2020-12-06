package main

func matrixScore(a [][]int) int {
	// 权重最大的是 最左边的一列, 将最左边的一列通过行移动全变成1
	// 然后通过列变化让剩余的列里1的数量大于0的数量
	m, n := len(a), len(a[0])
	// 假设第一列都是1, 这是计算的基础
	ans := 1 << (n - 1) * m

	// 处理的时候不需要真的去转置矩阵, 只需要直接计算结果即可

	// 遍历剩下的列
	for j := 1; j < n; j++ {
		// 每一列中和行头相同的数的个数, 可以理解为 首列 行移动后 当前列中1的个数
		ones := 0
		for i := 0; i < m; i++ {
			if a[i][j] == a[i][0] {
				ones++
			}
		}
		// 如果1的个数少, 就将这一列进行反转
		if ones < m-ones {
			ones = m - ones
		}
		// 结果加上这一列1的数量
		ans += 1 << (n - 1 - j) * ones
	}
	return ans
}
