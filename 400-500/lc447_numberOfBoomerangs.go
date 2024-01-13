package main

func numberOfBoomerangs(points [][]int) int {
	// 最简单的方式

	// 每个点距离其他点为不同长度的集合?

	var cc int

	var t = make(map[int]int, len(points))
	for i, c := range points {
		for j, p := range points {
			if j == i {
				continue
			}
			d := dis(c, p)
			// *2的意思是, 两个点可以互换
			cc += t[d] * 2
			t[d]++
		}
		clear(t)
	}
	return cc
}

func dis(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}
