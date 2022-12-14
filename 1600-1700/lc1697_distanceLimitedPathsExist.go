package main

import "sort"

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool { return edgeList[i][2] < edgeList[j][2] })

	// 并查集模板
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) {
		fa[find(from)] = find(to)
	}

	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	// 按照 limit 从小到大排序，方便离线
	sort.Slice(queries, func(i, j int) bool { return queries[i][2] < queries[j][2] })

	ans := make([]bool, len(queries))
	k := 0
	for _, q := range queries {
		// 按照limit进行merge, 如果可以merge成功, 那么二者之间互联的边一定是小于等于limit的
		for ; k < len(edgeList) && edgeList[k][2] < q[2]; k++ {
			merge(edgeList[k][0], edgeList[k][1])
		}
		// 因为limit已经是按照从小到大进行排序后的序列, 所以可以方便的确认是否可以在 <= limit的距离下, 进行联通
		ans[q[3]] = find(q[0]) == find(q[1])
	}
	return ans
}
