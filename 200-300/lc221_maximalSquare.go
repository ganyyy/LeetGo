package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	ln := len(matrix)
	if 0 == ln {
		return 0
	}
	lm := len(matrix[0])
	if 0 == lm {
		return 0
	}
	var res int
	for i := 0; i < ln; i++ {
		for j := 0; j < lm; j++ {
			// 0表示无法扩散, 2表示已经扩散
			if '0' == matrix[i][j] || '2' == matrix[i][j] {
				continue
			}
			res = max(getMax(matrix, i, j), res)
		}
	}
	return res
}

// 返回最大正方形距离
func getMax(matrix [][]byte, row, col int) int {
	var res int
	ln := len(matrix)
	lm := len(matrix[0])
	// 三个方向, 依次扩散
	var check = func(lv int) bool {
		if lv+col >= lm || lv+row >= ln {
			return false
		}
		// 下边行
		for t := col; t <= col+lv; t++ {
			if '0' == matrix[row+lv][t] {
				return false
			}
		}
		// 右边列
		for t := row; t <= row+lv; t++ {
			if '0' == matrix[t][col+lv] {
				return false
			}
		}
		// 对角线
		if '0' == matrix[row+lv][col+lv] {
			return false
		}
		return true
	}

	for {
		if check(res) {
			res++
		} else {
			break
		}
	}
	return res * res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maximalSquareDP(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1)
	if len(matrix) == 0 {
		return 0
	}
	for i := 0; i <= len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	var max int
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				// 取左，上，左上三个值中的最小值 + 1, 如果任意一点为0说明 不能构成 正方形
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max
}

func min(a, b int) int {
	if b > a {
		return a
	}
	return b
}

func main() {
	/**

	 */
	fmt.Println(maximalSquare([][]byte{{'0', '0', '0', '1'}, {'1', '1', '0', '1'}, {'1', '1', '1', '1'}, {'0', '1', '1', '1'}, {'0', '1', '1', '1'}}))
}
