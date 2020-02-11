package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	addVal(nil, nums, &res)
	return res
}

func addVal(pre, next []int, res *[][]int) {
	if len(next) == 1 {
		*res = append(*res, append(pre, next[0]))
		return
	}
	for i := 0; i < len(next); i++ {
		if i == 0 || next[0] != next[i] {
			next[0], next[i] = next[i], next[0]
			addVal(append(pre, next[0]), next[1:], res)
			next[i], next[0] = next[0], next[i]
		}
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 1}))
}
