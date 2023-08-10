package main

func calculateMinimumHP(dungeon [][]int) int {
	ln := len(dungeon)
	if ln == 0 {
		return 1
	}
	lm := len(dungeon[0])
	if lm == 0 {
		return 1
	}
	// 从终点向起点逆推
	dungeon[ln-1][lm-1] = max(1, 1-dungeon[ln-1][lm-1])
	// 处理边界

	// 最后一列
	lasCol, lasRow := lm-1, ln-1
	for i := ln - 2; i >= 0; i-- {
		dungeon[i][lasCol] = max(1, dungeon[i+1][lasCol]-dungeon[i][lasCol])
	}
	// 最后一行
	for i := lm - 2; i >= 0; i-- {
		dungeon[lasRow][i] = max(1, dungeon[lasRow][i+1]-dungeon[lasRow][i])
	}

	// 看其他的
	for i := ln - 2; i >= 0; i-- {
		for j := lm - 2; j >= 0; j-- {
			// 当前位置可以从 右边过来, 也可以从下边过来
			// 这里取二者消耗血量的最小值, 如果小于0 那就取 1
			dungeon[i][j] = max(1, min(dungeon[i][j+1]-dungeon[i][j], dungeon[i+1][j]-dungeon[i][j]))
		}
	}
	return dungeon[0][0]
}
