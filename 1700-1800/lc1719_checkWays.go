package main

import "math"

func checkWays(pairs [][]int) int {
	// 第一步: 搭建不同节点之间的映射关系
	adj := map[int]map[int]bool{}
	for _, p := range pairs {
		x, y := p[0], p[1]
		if adj[x] == nil {
			adj[x] = map[int]bool{}
		}
		adj[x][y] = true
		if adj[y] == nil {
			adj[y] = map[int]bool{}
		}
		adj[y][x] = true
	}

	// 检测是否存在根节点
	// 注意: pair表示的是一组父字节点关系, 如果 存在一组父子关系
	// 不一定是直系的!
	// 所以根节点一定是所有节点的父节点, 即和根节点相关联的节点的个数为总结点数量-1
	root := -1
	for node, neighbours := range adj {
		if len(neighbours) == len(adj)-1 {
			root = node
			break
		}
	}
	if root == -1 {
		return 0
	}

	ans := 1
	for node, neighbours := range adj {
		if node == root {
			continue
		}
		// 迭代每一个子节点对应的关系节点
		currDegree := len(neighbours)
		parent := -1
		parentDegree := math.MaxInt32
		// 根据 degree 的大小找到 node 的父节点 parent
		// 这里找的所有关系节点中, 相连度最小的父节点
		for neighbour := range neighbours {
			if len(adj[neighbour]) < parentDegree && len(adj[neighbour]) >= currDegree {
				parent = neighbour
				parentDegree = len(adj[neighbour])
			}
		}
		if parent == -1 {
			return 0
		}
		// 检测 neighbours 是否为 adj[parent] 的子集
		for neighbour := range neighbours {
			if neighbour != parent && !adj[parent][neighbour] {
				return 0
			}
		}

		if parentDegree == currDegree {
			// 这里不直接返回, 是因为后续可能存在不符合要求的节点
			ans = 2
		}
	}
	return ans
}
