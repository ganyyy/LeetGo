package main

func getRow(rowIndex int) []int {
	var res = make([]int, rowIndex+1)
	res[0] = 1
	// 二维dp优化到一维, 需要用到前态, 从后先前搞
	for i := 1; i <= rowIndex; i++ {
		res[i] = 1
		for j := i - 1; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}

func getRowMath(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	// 数学还是好用.
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}
