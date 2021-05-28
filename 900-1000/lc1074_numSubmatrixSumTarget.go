package main

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	for i := 0; i < len(sum); i++ {
		sum[i] = make([]int, n+1)
	}
	for i := 0; i < len(matrix); i++ {
		t := 0
		for j := 0; j < len(matrix[i]); j++ {
			t += matrix[i][j]
			sum[i+1][j+1] = t + sum[i][j+1]
		}
	}
	count := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for x := 1; x <= i; x++ {
				for y := 1; y <= j; y++ {
					v := sum[i][j] - sum[x-1][j] - sum[i][y-1] + sum[x-1][y-1]
					if v == target {
						count++
					}
				}
			}
		}
	}
	return count
}

// 硬解就完事了
func numSubmatrixSumTarget2(matrix [][]int, target int) int {
	// 首行, 首列, 总计数
	var row = len(matrix)
	var col = len(matrix[0])
	var cnt int
	for i := 1; i < col; i++ {
		matrix[0][i] += matrix[0][i-1]
	}
	for i := 1; i < row; i++ {
		matrix[i][0] += matrix[i-1][0]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			matrix[i][j] += matrix[i-1][j] + matrix[i][j-1] - matrix[i-1][j-1]
		}
	}

	// 开始计数了呗..

	// 这个很耗时间的, 如果需要的话, 完全可以用上边的方法使用空间换时间
	var getVal = func(i1, i2, j1, j2 int) int {
		// 从零开始计数
		if i1 == 0 && j1 == 0 {
			return matrix[i2][j2]
		}
		var ret = matrix[i2][j2]
		if i1 > 0 {
			ret -= matrix[i1-1][j2]
		}
		if j1 > 0 {
			ret -= matrix[i2][j1-1]
		}
		if i1 > 0 && j1 > 0 {
			ret += matrix[i1-1][j1-1]
		}
		return ret
	}

	// 四重循环的暴力解法
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			for i2 := i; i2 < row; i2++ {
				for j2 := j; j2 < col; j2++ {
					if getVal(i, i2, j, j2) == target {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}

// 将二维问题转换为一维问题
func subarraySum(nums []int, k int) (ans int) {
	mp := map[int]int{0: 1}
	// 这里面相当于计算的左右边界的差值
	for i, pre := 0, 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := mp[pre-k]; ok {
			ans += v
		}
		mp[pre]++
	}
	return
}

func numSubmatrixSumTarget3(matrix [][]int, target int) (ans int) {
	sum := make([]int, len(matrix[0]))
	for i := range matrix { // 枚举上边界
		for _, row := range matrix[i:] { // 枚举下边界
			for c, v := range row {
				sum[c] += v // 更新每列的元素和
			}
			ans += subarraySum(sum, target)
		}
		for i := range sum {
			sum[i] = 0
		}
	}
	return
}
