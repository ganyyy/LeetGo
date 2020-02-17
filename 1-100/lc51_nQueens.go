package main

const (
	Empty = '.'
	Queen = 'Q'
)

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	chess := make([][]byte, n)
	for i := 0; i < n; i++ {
		chess[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chess[i][j] = Empty
		}
	}
	dfs(chess, 0, &res)
	return res
}

// row 之前的数据都已经正确的放置了皇后
func dfs(chess [][]byte, row int, res *[][]string) {
	l := len(chess)
	if row == l {
		r := make([]string, l)
		for i := 0; i < l; i++ {
			r = append(r, string(chess[i]))
		}
		*res = append(*res, r)
		return
	}
	// 一行只能存在一个皇后
	for col := 0; col < l; col++ {
		if !isValid(chess, row, col) {
			continue
		}
		// 交换在换回去
		chess[row][col] = Queen
		dfs(chess, row+1, res)
		chess[row][col] = Empty
	}
}

func isValid(chess [][]byte, row, col int) bool {
	// 检查列
	// 检查左上
	// 检查右上

	return true
}

func main() {

}
