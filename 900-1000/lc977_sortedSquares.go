package main

import "sort"

func sortedSquares2(A []int) []int {
	// 双指针
	var left, right = 0, len(A) - 1
	// 返回结果
	var res = make([]int, len(A))
	// 插入位置
	var p = len(A) - 1

	// 每次都把最大的放在后边, 然后根据插入的位置进行相应的变化
	for left <= right {
		if abs(A[left]) > abs(A[right]) {
			res[p] = A[left] * A[left]
			p--
			left++
		} else {
			res[p] = A[right] * A[right]
			p--
			right--
		}
	}
	return res
}

// 试试sort的性能
func sortedSquares(A []int) []int {
	for i := range A {
		A[i] *= A[i]
	}
	sort.Ints(A)
	return A
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

}
