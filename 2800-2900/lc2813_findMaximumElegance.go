package main

import "sort"

func findMaximumElegance(items [][]int, k int) int64 {
	// 首先按照 profit 降序排序
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})
	// 此时的total_profit部分一定是最大的
	// 然后再看distinct_categories部分
	categorySet := map[int]bool{}
	var res, profit int64
	var st []int
	for i, item := range items {
		if i < k {
			// 前k个数据直接加入
			profit += int64(item[0])
			if categorySet[item[1]] {
				// 针对多次出现的类型, 入栈, 用作后续的替换使用
				st = append(st, item[0])
			} else {
				// 首次出现的类型进行标记
				categorySet[item[1]] = true
			}
		} else if !categorySet[item[1]] && len(st) > 0 {
			// 此时需要找到一个可以替换的类型
			// 首先需要保证的是: 这个类型不能是出现过的, 因为已经出现过的类型不会增加任何新的价值
			// 同时, 替换也必须要从前k个已经重复出现的类型中的最小profit进行替换

			// 这一步可能会降低profit
			profit += int64(item[0] - st[len(st)-1])
			st = st[:len(st)-1]
			categorySet[item[1]] = true
		} else if len(st) == 0 {
			// 此时已经没有可以替换的类型了
			break
		}
		// 更新可能的最大值
		res = max(res, profit+int64(len(categorySet)*len(categorySet)))
	}
	return res
}
