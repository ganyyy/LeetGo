package main

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}

	var ret = make([][]int, m)

	for i := range ret {
		ret[i] = original[i*n : (i+1)*n]
	}

	return ret
}
