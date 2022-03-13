package main

func platesBetweenCandles(s string, queries [][]int) []int {

	var cnt = make([]int, len(s)) // * 的计数
	var left, right = make([]int, len(s)+1), make([]int, len(s)+1)

	// 一个左右缩进的问题
	// 试想一下:
	//   |***|***|***|
	//cnt: 累计的*的个数                 [ 0, 1, 2, 3, 0, 4, 5, 6, 0, 7, 8, 9, 0]
	//left : 从左向右, 第一个 | 出现的位置 [ 0, 0, 0, 0, 4, 4, 4, 4, 8, 8, 8, 8,12]
	//right: 从右向左, 后续的 | 出现的位置 [ 0, 4, 4, 4, 4, 8, 8, 8, 8,12,12,12,12]
	//二者在 | 出现的位置上有重合

	var cur int
	for i, v := range s {
		left[i+1] = left[i]
		if v == '*' {
			cur++
		} else {
			left[i+1] = i
		}
		cnt[i] = cur
	}
	for i := len(s) - 1; i >= 0; i-- {
		right[i] = right[i+1]
		if s[i] == '|' {
			right[i] = i
		}
	}

	left = left[1:]
	right = right[:len(right)-1]

	// fmt.Println(left, right)
	// fmt.Println(cnt)

	var ret = make([]int, len(queries))
	for i, query := range queries {
		// 位置相同, 肯定是0
		if query[0] == query[1] {
			continue
		}
		// 通过right向后缩进
		// 通过left 向前缩进
		var l = right[query[0]]
		var r = left[query[1]]
		// fmt.Println(l, r, cnt[l], cnt[r])
		// 如果在一个区间也不需要处理
		if l >= r {
			continue
		}
		// 最终的结果就是: 两个边界值对应的差值
		ret[i] = cnt[r] - cnt[l]
	}
	return ret
}
