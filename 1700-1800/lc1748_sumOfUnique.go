package main

func sumOfUnique(nums []int) int {
	var count = make(map[int]int, len(nums))

	for _, v := range nums {
		count[v]++
	}

	var ret int

	for v, c := range count {
		if c != 1 {
			continue
		}
		ret += v
	}
	return ret
}
