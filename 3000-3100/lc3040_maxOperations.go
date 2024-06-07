package main

func maxOperations(nums []int) int {
	n := len(nums)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(start, end, pre int) int
	dfs = func(start, end, pre int) int {
		if end-start < 1 {
			return 0
		}

		if ret := cache[start][end]; ret != -1 {
			return ret
		}

		var n int
		if nums[start]+nums[start+1] == pre {
			n = dfs(start+2, end, pre) + 1
		}
		if nums[end-1]+nums[end] == pre {
			n = max(n, dfs(start, end-2, pre)+1)
		}
		if nums[start]+nums[end] == pre {
			n = max(n, dfs(start+1, end-1, pre)+1)
		}
		cache[start][end] = n
		return n
	}

	end := len(nums) - 1

	return max(dfs(2, end, nums[0]+nums[1]), dfs(0, end-2, nums[end]+nums[end-1]), dfs(1, end-1, nums[0]+nums[end])) + 1
}
