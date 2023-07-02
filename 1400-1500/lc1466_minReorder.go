package main

func minReorder(n int, connections [][]int) int {
	// [i] -> [j]
	// 单向
	// 边的数量为 n-1
	// 这尼玛, 从0出发, 看有几个正边?
	adj := map[int][]int{}
	for _, conn := range connections {
		// a -> b 是正数, b -> a 是负数.
		// 这里面要做的就是把所有的边都变成负数, 所以最后的结果是所有的正数边的数量
		adj[conn[0]] = append(adj[conn[0]], conn[1])
		adj[conn[1]] = append(adj[conn[1]], -conn[0])
	}
	ans := 0
	var dfs func(u, p int)
	dfs = func(u, p int) {
		for _, v := range adj[u] {
			// 单向的, 不存在依赖, 只需要判断是否是父节点就行了
			if v != p && -v != p {
				if v > 0 {
					ans++
				} else {
					v = -v
				}
				dfs(v, u)
			}
		}
	}
	dfs(0, n)
	return ans
}
