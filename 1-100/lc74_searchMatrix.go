package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	// 二分走起
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	total := col * row
	for left, right := 0, total-1; left <= right; {
		mid := (right + left) / 2
		r, c := mid/col, mid%col
		midv := matrix[r][c]
		if midv == target {
			return true
		} else {
			if midv > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	// 二分呗.. 这还不好搞..
	var n = len(matrix)
	if n == 0 {
		return false
	}
	var m = len(matrix[0])
	if m == 0 {
		return false
	}
	// 总数
	n *= m
	var getVal = func(idx int) int {
		return matrix[idx/m][idx%m]
	}
	var l, r = 0, n
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if v := getVal(mid); v == target {
			return true
		} else if v > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	v := [][]int{
		//{1,   3,  5,  7},
		//{10, 11, 16, 20},
		//{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(v, 13))
}
