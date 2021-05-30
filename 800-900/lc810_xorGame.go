package main

func xorGame(nums []int) bool {
	// 数学推论, 没那个思维还是别想了. 反正也搞不起来
	if len(nums)&1 == 0 {
		return true
	}
	var sum int

	for _, v := range nums {
		sum ^= v
	}

	return sum == 0
}
