package main

func findDiagonalOrder(mat [][]int) []int {
	var row, col = len(mat), len(mat[0])

	var ret = make([]int, 0, row*col)

	var needReverse = false
	var reverse = func(data []int) {
		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
	}

	// 行
	for i := 0; i < col; i++ {
		var start = len(ret)

		for j, k := 0, i; k >= 0 && j < row; j, k = j+1, k-1 {
			ret = append(ret, mat[j][k])
		}

		needReverse = !needReverse
		if needReverse {
			reverse(ret[start:])
		}
	}
	// 列
	for i := 1; i < row; i++ {
		var start = len(ret)

		for j, k := i, col-1; j < row && k >= 0; j, k = j+1, k-1 {
			ret = append(ret, mat[j][k])
		}

		needReverse = !needReverse
		if needReverse {
			reverse(ret[start:])
		}
	}

	return ret
}

func findDiagonalOrderDirect(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n)
	// 总共的行数
	for i := 0; i < m+n-1; i++ {
		if i%2 == 1 {
			// 奇数行, 从上往下,
			// 行++, 列--
			x := max(i-n+1, 0) // 行下限0
			y := min(i, n-1)   // 列上限 n-1
			for x < m && y >= 0 {
				ans = append(ans, mat[x][y])
				x++
				y--
			}
		} else {
			// 偶数行, 从下往上
			// 行--, 列++
			x := min(i, m-1)   // 行上限 m-1
			y := max(i-m+1, 0) // 列下限 0
			for x >= 0 && y < n {
				ans = append(ans, mat[x][y])
				x--
				y++
			}
		}
	}
	return ans
}
