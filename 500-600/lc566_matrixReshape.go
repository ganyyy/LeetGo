package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	var n, m = len(nums), len(nums[0])
	if n*m != r*c {
		return nums
	}

	var res = make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}

	for i := 0; i < m*n; i++ {
		res[i/c][i%c] = nums[i/m][i%m]
	}

	return res
}
