package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combination(candidates, target)
}

func combination(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 || target < candidates[0] {
		return res
	}
	for i, v := range candidates {
		if i != 0 && v == candidates[i-1] {
			continue
		}
		if v == target {
			res = append(res, []int{v})
			continue
		}
		if v > target {
			break
		}
		combs := combination(candidates[i+1:], target-v)
		if len(combs) > 0 {
			tmp := []int{v}
			for _, comb := range combs {
				res = append(res, append(tmp, comb...))
			}
		}
	}
	return res
}

/**
[10,1,2,7,6,1,5]
8
*/
func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
