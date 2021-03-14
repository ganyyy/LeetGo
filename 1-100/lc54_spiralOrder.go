package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	ln := len(matrix)
	if ln == 0 {
		return []int{}
	}
	if ln == 1 {
		return matrix[0]
	}
	ln2 := len(matrix[0])
	res := make([]int, ln*ln2)
	// 边界
	rowStart, rowEnd, colStart, colEnd := 0, ln-1, 0, ln2-1
	// 增值, 从第一行开始
	row, rowAdd, col, colAdd := 0, 0, 0, 1

	for i := 0; i < len(res); i++ {
		res[i] = matrix[row][col]
		row += rowAdd
		col += colAdd
		// 行到头了
		if col > colEnd {
			// 行+1
			rowStart += 1
			// 换行, 上到下
			col = colEnd
			colAdd = 0
			rowAdd = 1
			row += 1
			continue
		}
		if col < colStart {
			// 行-1
			rowEnd -= 1
			// 换行, 下到上
			col = colStart
			colAdd = 0
			rowAdd = -1
			row -= 1
			continue
		}
		// 列到头了
		if row > rowEnd {
			// 列-1
			colEnd -= 1
			// 换列, 右到左
			row = rowEnd
			rowAdd = 0
			colAdd = -1
			col -= 1
			continue
		}
		if row < rowStart {
			// 列+1
			colStart += 1
			// 换列, 左到右
			row = rowStart
			rowAdd = 0
			colAdd = 1
			col += 1
			continue
		}
	}
	return res
}

func spiralOrder2(matrix [][]int) []int {
	var n = len(matrix)
	if n == 0 {
		return nil
	}
	var m = len(matrix[0])

	var res = make([]int, 0, m*n)

	// 可以搞一个递归, 输入边界值决定计算流程

	var helper func(sw, sh, w, h int)

	helper = func(sw, sh, ew, eh int) {
		// 如果都相等了, 说明还剩一个中间的数, 直接添加并返回即可
		fmt.Println(sw, sh, ew, eh)
		if sw > ew || sh > eh {
			return
		}
		// 特殊处理一下两种情况吧
		// 代码写的太难看了...
		if sw == ew {
			for i := sh; i <= eh; i++ {
				res = append(res, matrix[i][sw])
			}
			return
		}
		if sh == eh {
			for i := sw; i <= ew; i++ {
				res = append(res, matrix[sh][i])
			}
			return
		}

		// 横着
		for i := sw; i <= ew; i++ {
			res = append(res, matrix[sh][i])
		}
		// 竖着
		for i := sh + 1; i < eh; i++ {
			res = append(res, matrix[i][ew])
		}
		// 倒横着
		for i := ew; i >= sw; i-- {
			res = append(res, matrix[eh][i])
		}
		// 倒着竖
		for i := eh - 1; i > sh; i-- {
			res = append(res, matrix[i][sw])
		}
		helper(sw+1, sh+1, ew-1, eh-1)
	}

	helper(0, 0, m-1, n-1)

	return res
}

func spiralOrder3(matrix [][]int) []int {
	var n = len(matrix)
	if n == 0 {
		return nil
	}
	var m = len(matrix[0])

	var res = make([]int, 0, m*n)

	var left, right, up, down = 0, m - 1, 0, n - 1
	for {
		// 右
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++ // 右完了是下
		if up > down {
			break
		}
		// 下
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right-- // 下完了是左
		if right < left {
			break
		}
		// 左
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down-- // 左完了是上
		if down < up {
			break
		}
		// 上
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++ // 上完了是右
		if left > right {
			break
		}
	}

	return res
}

func main() {
	//param := [][]int {
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9,10,11,12},
	//	{13,14,15,16},
	//	{17,18,19,20},
	//}
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}}))
}
