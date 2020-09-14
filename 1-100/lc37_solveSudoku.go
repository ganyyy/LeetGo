package main

import "fmt"

type Bitset int

func (b *Bitset) Set(pos int) {
	*b |= 1 << pos
}

func (b *Bitset) Check(pos int) bool {
	return *b&(1<<pos) != 0
}

func (b *Bitset) Clear(pos int) {
	*b &^= 1 << pos
}

func solveSudoku(board [][]byte) bool {
	// 行, 列, 块, 指定位置是否被选中
	var row, col, block [9]Bitset

	// 初始化指定的位置
	for i, r := range board {
		for j, v := range r {
			if v != '.' {
				pos := int(v - '1')
				row[i].Set(pos)
				col[j].Set(pos)
				block[i/3*3+j/3].Set(pos)
			}
		}
	}

	var dfs func(i, j int) bool

	dfs = func(i, j int) bool {
		// 寻找一个空位置
		for board[i][j] != '.' {
			if j >= 8 {
				i++
				j = 0
			} else {
				j++
			}
			// 没有空位了, 全部搞完了
			if i >= 9 {
				return true
			}
		}
		for pos := 0; pos < 9; pos++ {
			blockIdx := i/3*3 + j/3
			// 如果都没放进去过
			if !row[i].Check(pos) && !col[j].Check(pos) && !block[blockIdx].Check(pos) {
				row[i].Set(pos)
				col[j].Set(pos)
				block[blockIdx].Set(pos)
				board[i][j] = byte(pos + '1')
				if dfs(i, j) {
					return true
				} else {
					// 清空
					row[i].Clear(pos)
					col[j].Clear(pos)
					block[blockIdx].Clear(pos)
					board[i][j] = '.'
				}
			}
		}
		return false
	}
	return dfs(0, 0)
}

func main() {
	origin := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(origin)
	for _, v := range origin {
		fmt.Println(string(v))
	}
}
