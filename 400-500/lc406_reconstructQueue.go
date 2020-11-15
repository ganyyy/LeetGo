package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	// k == 0 的在最前边, 且按照从小到大的顺序排序

	// 先按照 h降序, k升序排列
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		}
		return people[i][1] < people[j][1]
	})
	// 此时, 对于 h不同的人, 可以保证 后边没有比当前身高要高的
	// 对于 h 相同的人, 可以保证相对顺序不会发生变化
	// 遍历排序后的 people.
	for i, p := range people {
		// 对于已经拍好顺序的数组而言, p的身高是最低的
		// 此时需要给当前p找到合适的位置

		// i 表示不小于当前身高的人的个数
		// 如果i > k, 说明p所处的位置应该在i前边, 因为后边不会存在比 p更高的人
		// 所以可以确定 p 所处的位置就是在已经拍好的数组的 索引k 上
		if k := p[1]; k < i {
			copy(people[k+1:i+1], people[k:i+1])
			people[k] = p
		}
	}

	return people
}
