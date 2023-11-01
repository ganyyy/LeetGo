package main

func maximumInvitations(favorite []int) int {
	n := len(favorite)

	// 统计基环树每个节点的入度
	degree := make([]int, n)
	for _, f := range favorite {
		degree[f]++
	}

	// reverseGraph := make([][]int, n) // 反图. a->b => b -> [a]
	var curQueue, nextQueue []int // 树枝队列
	for i, d := range degree {
		// 没有入度的就是树枝
		if d == 0 {
			curQueue = append(curQueue, i)
		}
	}

	maxDepth := make([]int, n)

	// 拓扑排序，剪掉图上所有树枝
	//  通过BFS递归处理, 只留下环
	//  因为是有向图, 所以 a->b->c->a 和  a->b->a 都是环
	for len(curQueue) > 0 {
		for _, x := range curQueue {
			y := favorite[x] // x 只有一条出边
			// 这里为啥不需要使用max(maxDepth[x], maxDepth[y]+1)呢?
			// 因为这是BFS, 所以较短的链会先被遍历到
			// 那么在遍历到较长的链时, 会覆盖掉较短的链
			maxDepth[y] = maxDepth[x] + 1
			if degree[y]--; degree[y] == 0 {
				// 消去树枝后自己也变成树枝了..?
				nextQueue = append(nextQueue, y)
			}
		}
		curQueue, nextQueue = nextQueue, curQueue[:0]
	}

	//
	/*最大的基环(节点数>2)
	  如果成环了, 那么就无法连接其他的树了
	  ① → ②
	  ↑   ↓       =>   4
	  ④ ← ③ ← 0
	*/

	/*最长的链(节点数=2的环的两个节点组成的树中最长边)之和.
	  多个没有交集的树之间可以在链的末尾相连, 所以需要累加所有不相交的树
	  ③ ←→ ① ← ⑥
	  ↑
	  ② ← 0       =>   6
	  ↑
	  ④ ← ⑤
	*/
	maxRingSize, sumChainSize := 0, 0
	for i, d := range degree {
		if d == 0 {
			continue
		}

		// 遍历基环上的点
		degree[i] = 0 // 将基环上的点的入度标记为 0，避免重复访问
		ringSize := 1 // 基环长度
		for x := favorite[i]; x != i; x = favorite[x] {
			degree[x] = 0 // 将基环上的点的入度标记为 0，避免重复访问
			ringSize++
		}

		if ringSize == 2 {
			// 基环长度为 2
			sumChainSize += maxDepth[i] + maxDepth[favorite[i]] + 2 // 两条链的长度之和+两个节点
		} else {
			// 取所有基环长度的最大值
			maxRingSize = max(maxRingSize, ringSize)
		}
	}
	return max(maxRingSize, sumChainSize)
}
