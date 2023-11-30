package main

func firstCompleteIndex(arr []int, mat [][]int) int {
	if len(arr) == 0 {
		return -1
	}
	m := len(mat)
	if m == 0 {
		return -1
	}
	n := len(mat[0])
	if n == 0 {
		return -1
	}

	if m == 1 || n == 1 {
		// 单行或者单列 情况
		return 0
	}

	var matIndex = make([][2]int32, m*n)
	var count = make([]int, m+n)

	for r, row := range mat {
		for c, val := range row {
			val--
			vr, vc := val/n, val%n
			matIndex[vr*n+vc] = [2]int32{int32(r), int32(c)}
		}
	}

	for i, num := range arr {
		num--
		r, c := num/n, num%n
		idx := r*n + c
		vr, vc := int(matIndex[idx][0]), int(matIndex[idx][1])
		r, c = vr, vc+m
		count[r]++
		count[c]++
		if count[r] == n || count[c] == m {
			return i
		}
	}
	return -1
}
