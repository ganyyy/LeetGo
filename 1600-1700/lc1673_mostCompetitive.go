package main

func mostCompetitive(nums []int, k int) []int {
	res := make([]int, 0, len(nums))
	for i, x := range nums {
		// 先上来总是可以保证ret中至少包含k个递增数, 到末尾时将不再进行替换
		for len(res) > 0 && len(nums)-i+len(res) > k && res[len(res)-1] > x {
			res = res[:len(res)-1]
		}
		res = append(res, x)
	}
	// 此时ret的前k个数就是字典序的最小值
	return res[:k]
}
