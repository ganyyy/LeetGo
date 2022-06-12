package main

func heightChecker(heights []int) int {
	var cnt [101]int
	for _, v := range heights {
		cnt[v]++
	}
	var ret int
	var idx int
	for i := 1; i <= 100; i++ {
		for cnt[i] > 0 {
			if heights[idx] != i {
				ret++
			}
			cnt[i]--
			idx++
		}
	}
	return ret
}
