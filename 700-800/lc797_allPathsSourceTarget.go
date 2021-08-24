package main

func allPathsSourceTarget(graph [][]int) [][]int {
	// 必须dfs啊

	var ret [][]int

	var dfs func(i int)
	var visited = make([]bool, len(graph))
	var cur = make([]int, 0, len(graph))
	dfs = func(i int) {
		if i == len(graph)-1 {
			cur = append(cur, i)
			var tmp = make([]int, len(cur))
			copy(tmp, cur)
			ret = append(ret, tmp)
			cur = cur[:len(cur)-1]
			return
		}
		visited[i] = true
		cur = append(cur, i)
		for _, p := range graph[i] {
			if visited[p] {
				continue
			}
			dfs(p)
		}
		cur = cur[:len(cur)-1]
		visited[i] = false
	}

	dfs(0)

	return ret
}
