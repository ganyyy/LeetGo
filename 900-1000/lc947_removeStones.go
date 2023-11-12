package main

import "fmt"

func removeStones(stones [][]int) int {
	// 搞一个矩阵..?
	// 还他妈并查集???

	// 每一行, 每一列只保留一个.
	// 换而言之, 可以通过 dfs 方式遍历所有点.

	var n = len(stones)
	// 边之间的关系
	var edge = make([][]int, n)
	// 同一行/列之间的点位
	var rec = make(map[int][]int, n)

	for i, s := range stones {
		rec[s[0]] = append(rec[s[0]], i)
		// 区分横坐标和纵坐标
		rec[s[1]+10000] = append(rec[s[1]+10000], i)
	}

	for i, points := range rec {
		// 同一行/列的数据相互管理
		fmt.Println(i, points)
		for i := 1; i < len(points); i++ {
			edge[points[i-1]] = append(edge[points[i-1]], points[i])
			edge[points[i]] = append(edge[points[i]], points[i-1])
		}
	}

	var visiable = make([]bool, n)

	var dfs func(i int)

	dfs = func(i int) {
		visiable[i] = true
		for _, p := range edge[i] {
			if visiable[p] {
				continue
			}
			dfs(p)
		}
	}

	var num int
	for i := 0; i < n; i++ {
		if visiable[i] {
			continue
		}
		// 最终要保留的点
		num++
		dfs(i)
	}

	// 删除的点的数量
	return n - num

}

func removeStones2(stones [][]int) int {
	// 搞一个矩阵..?
	// 还他妈并查集???

	// 底层数组使用map表示
	// rank表示权重
	var n = len(stones)
	var f, rank = make(map[int]int, n), make(map[int]int, n)

	var find func(i int) int
	find = func(i int) int {
		if v, ok := f[i]; !ok {
			f[i] = i
			rank[i] = 1
			return i
		} else {
			if v == i {
				return v
			}
			return find(v)
		}
	}

	var union = func(a, b int) {
		var fa, fb = find(a), find(b)
		if fa == fb {
			return
		}
		// fa始终是大的一方
		if rank[fa] < rank[fb] {
			fa, fb = fb, fa
		}
		rank[fa] += rank[fb]
		// 将fb合并到fa
		f[fb] = fa
	}

	for _, s := range stones {
		union(s[0], s[1]+10000)
	}

	var num int
	for a, fa := range f {
		if a == fa {
			num++
		}
	}

	return n - num
}

func removeStones3(stones [][]int) int {
	arr := make([]int, len(stones))
	for i := range arr {
		arr[i] = i
	}
	// 行, 列map
	rowMp := make(map[int]int, len(stones))
	colMp := make(map[int]int, len(stones))
	for k, v := range stones {
		// 分别合并行和列
		row, ok1 := rowMp[v[0]]
		col, ok2 := colMp[v[1]]
		if ok1 {
			UnionStones947(arr, row, k)
		} else {
			rowMp[v[0]] = k
		}
		if ok2 {
			UnionStones947(arr, col, k)
		} else {
			colMp[v[1]] = k
		}
	}
	res := 0
	// 最后计算剩余的点的个数
	for k, v := range arr {
		if k != v {
			res++
		}
	}
	return res
}

func UnionStones947(arr []int, i, j int) int {
	a, b := findUnion947(arr, i), findUnion947(arr, j)
	if a != b {
		arr[b] = a
		return 1
	}
	return 0
}

func findUnion947(arr []int, i int) int {
	for arr[i] != i {
		i = arr[i]
	}
	return i
}
