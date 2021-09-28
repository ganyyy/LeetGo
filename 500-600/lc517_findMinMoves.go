package main

func findMinMoves(machines []int) (ans int) {
	tot := 0
	for _, v := range machines {
		tot += v
	}
	n := len(machines)
	// 不能平均分配, 直接返回即可
	if tot%n > 0 {
		return -1
	}
	// 计算每个洗衣机应有的衣服数量
	avg := tot / n
	sum := 0
	for _, num := range machines {
		num -= avg                         // 这个值是当前位置需要移入/移出的衣服的数量. 如果这个值非常高, 就相当于需要同时向左右移入/移出, 每件等同于一次
		sum += num                         // 这个值是从[0:i+1]这半部分相对于平均值需要移入/移出到另一半部分的衣服的数量. 这半部分的整体向另一部分移入/移出, 每件等同于一次
		ans = max(ans, max(abs(sum), num)) // 计算结果时, 取二者的最大值
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
