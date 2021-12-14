package main

import "fmt"

func loudAndRich(richer [][]int, quiet []int) []int {
	// rich[i] 表示所有比i有钱的人
	var rich = make([][]int, len(quiet))

	for _, r := range richer {
		rich[r[1]] = append(rich[r[1]], r[0])
	}

	// 暴力方法: 依次遍历每一个比当前值有钱的玩家, 然后记录一下最小值(?)

	var ret = make([]int, len(quiet))
	for i := range ret {
		ret[i] = -1
	}

	var findMin func(i int) int

	findMin = func(i int) int {
		if ret[i] != -1 {
			return quiet[ret[i]]
		}
		var min = quiet[i]
		var mip = i
		for _, p := range rich[i] {
			if t := findMin(p); t < min {
				min = t
				mip = ret[p]
			}
		}
		ret[i] = mip
		return min
	}

	for i := range quiet {
		findMin(i)
	}

	return ret
}

func main() {
	fmt.Println(loudAndRich(
		[][]int{
			{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3},
		},
		[]int{
			3, 2, 5, 4, 6, 1, 7, 0,
		},
	))
}
