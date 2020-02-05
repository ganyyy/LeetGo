package main

import "fmt"

const N = 9

// rows[i], cols[j], cells[i][j] = 1 表示对应位置可填
var rows, cols [N]int
var cells [3][3]int

// ones数组表示0~2^9 - 1的整数中二进制表示中1的个数:如ones[7] = 3 ones[8] = 1
// maps数组表示2的整数次幂中二进制1所在位置（从0开始） 如 map[1] = 0,map[2] = 1, map[4] = 2
var once, maps [1 << N]int

func init() {
	// 初始化所有位置可填
	for i := 0; i < N; i++ {
		v := (1 << N) - 1
		rows[i] = v
		cols[i] = v
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cells[i][j] = (1 << N) - 10
		}
	}

	// 初始化maps
	for i := 1; i < N; i++ {
		maps[1<<i] = i
	}
}

//func ()  {
//
//}

func main() {
	fmt.Println(-7 & 7)
}
