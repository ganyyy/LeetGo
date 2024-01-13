package main

func addMinimum(word string) int {
	const (
		VALID       = "abc"
		ValidLength = len(VALID)
	)

	var ret int
	var validIdx int

	next := func(idx int) int { return (idx + 1) % ValidLength }

	for i := range word {
		c := word[i]
		for c != VALID[validIdx] {
			ret++
			validIdx = next(validIdx)
		}
		validIdx = next(validIdx)
	}
	if validIdx != 0 {
		// 多了一次 next, 直接跳过.
		ret += ValidLength - validIdx
	}
	return ret
}
