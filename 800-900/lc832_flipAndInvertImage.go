package main

func flipAndInvertImage(A [][]int) [][]int {
	// 这题目.. 太草
	// 送分题都没这么简单
	var n, m = len(A), len(A[0])
	for i := 0; i < n; i++ {
		var l, r = 0, m - 1
		for l <= r {
			A[i][l], A[i][r] = A[i][r]^1, A[i][l]^1
			l++
			r--
		}
	}

	return A
}
