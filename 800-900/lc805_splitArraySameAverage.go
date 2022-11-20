package main

func sum(nums []int) int {
	var ret int
	for _, v := range nums {
		ret += v
	}
	return ret
}

func filterSum(nums []int, bit int) int {
	var ret int
	for i, v := range nums {
		if bit>>i&1 != 0 {
			ret += v
		}
	}
	return ret
}

func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	allSum := sum(nums)
	for i := range nums {
		nums[i] = nums[i]*n - allSum
	}

	var leftSet = make(map[int]struct{})

	// 折半
	m := n / 2
	for i := 1; i < 1<<m; i++ {
		tot := filterSum(nums[:m], i)
		if tot == 0 {
			// fmt.Println("reture left")
			return true
		}
		leftSet[tot] = struct{}{}
	}

	// 右半部分
	rightSum := sum(nums[m:])
	for i := 1; i <= (1<<(n-m))-1; i++ {
		tot := filterSum(nums[m:], i)
		if tot == 0 {
			// fmt.Println("reture right zero")
			return true
		}
		// 单调递增的数据, 如果相同了, 就意味着right是一个空的集合
		if tot == rightSum {
			continue
		}
		if _, ok := leftSet[-tot]; ok {
			// fmt.Println("reture right equal")
			return true
		}
	}
	return false
}
