package main

func gardenNoAdj(n int, paths [][]int) []int {
	// 构建无向图
	var allPath = make([][]int, n+1)
	for _, path := range paths {
		a, b := path[0], path[1]
		allPath[a] = append(allPath[a], b)
		allPath[b] = append(allPath[b], a)
	}

	const AllFlower = 4
	// fmt.Println(allPath)

	// 所有直接相连的, 花都不能一致
	// 因为每个节点的出度最多是3...
	var ret = make([]int, n+1)
	for i := 1; i <= n; i++ {
		var flags [AllFlower + 1]bool
		for _, nxt := range allPath[i] {
			flags[ret[nxt]] = true
		}
		// fmt.Println(allPath[i], flags)
		for flower := 1; flower <= AllFlower; flower++ {
			if flags[flower] {
				continue
			}
			ret[i] = flower
			break
		}
	}
	return ret[1:]
}
