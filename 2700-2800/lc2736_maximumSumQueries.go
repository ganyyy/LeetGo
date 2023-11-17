package main

import (
	"cmp"
	"slices"
	"sort"
)

func maximumSumQueries(nums1, nums2 []int, queries [][]int) []int {
	// 相同位置的nums1和nums2的值
	type pair struct{ x, y int }
	pairs := make([]pair, len(nums1))
	for i, x := range nums1 {
		pairs[i] = pair{x, nums2[i]}
	}
	// 注意这里比较的是b和a的x, 所以等同于按照x降序排列
	slices.SortFunc(pairs, func(a, b pair) int { return cmp.Compare(b.x, a.x) })

	// 查询id, 后续进一步排序处理
	qid := make([]int, len(queries))
	for i := range qid {
		qid[i] = i
	}

	// 对比的是j和i的x, 所以等同于按照查询的x降序排列
	slices.SortFunc(qid, func(i, j int) int { return cmp.Compare(queries[j][0], queries[i][0]) })

	ans := make([]int, len(queries))
	type data struct{ y, s int }

	/*
		1. pairs 按照 x 降序排列
		2. query x 降序排列
		3. st 中 total 降序排列, y 升序排列
	*/

	// 单调栈, 双指针, 二分查询
	// pairs按照x降序排列, qid按照查询的x降序排列, 通过最终运算保证了[y]是递增的
	var st []data
	j := 0
	for _, i := range qid {
		x, y := queries[i][0], queries[i][1]
		// 入栈条件:
		// 1. pairs[j].x >= x
		// 2. pairs[j].y >= st[len(st)-1].y
		// 3. pairs[j].x+pairs[j].y >= st[len(st)-1].s

		// 先将所有>=x的pairs[j]的{y, x+y}放入栈中
		// st是一个单调递减栈, 保证了st中的元素的x+y是单调递减的
		for ; j < len(pairs) && pairs[j].x >= x; j++ { // 下面只需关心 pairs[j].y
			// 因为pairs.[x]本身已经是单调递减的, 所以如果total(pairs[j]) >= total(st[len(st)-1]),
			// 那么必定pairs[j]的y更大
			for len(st) > 0 && st[len(st)-1].s <= pairs[j].x+pairs[j].y { // pairs[j].y >= st[len(st)-1].y
				st = st[:len(st)-1]
			}
			// 在上一个循环结束后, st中的元素的x+y是单调递减的,
			// 如果pairs[j].y > st[len(st)-1].y, 那么pairs[j]的y更大,
			if len(st) == 0 || st[len(st)-1].y < pairs[j].y {
				st = append(st, data{pairs[j].y, pairs[j].x + pairs[j].y})
			}
		}
		// 查询需要获取的是大于等于 query.x/query.y 的 最大的 x+y
		// 栈中的数据对应的x一定是大于等于query.x的, 所以只需要找到第一个大于等于query.y的元素即可, 其对应的x+y就是最大的
		// 此时st中存储的数据是按照 total(query[x]) 降序排列的, y是递增的
		p := sort.Search(len(st), func(i int) bool { return st[i].y >= y })
		if p < len(st) {
			ans[i] = st[p].s
		} else {
			ans[i] = -1
		}
	}
	return ans
}
