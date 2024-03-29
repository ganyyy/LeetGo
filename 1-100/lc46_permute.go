package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	addVal46(nil, nums, &res)
	return res
}

func addVal46(pre, next []int, res *[][]int) {
	if len(next) == 1 {
		*res = append(*res, append(pre, next[0]))
		return
	}
	for i := 0; i < len(next); i++ {
		if i == 0 || next[0] != next[i] {
			next[0], next[i] = next[i], next[0]
			addVal46(append(pre, next[0]), next[1:], res)
			next[i], next[0] = next[0], next[i]
		}
	}
}

func permute2(nums []int) [][]int {
	var ret [][]int

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			var tmp = make([]int, len(nums))
			copy(tmp, nums)
			ret = append(ret, tmp)
			return
		}
		for j := i; j < len(nums); j++ {
			nums[i], nums[j] = nums[j], nums[i]
			dfs(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	dfs(0)
	return ret
}

func permute3(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var ret [][]int
	var dfs func(candidates, buf []int)
	dfs = func(candidates, buf []int) {
		if len(candidates) == 0 {
			ret = append(ret, append([]int(nil), buf...))
			return
		}
		var start = candidates[0]
		for idx, val := range candidates {
			candidates[idx] = start
			dfs(candidates[1:], append(buf, val))
			candidates[idx] = val
		}
	}
	dfs(nums, make([]int, 0, len(nums)))
	return ret
}

func main() {
	fmt.Println(permute([]int{1, 2, 1}))
}
