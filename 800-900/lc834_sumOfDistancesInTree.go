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

func sumOfDistancesInTree3(n int, edges [][]int) []int {
	// 暴力解法肯定是有的, 但是怎么运用dp呢?

	next := make([][]int, n)
	for _, edge := range edges {
		s, e := edge[0], edge[1]
		next[s] = append(next[s], e)
		next[e] = append(next[e], s)
	}

	ret := make([]int, n)

	// 以i节点作为根节点, 刨去父节点这个树的节点个数
	// 起始情况下, 每个节点都看成是一个单独的树, 对应的节点数量就是1
	size := make([]int, n)
	for i := range size {
		size[i] = 1
	}

	// start: 起点
	// father: 父节点
	// depth: 深度
	// 计算每个节点作为树的根节点时, 这个树对应的节点个数
	var dfs func(int, int, int)

	dfs = func(start, father, depth int) {
		ret[0] += depth // 0 到其他节点的累加和, 直接算
		for _, nxt := range next[start] {
			if nxt == father {
				continue
			}
			dfs(nxt, start, depth+1)
			size[start] += size[nxt]
		}
	}

	dfs(0, -1, 0) // 统计所有的节点的子树节点和

	fmt.Println(size)

	// start: 起点
	// father: 父节点
	var reRoot func(int, int)

	/*
	   对于A,B两个相连接的节点, 分别拆成两棵子树
	   sum(A): 所有子节点到A的距离和(注意: 这里是已经断开了A-B之间连接后的结果)

	   ret[A] = sum(A) + sum(B) + size[B] , 加一个size[B]是因为从B子树中任意一个点位到A都需要额外+1
	   ret[B] = sum(B) + sum(A) + size[A] , 加一个size[A]同理

	   size[A]+size[B] = N, 因为他们将整个树分成了两部分
	   ret[A] - ret[B] = size[B]-size[A] => ret[A] = ret[B]+size[B]-size[A] = ret[B] + N - 2*size[A]

	   假设现在已经获取了父节点对应的结果ret[P], 此时要求P节点所有的子节点对应的ret[X]
	   ret[X] = ret[P] + N - 2*size[X]

	   以此类推到所有的子节点和父节点之间的关系

	   0 已经在计算size的时候知道了结果诶.. 所以可以作为计算其他节点的初始值
	*/

	reRoot = func(start, father int) {
		for _, nxt := range next[start] {
			if nxt == father {
				continue
			}
			// 这个怎么理解很关键.
			ret[nxt] = ret[start] + n - 2*size[nxt]
			reRoot(nxt, start)
		}
	}
	reRoot(0, -1)

	return ret
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
