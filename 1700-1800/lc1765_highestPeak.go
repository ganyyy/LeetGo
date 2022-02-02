package main

var dir = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

const (
	SHIFT = 10
	MASK  = (1 << 10) - 1
)

func Pack(x, y int) int {
	return (x << SHIFT) | y
}

func Unpack(v int) (x, y int) {
	return v >> SHIFT, v & MASK
}

func highestPeak(isWater [][]int) [][]int {

	var queue []int
	for i, row := range isWater {
		for j, v := range row {
			if v == 1 {
				row[j] = 0
				queue = append(queue, Pack(i, j))
			} else {
				row[j] = -1
			}
		}
	}

	var li, lj = len(isWater), len(isWater[0])

	// 标记

	for len(queue) > 0 {
		// fmt.Println(queue)
		var top = queue[0]
		queue = queue[1:]

		var i, j = Unpack(top)
		var val = isWater[i][j]
		for _, d := range dir {
			var ni, nj = i + d[0], j + d[1]
			if ni < 0 || nj < 0 || ni >= li || nj >= lj {
				continue
			}
			var nv = isWater[ni][nj]
			if nv == 0 {
				// 这是一个水格子, 肯定已经在队列中了
				// 不需要重复处理
				continue
			}
			// 这个格子已经被标记了, 直接看下一个
			if nv >= 0 {
				continue
			}
			// 这是一个未分配陆地的格子, 所以需要特殊处理一下
			isWater[ni][nj] = val + 1
			queue = append(queue, Pack(ni, nj))

			// 这里不会存在一个当前格子的高度和相邻格子之间的高度差> 1的情况
			// 因为每次计算的轮次都是相同的, 由此决定的最大高度也是相同的, 由此扩散的最大值也是相同的
			// 所以不会出现相邻的两个格子差值超过1的情况
		}
	}

	return isWater
}
