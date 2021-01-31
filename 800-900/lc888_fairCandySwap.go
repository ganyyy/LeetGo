package main

func fairCandySwap(A []int, B []int) []int {
	var sumA, sumB int
	for _, v := range A {
		sumA += v
	}
	var mb = make(map[int]bool, len(B))
	for _, v := range B {
		sumB += v
		mb[v] = true
	}

	var mid = (sumA + sumB) >> 1

	var t, v int
	for _, v = range A {
		if t = v - sumA + mid; mb[t] {
			break
		}
	}

	return []int{v, t}
}
