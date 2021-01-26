package main

type union struct {
	parent []int
	cnt    int
}

func newUnion(n int) *union {
	var u = union{
		parent: make([]int, n),
		cnt:    n, // 初始情况下看成n的独立的边
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}

	return &u
}

func (u *union) find(x int) int {
	if x != u.parent[x] {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *union) isSet(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u *union) merge(a, b int) bool {
	a, b = u.find(a), u.find(b)
	if a == b {
		return false
	}
	u.parent[b] = a
	u.cnt-- // 每次成功合并两条边, 独立边就-1
	return true
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	// 构建A, B两人的并查集
	var ua = newUnion(n)

	var ans = len(edges)
	// 先看公共边
	for _, edge := range edges {
		var x, y = edge[1] - 1, edge[2] - 1
		// 如果存在一条二者可能需要的公共边, 那么这条边一定不能删除
		if edge[0] == 3 && !ua.isSet(x, y) {
			ua.merge(x, y)
			ans--
		}
	}

	var ub = &union{
		parent: make([]int, n),
		cnt:    ua.cnt,
	}
	copy(ub.parent, ua.parent)

	// 再看独立边
	for _, edge := range edges {
		var tp = edge[0]
		if tp == 3 {
			continue
		}
		var t = ua
		if tp == 2 {
			t = ub
		}
		if t.merge(edge[1]-1, edge[2]-1) {
			// 如果有一方需要这条独占边, 就-1
			ans--
		}
	}

	// 无法联通
	if ua.cnt > 1 || ub.cnt > 1 {
		return -1
	}
	// 可以删除的边的数量
	return ans
}
