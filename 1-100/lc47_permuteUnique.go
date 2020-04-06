package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	addVal(nums, 0, &res)
	return res
}

func addVal(nums []int, left int, res *[][]int) {
	if left == len(nums) {
		*res = append(*res, append([]int(nil), nums...))
		return
	}
	m := make(map[int]struct{})
	for i := left; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		nums[left], nums[i] = nums[i], nums[left]
		addVal(nums, left+1, res)
		nums[left], nums[i] = nums[i], nums[left]
		m[nums[i]] = struct{}{}
	}
}

func main() {
	fmt.Println(len(permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1})))
}
