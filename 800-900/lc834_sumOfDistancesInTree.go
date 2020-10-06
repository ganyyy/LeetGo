package main

import "fmt"

func sumOfDistancesInTree2(N int, edges [][]int) []int {
	var res = make([]int, N)
	// 首先要构建一个映射
	// 这个映射是所有从该点出发可以 到达的所有点
	// 因为是无环的无相联通图, 所以 N个点一定有N-1条边
	var p = make([][]int, N)
	var empty = make([]int, N)
	for _, v := range edges {
		p1, p2 := v[0], v[1]
		p[p1] = append(p[p1], p2)
		p[p2] = append(p[p2], p1)
	}
	// 找出这个点所有连接点的距离
	var dfs func(i, now int) int

	dfs = func(i, now int) int {
		var cur int
		for _, point := range p[i] {
			if empty[point] == 1 {
				cur += now
				continue
			}
			empty[point] = 1
			cur += dfs(point, now+1)
			empty[point] = 0
		}
		return cur
	}

	// 然后, emmm, 暴力dfs 试试行不行
	for i := 0; i < N; i++ {
		empty[i] = 1
		res[i] = dfs(i, 0)
		empty[i] = 0
	}
	return res
}

func sumOfDistancesInTree(N int, edges [][]int) []int {
	// 构建一个图
	var graph = make([][]int, N)
	for _, e := range edges {
		var u, v = e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// sz[u]表示以u为root节点, 它的所有子节点的数量之和(包括自己 )
	var sz = make([]int, N)
	// dp[u]表示以u为root节点, 它的所有子节点到它的距离之和
	var dp = make([]int, N)

	var ans = make([]int, N)

	// u: 当前节点
	// f: 父节点
	var dfs func(u, f int)
	dfs = func(u, f int) {
		// 很容易想到一个树形dp(划重点, 很容易)  dp[u] = sum(dp[v]+sz[v] for v in son[u])
		// 首先, 先取得所有子节点的距离和, 在加上 u->v , 一共 sz[u]条距离为1的边
		sz[u] = 1
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			dfs(v, u)
			dp[u] += dp[v] + sz[v]
			sz[u] += sz[v]
		}
	}
	// 先进行以此遍历
	dfs(0, -1)

	var dfs2 func(u, f int)
	dfs2 = func(u, f int) {
		ans[u] = dp[u]
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			var pu, pv = dp[u], dp[v]
			var su, sv = sz[u], sz[v]
			// 这里进行了换根操作, 将根由 u -> v
			// 首先, v 不属于 u的子节点, 所以要在 dp[u]中去掉 v相关的数据,
			// 同样的, 子节点也不能包含 v相关的子节点
			dp[u] -= dp[v] + sz[v]
			sz[u] -= sz[v]
			// 对于 v而言, 需要加上 u相关的数据,
			// 子节点也要加上 u的所有子节点数量
			dp[v] += dp[u] + sz[u]
			sz[v] += sz[u]

			dfs2(v, u)
			// 复原
			dp[u], dp[v] = pu, pv
			sz[u], sz[v] = su, sv
		}
	}
	dfs2(0, -1)
	return ans
}

func main() {
	var N = 6
	var edges = [][]int{
		{0, 1},
		{0, 2},
		{2, 3},
		{2, 4},
		{2, 5},
	}
	fmt.Println(sumOfDistancesInTree(N, edges))
}
