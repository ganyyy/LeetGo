package main

func oddCells(m int, n int, indices [][]int) int {
	var row = make([]int, n)
	var col = make([]int, m)

	for _, indice := range indices {
		row[indice[0]]++
		col[indice[1]]++
	}

	var ret int
	for _, v1 := range row {
		for _, v2 := range col {
			ret += (v1 + v2) % 2
		}
	}
	return ret
}
