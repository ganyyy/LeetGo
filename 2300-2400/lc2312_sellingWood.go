package main

func sellingWood(m, n int, prices [][]int) int64 {
	// 将小矩形转变成
	pr := make([][]int, m+1)
	for i := range pr {
		pr[i] = make([]int, n+1)
	}
	for _, price := range prices {
		pr[price[0]][price[1]] = price[2]
	}

	f := make([][]int64, m+1)
	// 从(1,1)开始到(m,n), 依次迭代子矩形的最大面积和
	for i := 1; i <= m; i++ {
		f[i] = make([]int64, n+1)
		for j := 1; j <= n; j++ {
			// 直接售卖
			f[i][j] = int64(pr[i][j])
			for k := 1; k < j; k++ { // 垂直切割
				f[i][j] = max(f[i][j], f[i][k]+f[i][j-k])
			}
			for k := 1; k < i; k++ { // 水平切割
				f[i][j] = max(f[i][j], f[k][j]+f[i-k][j])
			}
		}
	}
	return f[m][n]
}
