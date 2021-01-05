package main

import "fmt"

// 此路不通, 需要进行连锁更新
func calcEquationErr(equations [][]string, values []float64, queries [][]string) []float64 {
	// 先把每个式子的值求出来
	var m = make(map[string]float64, len(equations)*2)
	var v, v1, v2 float64
	var ok bool
	for i, equation := range equations {
		v = values[i]
		v1, ok = m[equation[0]]
		if !ok {
			v1 = 1.0
		}
		v2, ok = m[equation[1]]
		if !ok {
			v2 = v1 / v
		} else {
			v1 = v2 * v
		}
		m[equation[0]] = v1
		m[equation[1]] = v2
	}

	fmt.Println(m)

	var res = make([]float64, len(queries))

	for i, query := range queries {
		v1, ok = m[query[0]]
		if !ok {
			res[i] = -1.0
			continue
		}
		v2, ok = m[query[1]]
		if !ok {
			res[i] = -1.0
			continue
		}

		res[i] = v1 / v2
	}

	return res
}

// 使用并查集进行处理
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给每一个字母进行编号
	var id = make(map[string]int, len(equations)*2)
	for _, equation := range equations {
		var a, b = equation[0], equation[1]
		if _, ok := id[a]; !ok {
			id[a] = len(id)
		}
		if _, ok := id[b]; !ok {
			id[b] = len(id)
		}
	}

	// 关系数组和权重数组
	var fa = make([]int, len(id))
	var w = make([]float64, len(id))

	// 初始化关系
	for i := 0; i < len(id); i++ {
		fa[i] = i
		w[i] = 1
	}

	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			// 寻找父节点, 直到找到根节点为止
			var f = find(fa[x])
			// 更新权重
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}

	var merge = func(from, to int, val float64) {
		// 找到from和to的根节点
		fFrom, fTo := find(from), find(to)
		w[fFrom] = val * w[to] / w[from]
		// 合并
		fa[fFrom] = fTo
	}

	// 合并有关系的选项
	for i, eq := range equations {
		merge(id[eq[0]], id[eq[1]], values[i])
	}

	var ans = make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]

		if hasS && hasE && find(start) == find(end) {
			ans[i] = w[start] / w[end]
		} else {
			ans[i] = -1
		}
	}

	return ans
}
