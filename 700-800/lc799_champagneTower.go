package main

import "math"

func champagneTower(poured, queryRow, queryGlass int) float64 {
	row := make([]float64, 1, queryRow+1)
	row[0] = float64(poured)
	nextRow := make([]float64, 0, queryRow+1)
	for i := 1; i <= queryRow; i++ {
		nextRow = nextRow[:i+1]
		for idx := range nextRow {
			nextRow[idx] = 0
		}
		// 逐层计算, 逢1减半
		// 越靠近中间的杯子, 来源的水的数量就越多
		for j, volume := range row {
			// 当前行的杯子会分流给下一行 临近的杯子
			if volume > 1 {
				sub := (volume - 1) / 2
				nextRow[j] += sub
				nextRow[j+1] += sub
			}
		}
		row, nextRow = nextRow, row
	}
	return math.Min(1, row[queryGlass])
}
