package main

func smallestMissingValueSubtree(parents []int, genes []int) []int {
	length := len(parents)
	var children = make([][]int, length)
	for i := 1; i < length; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}

	// 将整棵树分成两部分:
	// 根节点带有1的部分, 根节点不带有1的部分
	//  根节点带有1的部分, 需要递归处理
	//  根节点不带有1的部分, 结果就是1
	res := make([]int, length)
	// 所有已经被标记的基因
	geneMarked := make(map[int]bool, length)
	// 标记树上节点是否已经被访问过
	rootVisited := make([]bool, length)
	// 首次出现的基因为1的根节点, 自底向上处理
	var minGeneParent = -1
	for node := 0; node < length; node++ {
		// 如果没有基因为1的节点, 那么所有的节点的最小缺失值都是1
		res[node] = 1
		if genes[node] == 1 {
			minGeneParent = node
		}
	}

	// 从根节点开始, 递归处理所有的子节点
	var visitNode func(int)
	visitNode = func(node int) {
		if rootVisited[node] {
			return
		}
		// 节点已被访问
		rootVisited[node] = true
		// 标记当前节点的基因
		geneMarked[genes[node]] = true
		// 处理所有的子节点
		for _, child := range children[node] {
			visitNode(child)
		}
	}

	// 最小的未标记基因
	var minGene = 1

	// 以首次出现的minGeneNode为子节点, 向上不停查找父节点, 直到找到根节点
	for minGeneParent != -1 {
		visitNode(minGeneParent)
		// 原理就是:
		//  标记当前以 minGeneParent 为根节点所有子节点的基因,
		//  通过递增的形式找到首个未标记的最小基因
		for geneMarked[minGene] {
			minGene++
		}
		res[minGeneParent] = minGene
		// 再找当前 minGeneParent 的父节点
		minGeneParent = parents[minGeneParent]
	}
	return res
}
