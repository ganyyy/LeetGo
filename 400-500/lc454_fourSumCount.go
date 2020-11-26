package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	// A+B 的集合

	// 将两组数据进行平分, 降低时间复杂度

	var abMap = make(map[int]int, len(A)*len(B))

	for _, va := range A {
		for _, vb := range B {
			abMap[va+vb]++
		}
	}

	var res int
	for _, vc := range C {
		for _, vd := range D {
			res += abMap[-(vc + vd)]
		}
	}

	return res
}
