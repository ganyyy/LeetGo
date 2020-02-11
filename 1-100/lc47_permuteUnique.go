package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	var res [][]int
	addVal(nil, nums, &res)
	return res
}

func addVal(pre, next []int, res *[][]int) {
	if len(next) == 1 {
		*res = append(*res, append(pre, next[0]))
		return
	}
	mp := make(map[int]struct{})
	for i := 0; i < len(next); i++ {
		if _, ok := mp[next[i]]; ok {
			continue
		}
		next[0], next[i] = next[i], next[0]
		addVal(append(pre, next[0]), next[1:], res)
		next[i], next[0] = next[0], next[i]
		mp[next[i]] = struct{}{}
	}
}

func main() {
	fmt.Println(len(permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1})))
}
