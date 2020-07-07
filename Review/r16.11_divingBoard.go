package main

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int(nil)
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	res := make([]int, k+1)
	for i := 0; i <= k; i++ {
		res[i] = shorter*(k-i) + longer*i
	}
	return res
}
