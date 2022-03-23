package main

func imageSmoother(img [][]int) [][]int {
	var m = len(img)
	if m == 0 {
		return img
	}
	var n = len(img[0])
	if n == 0 {
		return img
	}
	var ret = make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var cnt, sum int

			for row := i - 1; row <= i+1; row++ {
				if row < 0 || row >= m {
					continue
				}
				for col := j - 1; col <= j+1; col++ {
					if col < 0 || col >= n {
						continue
					}
					cnt++
					sum += img[row][col]
				}
			}
			// fmt.Println(i, j, cnt, sum)
			ret[i][j] = sum / cnt
		}
	}
	return ret
}
