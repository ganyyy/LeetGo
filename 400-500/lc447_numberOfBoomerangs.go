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
			t[dis(c, p)]++
		}
		for _, cnt := range t {
			if cnt >= 2 {
				cc += cnt * (cnt - 1)
			}
		}
		for k := range t {
			delete(t, k)
		}
	}
	return cc
}

func dis(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}
