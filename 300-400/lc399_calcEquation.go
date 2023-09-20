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
	// 并查集
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
	// 起始情况下, w[a], w[b], w[c] = 1
	// 如果a/b=2, 那么w[a]=2*w[b]=2
	// 如果c/a=3, 那么w[c]=3*(2*w[b])=6*w[b]=6
	var w = make([]float64, len(id))

	// 初始化关系
	for i := 0; i < len(id); i++ {
		fa[i] = i
		w[i] = 1 // 自己到自己的权重为1, 相当于自己/自己=1
	}

	// 寻找父节点, 同时更新权重
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			// 寻找父节点, 直到找到根节点为止
			var f = find(fa[x])
			// 通过父节点的权重更新当前节点的权重
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}

	var merge = func(from, to int, val float64) {
		// 找到from和to的根节点
		fFrom, fTo := find(from), find(to)
		// 合并两个根节点时, 更新权重
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

func calcEquation3(equations [][]string, values []float64, queries [][]string) []float64 {
	if len(queries) == 0 {
		return nil
	}

	// 给每个数组编号
	var identity = make(map[string]int, len(queries))
	add := func(v string) {
		if _, ok := identity[v]; ok {
			return
		}
		identity[v] = len(identity)
	}
	id := func(v string) (r int, ok bool) {
		r, ok = identity[v]
		return
	}

	for _, equation := range equations {
		add(equation[0])
		add(equation[1])
	}

	// 初始化父子关系和权重比值
	var parents = make([]int, len(identity))
	var weight = make([]float64, len(identity))

	for _, id := range identity {
		parents[id] = id
		weight[id] = 1
	}

	var findParent func(int) int
	findParent = func(id int) int {
		p := parents[id]
		if p != id {
			np := findParent(p)
			// 为什么可以这么做呢?
			// 最终结果上, 根节点的weight一定是1
			// 假设先上来的链式关系是 a->b->c->d, 那么压平后就是 a->d, b->d, c->d
			// 首次需要通过累乘获取到 a 和 d的关系(weight[c]*weight[b]*weight[a]),
			// 后续 a的直接父节点本身就是d, 那么此时的 *= 相当于没发生变化
			// 同样道理, 如果 d 又出现了新的父节点e, 那么就会重新累乘到a身上, 同时会更新 a/d和e之间的关系
			weight[id] *= weight[p]
			p = np
			parents[id] = p
		}
		return p
	}
	// [②]/[①] = 2 => parents[②] = ①, weight[②] = 2, weight[①] = 1
	// [④]/[③] = 4 => parents[④] = ③, weight[④] = 4, weight[③] = 1
	// [④]/[②] = 2 => parents[③] = ①, weight[③] = 2, weight[①] = 1
	var merge = func(from, to int, value float64) {
		pf, pt := findParent(from), findParent(to)
		if pf == pt {
			// 后续不用看了, 因为最后的计算结果一定是不变的
			// 否则就是出现了错误的倍数关系
			return
		}
		// 现在相当于连上两个之前不相关的
		// 这两个关系要怎么计算呢? 首先, weight[pf]和weight[pt]此时一定都是1
		// 那么可得
		// from = weight[from]*weight[pf]
		// to   = weight[to]  *weight[pt]
		// from/to = (weight[from]*weight[pf])/(weight[to]  *weight[pt]) = value
		// (weight[pf]/weight[pt]) = value * weight[to]/weight[from]
		// 令weight[pt]保持为1
		// weight[pf] = value * weight[to]/weight[from]
		weight[pf] = value * weight[to] / weight[from]
		parents[pf] = pt
	}

	// 基于 equations 重构权重比值
	for i, eq := range equations {
		id1, ok := id(eq[0])
		if !ok {
			continue
		}
		id2, ok := id(eq[1])
		if !ok {
			continue
		}
		merge(id1, id2, values[i])
	}
	ret := make([]float64, 0, len(queries))
	// 基于 queries 生成结果集合
	for _, query := range queries {
		id1, ok1 := id(query[0])
		id2, ok2 := id(query[1])
		if ok1 && ok2 && findParent(id1) == findParent(id2) {
			ret = append(ret, weight[id1]/weight[id2])
		} else {
			ret = append(ret, -1)
		}
	}
	return ret
}
