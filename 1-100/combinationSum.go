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

func combinationSum3(candidates []int, target int) (ans [][]int) {
	var comb []int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		// 核心点是: 不会回退!
		// 很有意思的思路: 从小往大的发散, 保证了不会出现重复的组合

		// 直接跳过: 不选择当前数
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

/**
[7,3,2]
18
*/

func main() {
	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
}
