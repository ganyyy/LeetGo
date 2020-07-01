package main

import "fmt"

// // 这里利用了值二分法 找指定的第K小的元素
// func kthSmallest(matrix [][]int, k int) int {
//     n := len(matrix)
//     left, right := matrix[0][0], matrix[n-1][n-1]
//     for left < right {
//         mid := left + (right - left) / 2
//         // 如果数组中 比 mid 小的数个数大于 k,
//         // 说明比k小的数一定小于 mid, 否则一定大于 mid
//         if check(matrix, mid, n) >= k {
//             right = mid
//         } else {
//             left = mid + 1
//         }
//     }
//     return left
// }

// // 检查数组中比mid小的数的个数
// func check(matrix [][]int, mid, n int) int {
//     // 从左下角开始查询
//     i, j := n - 1, 0
//     num := 0
//     // 该二维数组每一列都是有序的
//     for i >= 0 && j < n {
//         // 如果mid比 当前值要大
//         if matrix[i][j] <= mid {
//             // 直接加上比这个列要小的所有数的个数
//             num += i + 1
//             // 换下一列
//             j++
//         } else {
//             // 上升一行
//             i--
//         }
//     }
//     return num
// }

func kthSmallest(matrix [][]int, k int) int {
	// 值二分法
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	// 在矩阵中找出比v小的数
	var check = func(v int) int {
		// 从左下角找起
		i, j := n-1, 0
		var cnt int
		for i >= 0 && j < n {
			if matrix[i][j] <= v {
				// 加上当前的列数
				cnt += i + 1
				j++
			} else {
				i--
			}
		}
		return cnt
	}

	for left < right {
		// 如何保证 mid 一定是在 矩阵中的值呢?
		mid := left + (right-left)>>1
		if check(mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return left
}
