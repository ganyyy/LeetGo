package main

func maximumWealth(accounts [][]int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var sum = func(nums []int) int {
		var ret int
		for _, v := range nums {
			ret += v
		}
		return ret
	}

	var ret int
	for _, nums := range accounts {
		ret = max(sum(nums), ret)
	}
	return ret
}
