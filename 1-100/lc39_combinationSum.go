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
	if len(candidates) == 0 {
		return nil
	}
	var dfs func(idx, target int, buf []int)

	dfs = func(idx, target int, buf []int) {
		if idx == len(candidates) || target < 0 {
			return
		}
		if target == 0 {
			// 添加到结果集
			ans = append(ans, append([]int(nil), buf...))
			return
		}
		// 跳过当前元素
		dfs(idx+1, target, buf)
		if rem := target - candidates[idx]; rem >= 0 {
			// 选择当前元素
			dfs(idx, rem, append(buf, candidates[idx]))
		}
	}
	dfs(0, target, make([]int, 0, len(candidates)))
	return
}

/**
[7,3,2]
18
*/

func main() {
	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
}
