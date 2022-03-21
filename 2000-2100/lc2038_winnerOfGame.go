package main

func winnerOfGame(colors string) bool {
	// 连续数量超过3的个数吗?

	var aIdx, bIdx int

	var aCnt, bCnt int

	for i := range colors {
		switch colors[i] {
		case 'A':
			aIdx++
			if aIdx > 2 {
				aCnt++
			}
			bIdx = 0
		case 'B':
			bIdx++
			if bIdx > 2 {
				bCnt++
			}
			aIdx = 0
		}
	}
	return aCnt > bCnt
}
