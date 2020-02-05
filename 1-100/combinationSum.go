package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combination(candidates, target)
}

func combination1(candidates []int, target int) [][]int {
	var res [][]int
	if target < candidates[0] || len(candidates) == 0 {
		return res
	}
	for i, v := range candidates {
		if v == target {
			res = append(res, []int{v})
		} else if v > target {
			break
		} else {
			comb := combination(candidates[i:], target-v)
			if len(comb) > 0 {
				for _, c := range comb {
					res = append(res, append([]int{v}, c...))
				}
			}
		}
	}
	return res
}

/**
[7,3,2]
18
*/

func main() {
	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
}
