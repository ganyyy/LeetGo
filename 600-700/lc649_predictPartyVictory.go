package main

func predictPartyVictory(senate string) string {
	// 就是看剩下谁呗

	// 统计人数?
	// 记录两种阵营出现的索引位置
	var n = len(senate)
	var rStack = make([]int, 0, n)
	var dStack = make([]int, 0, n)

	for i := range senate {
		if senate[i] == 'R' {
			rStack = append(rStack, i)
		} else {
			dStack = append(dStack, i)
		}
	}

	// 依次出队, 每次消掉索引大的, 并将小的放到身后
	for len(rStack) != 0 && len(dStack) != 0 {
		if rStack[0] < dStack[0] {
			rStack = append(rStack, rStack[0]+n)
		} else {
			dStack = append(dStack, dStack[0]+n)
		}
		rStack = rStack[1:]
		dStack = dStack[1:]
	}

	if len(rStack) == 0 {
		return "Dire"
	}
	return "Radiant"
}
