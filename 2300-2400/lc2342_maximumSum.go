package main

func maximumSum(nums []int) int {
	// nums中数字的范围在 [1, 1e9] 之间, 所以每个位上的累加和最大是 9*9=81
	var maxSum [82]int
	var ret = -1
	for _, num := range nums {
		var total int
		for n := num; n != 0; n /= 10 {
			total += n % 10
		}
		old := maxSum[total]
		if old != 0 {
			ret = max(ret, old+num)
		}
		if num > old {
			maxSum[total] = num
		}
	}
	return ret
}
