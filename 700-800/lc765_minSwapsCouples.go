package main

func minSwapsCouples(row []int) int {
	var tmp = make([]int, len(row))
	// 贪心算法

	// 逆转iv, 当字典用. 根据值取出队对应的下标位置
	for i, v := range row {
		tmp[v] = i
	}

	var cnt int
	for i := 0; i < len(row); i += 2 {
		var f = friend(row[i])
		if f != row[i+1] {
			var old = row[i+1]
			// 交换值
			row[i+1], row[tmp[f]] = row[tmp[f]], row[i+1]
			// 交换位置
			tmp[f], tmp[old] = tmp[old], tmp[f]
			cnt++
		}
	}

	return cnt
}

func minSwapsCouples2(row []int) int {
	n, res := len(row), 0
	for i := 0; i < n; i += 2 {
		x := row[i]
		if row[i+1] == x^1 {
			continue
		}
		res++
		// 只有可能是后边的有问题
		for j := i + 1; j < n; j++ {
			if row[j] == x^1 {
				row[i+1], row[j] = row[j], row[i+1]
				break
			}
		}
	}
	return res
}

func minSwapsCouples3(row []int) int {
	// 并查集. 忘得差不多了
	var ln = len(row)
	var fa = make([]int, ln>>1)
	// 初始化一下, 每一对的父节点指向自己
	for i := range fa {
		fa[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if i != fa[i] {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}

	var cnt int
	for i := 0; i < ln>>1; i++ {
		var a, b = find(row[i<<1] >> 1), find(row[(i<<1)+1] >> 1)
		if a != b {
			// 这一对不是情侣, 需要互换一下
			if b < a {
				a, b = b, a
			}
			fa[a] = b
			cnt++
		}
	}

	return cnt
}

func friend(a int) int {
	return a ^ 1
}
