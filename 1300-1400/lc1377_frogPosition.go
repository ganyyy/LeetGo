package main

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	// 计算每一个点的可达次数?

	// 树说明一定不会有环
	// 第一步: 建树

	var allSub = make([][]int, n+1)
	for _, edge := range edges {
		start, end := edge[0], edge[1]
		allSub[start] = append(allSub[start], end)
		allSub[end] = append(allSub[end], start)
	}

	// 代表递归的深度
	var dfs func(root, d, from int, n float64) bool
	var ret float64
	dfs = func(root, d, from int, n float64) bool {
		// fmt.Println(root, d, n)
		if d < 0 {
			return false
		}
		nl := len(allSub[root])
		if root != 1 {
			// root != 1, 说明是从其他点转移过来的
			// root == 1, 说明这是起点
			nl--
		}
		if nl == 0 || d == 0 {
			// 没有子节点了
			if root == target {
				ret = n
				return true
			}
			return false
		}
		nn := n / float64(nl)
		for _, nxt := range allSub[root] {
			if nxt == from {
				continue
			}
			if dfs(nxt, d-1, root, nn) {
				return true
			}
		}
		return false
	}
	dfs(1, t, 0, 1)
	return ret
}
