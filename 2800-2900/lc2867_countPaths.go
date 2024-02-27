package main

const Max int = 1e5 + 1

var np = [Max]bool{1: true}

func init() { // 质数=false 非质数=true
	// 从2开始, 2的倍数都是非质数. 2,4,6,8,10,12,14,16,18,20...
	// 然后从3开始, 3的倍数都是非质数. 3,6,9,12,15,18,21,24,27,30...
	// 在计算4的时候, 4已经被2计算过了. 4,8,12,16,20,24,28,32,36,40...
	for i := 2; i*i < Max; i++ {
		if !np[i] {
			for j := i * i; j < Max; j += i {
				np[j] = true
			}
		}
	}
}

func countPaths(n int, edges [][]int) (ans int64) {
	// 构建一个邻接表
	graph := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	size := make([]int, n+1)
	// 每个节点的质数节点组成的集合
	var nodes []int
	var dfs func(int, int)
	dfs = func(x, fa int) {
		// 找到所有的非质数点位, 加入到集合中
		nodes = append(nodes, x)
		for _, y := range graph[x] {
			if y != fa && np[y] {
				dfs(y, x)
			}
		}
	}
	/*
	        3
	   (2) (3) (4)

	   根节点是3, 有三个子节点, 这些子节点中的质数的个数分别是2,3,4个
	   从左开始, (3)和(2)两两组合其路径个数为6, (4)和(3)(2)两两组合其路径个数为24个.

	   最后再加上3个单向的路径, 3->(2), 3->(3), 3->(4) = 9个
	*/
	for x := 1; x <= n; x++ {
		if np[x] { // 跳过非质数
			continue
		}
		sum := 0
		for _, y := range graph[x] { // 质数 x 把这棵树分成了若干个连通块
			if !np[y] {
				continue
			}
			if size[y] == 0 { // 尚未计算过
				nodes = []int{}
				dfs(y, -1) // 遍历 y 所在连通块，在不经过质数的前提下，统计有多少个非质数
				for _, z := range nodes {
					// 当size[z]为根的时候, 它的质数节点个数和这个相同(!)
					size[z] = len(nodes)
				}
			}
			// 这 size[y] 个非质数与之前遍历到的 sum 个非质数，两两之间的路径只包含质数 x
			ans += int64(size[y]) * int64(sum)
			sum += size[y]
		}
		// 单向的路径
		ans += int64(sum) // 从 x 出发的路径
	}
	return
}
