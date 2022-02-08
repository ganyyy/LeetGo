package main

func countKDifference(nums []int, k int) int {
	var cnt [101]int

	var ret int
	for _, v := range nums {
		if v+k <= 100 {
			ret += cnt[v+k]
		}
		if v-k > 0 {
			ret += cnt[v-k]
		}
		cnt[v]++
	}
	return ret
}
