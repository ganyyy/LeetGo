package main

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	// 计算前缀和
	var n = len(matrix)
	if n == 0 {
		return NumMatrix{}
	}
	var m = len(matrix[0])
	var tmp = make([][]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			tmp[i][j] = matrix[i][j]
			if i > 0 {
				tmp[i][j] += tmp[i-1][j]
			}
			if j > 0 {
				tmp[i][j] += tmp[i][j-1]
			}
			if i > 0 && j > 0 {
				tmp[i][j] -= tmp[i-1][j-1]
			}
		}
	}
	return NumMatrix{
		sum: tmp,
	}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 整体看成四个部分的运算
	// 右下角 - 右上角 - 左下角 + 左上角
	var rb = m.sum[row2][col2]
	var rt, lb, lt int
	if row1-1 >= 0 {
		rt = m.sum[row1-1][col2]
	}
	if col1-1 >= 0 {
		lb = m.sum[row2][col1-1]
	}
	if row1-1 >= 0 && col1-1 >= 0 {
		lt = m.sum[row1-1][col1-1]
	}
	return rb - rt - lb + lt
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
