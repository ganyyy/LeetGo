package main

func goodSubsetofBinaryMatrix(grid [][]int) []int {
	// 当只有一行时, 则每一列的值都应该是0 (1/2)
	// 当只有两行时, 每一列的的和不应该超过1, 即二者相与 = 0 (2/2)
	// 当有三行时, 等同于两行 (3/2 = 2/2), 所以两行肯定就够了
	// 当有超过四行时, 至少需要6列...
	maskToIdx := map[int]int{}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		maskToIdx[mask] = i
	}

	for x, i := range maskToIdx {
		for y, j := range maskToIdx {
			if x&y == 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return nil
}
