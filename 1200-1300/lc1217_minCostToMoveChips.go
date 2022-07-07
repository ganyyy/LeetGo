package main

func minCostToMoveChips(position []int) int {
	// 奇变偶不变

	var odd, even int
	for _, v := range position {
		if v&1 == 1 {
			odd++
		} else {
			even++
		}
	}
	if odd > even {
		return even
	} else {
		return odd
	}
}
