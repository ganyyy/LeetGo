package main

func numRabbits(answers []int) int {
	var cnt = make(map[int]int, 10)

	for _, v := range answers {
		cnt[v]++
	}

	var res int
	res += cnt[0]
	delete(cnt, 0)
	for k, v := range cnt {
		res += (v + k) / (k + 1) * (k + 1)
	}

	return res
}

func numRabbits2(answers []int) int {
	var cnt = make(map[int]int)

	// 一次遍历版本
	var res int
	var old int
	var find bool
	for _, v := range answers {
		if v == 0 {
			res++
			continue
		}
		old, find = cnt[v]
		if !find {
			old += v
		}
		old += 1
		if old >= v+1 {
			res += v + 1
			old = 0
		}
		cnt[v] = old
	}

	return res
}
