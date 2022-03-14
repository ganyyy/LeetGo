package main

func countMaxOrSubsets(nums []int) int {
	var max int
	for _, v := range nums {
		max |= v
	}

	var cnt int

	var dfs func(int, int)
	dfs = func(i, cur int) {
		var num = cur | nums[i]
		if num == max {
			cnt++
		}
		for j := i + 1; j < len(nums); j++ {
			dfs(j, num)
		}
	}
	for i := 0; i < len(nums); i++ {
		dfs(i, 0)
	}

	return cnt
}
