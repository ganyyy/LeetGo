package main

func maxScore2(cardPoints []int, k int) int {
	var total int
	for _, point := range cardPoints {
		total += point
	}
	ln := len(cardPoints)
	if ln <= k {
		return total
	}
	var current int
	for _, point := range cardPoints[:ln-k] {
		current += point
	}
	var ret = total - current
	for removeIdx, point := range cardPoints[ln-k:] {
		current += point - cardPoints[removeIdx]
		ret = max(ret, total-current)
	}
	return ret
}
