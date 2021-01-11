package main

func topSort(graph [][]int, deg, items []int) (orders []int) {
	// 一次拓扑排序
	// 首先找到所有入度为0的点
	// 然后依次遍历对应指向的点集合
	var q []int
	for _, i := range items {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		orders = append(orders, from)
		for _, to := range graph[from] {
			deg[to]--
			if deg[to] == 0 {
				q = append(q, to)
			}
		}
	}
	return
}

func sortItems(n, m int, group []int, beforeItems [][]int) (ans []int) {
	groupItems := make([][]int, m+n) // groupItems[i] 表示第 i 个组负责的项目列表
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i // 给不属于任何组的项目分配一个组
		}
		// 初始化每一个组对应的项目
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	// 为啥是m+n呢? 这是一个最大的数量. 即每一个项目都没有分配到组的情况
	// 此时相当于最大m+n个组
	groupGraph := make([][]int, m+n) // 组间依赖关系
	groupDegree := make([]int, m+n)
	itemGraph := make([][]int, n) // 组内依赖关系
	itemDegree := make([]int, n)
	for cur, items := range beforeItems {
		// 当前的组
		curGroupID := group[cur]
		for _, pre := range items {
			preGroupID := group[pre]
			if preGroupID != curGroupID { // 不同组项目，确定组间依赖关系
				groupGraph[preGroupID] = append(groupGraph[preGroupID], curGroupID)
				groupDegree[curGroupID]++
			} else { // 同组项目，确定组内依赖关系
				itemGraph[pre] = append(itemGraph[pre], cur)
				itemDegree[cur]++
			}
		}
	}

	// 组间拓扑序
	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}
	groupOrders := topSort(groupGraph, groupDegree, items)
	if len(groupOrders) < len(items) {
		return nil
	}

	// 按照组间的拓扑序，依次求得各个组的组内拓扑序，构成答案
	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := topSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}
	return
}
