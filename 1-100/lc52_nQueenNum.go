package main

import "fmt"

type Bit52 uint64

func (b Bit52) Get(idx int) bool    { return int(b)&(1<<idx) != 0 }
func (b Bit52) Set(idx int) Bit52   { return Bit52(int(b) | (1 << idx)) }
func (b Bit52) Unset(idx int) Bit52 { return Bit52(int(b) &^ (1 << idx)) }

func totalNQueens(n int) int {
	var dfs func(i int)

	// n的上限是9, 所以用uint64 足够了

	// 同一列, 最多n个值
	var cols Bit52
	// 主对角线(左上到右下) 主对角线的 row-col = const, 最大为(n-1, 0), 最小为(0, n-1), 需要通过+n 来保证都是>0的
	// 可以这么理解, 当到下一行时, row/col 都+1. 所以二者相减的值是恒定的
	var ml Bit52 // 范围: [0, 2n-1]
	// 次对角线(右上到左下) 次对角线的 row+col = const, 最大为(n-1, n-1)
	// 可以这么理解, 当到下一行时, row+1, col-1. 所以二者相加的值是恒定的
	var sl Bit52 // 范围: [0, 2n-2]

	var res int

	dfs = func(i int) {
		if i == n {
			// 增加一个解
			res++
		} else {
			for j := 0; j < n; j++ {
				// 在这一行的每个格子上尝试添加一枚皇后
				c, m, s := j, i-j+n, i+j
				if !cols.Get(c) && !ml.Get(m) && !sl.Get(s) {
					// 更新列, 主次对角线, 棋盘对应位置的值
					cols, ml, sl = cols.Set(c), ml.Set(m), sl.Set(s)
					// 迭代下一行
					dfs(i + 1)
					// 复原
					cols, ml, sl = cols.Unset(c), ml.Unset(m), sl.Unset(s)
				}

			}
		}
	}
	dfs(0)
	return res
}

func main() {
	fmt.Println(totalNQueens(4))
}
