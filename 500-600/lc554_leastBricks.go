package main

func leastBricks(wall [][]int) int {
	// 每一行求前缀和, 出现次数最多的那个就是 可以穿过最少墙的路径
	var m = make(map[int]int)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, row := range wall {
		var sum int
		// 不能算最后一个
		for _, v := range row[:len(row)-1] {
			sum += v
			// row[i] = sum
			m[sum]++
		}
	}

	var mv int

	// 找出最大的值
	for _, v := range m {
		mv = max(mv, v)
	}

	// 判断哪个没有这个值, 结果++

	// 有什么可以优化的吗?

	// 不需要知道是哪些具体的墙, 只需要知道数量即可!
	return len(wall) - mv
}
