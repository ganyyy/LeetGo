package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	lp := len(potions)

	// 借用idx对potions/spells进行排序
	var idx = make([]int, lp)
	for i := range idx {
		idx[i] = i
	}

	// 从小到大排序
	sort.Slice(idx, func(i, j int) bool {
		return potions[idx[i]] < potions[idx[j]]
	})

	for i, spell := range spells {
		// 找到满足条件的最小的idx
		idx := sort.Search(lp, func(i int) bool {
			// 如果返回了true, 说明满足条件, 尝试减小i; 否则尝试增大i
			return potions[idx[i]]*spell >= int(success)
		})
		spells[i] = lp - idx
	}
	return spells
}
