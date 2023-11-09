package main

import (
	"sort"
)

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

func successfulPairs2(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	var ret = make([]int, len(spells))
	ln := len(potions)
	for i, spell := range spells {
		// 向上取整
		// 假设spell是5, success是7
		// 那么最小满足的potion是 (7+5-1)/5=2
		// sort.SearchInts(potions, 2) 返回的是 potions 中第一个大于等于2的数的索引
		ret[i] = ln - sort.SearchInts(potions, int(success+int64(spell)-1)/spell)
	}
	return ret
}

func successfulPairs3(spells []int, potions []int, success int64) []int {
	var ret = make([]int, len(spells))
	for i := range ret {
		ret[i] = i
	}

	const (
		Shift = 32
		Mask  = (1 << Shift) - 1
	)

	// spell  正序
	sort.Slice(ret, func(i, j int) bool { return spells[ret[i]] <= spells[ret[j]] })
	// potions倒叙
	sort.Sort(sort.Reverse(sort.IntSlice(potions)))

	var potionIdx int
	for _, order := range ret {
		spell := spells[order&Mask]
		for potionIdx < len(potions) && int64(potions[potionIdx]*spell) >= success {
			potionIdx++
		}
		ret[order&Mask] |= potionIdx << Shift
	}

	for i := range ret {
		ret[i] >>= Shift
	}
	return ret

}
