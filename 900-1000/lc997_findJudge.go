package main

func findJudgeBad(n int, trust [][]int) int {

	// 有向图
	// 如果和数据的内容无关, 可以直接转换成 计数 来处理

	var relation1 = make([][]int, n+1) // 我信任的人. 不清楚规则是否是自洽的, 或许也可以用set/bitmap?
	var relation2 = make([][]int, n+1) // 信任我的人

	for _, t := range trust {
		relation1[t[0]] = append(relation1[t[0]], t[1])
		relation2[t[1]] = append(relation2[t[1]], t[0])
	}

	// 1. 查找一下, 是不是有一个信任的人数是0的(有且只有一个)
	var master = -1

	for i, r := range relation1[1:] {
		if len(r) != 0 {
			continue
		}
		if master != -1 {
			// 超过两个大法官
			return -1
		}
		master = i + 1
	}

	if master == -1 {
		return -1
	}

	if len(relation2[master]) == n-1 {
		return master
	}
	return -1
}

func findJudge(n int, trust [][]int) int {
	// 法官一定是入度 = n-1, 出度 = 0
	// 出度-1, 入度+1
	inToValues := make([]int, n+1)
	for _, v := range trust {
		inToValues[v[0]]--
		inToValues[v[1]]++
	}
	for i := 1; i <= n; i++ {
		if inToValues[i] == n-1 {
			return i
		}
	}
	return -1
}
