package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	lp := len(potions)
	var idx = make([]int, lp)
	for i := range idx {
		idx[i] = i
	}

	sort.Slice(idx, func(i, j int) bool {
		return potions[idx[i]] < potions[idx[j]]
	})

	for i, spell := range spells {
		idx := sort.Search(lp, func(i int) bool {
			return potions[idx[i]]*spell >= int(success)
		})
		spells[i] = lp - idx
	}
	return spells
}
