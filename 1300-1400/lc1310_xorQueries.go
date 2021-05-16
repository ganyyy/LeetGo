package main

func xorQueries(arr []int, queries [][]int) []int {
	// 有点类似于前缀和
	for i := 1; i < len(arr); i++ {
		arr[i] ^= arr[i-1]
	}

	var ret = make([]int, len(queries))

	for i, q := range queries {
		if q[0] == 0 {
			ret[i] = arr[q[1]]
		} else {
			// 利用了异或的基本原理, 即a^b^c ^ a^b = c
			ret[i] = arr[q[1]] ^ arr[q[0]-1]
		}
	}

	return ret
}
