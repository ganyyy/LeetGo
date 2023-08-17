package main

import (
	"fmt"
	"sort"
)

func combinationSum5(candidates []int, target int) [][]int {
	// 先排序
	sort.Ints(candidates)
	var res [][]int
	var tmp []int

	var helper func(index, target int)

	helper = func(index, target int) {
		// 如果目标值为0, 加一个答案
		if target == 0 {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}
		// 全选完了 或者 最小值比目标值要大, 直接返回
		if index >= len(candidates) || candidates[index] > target {
			return
		}
		// 从剩余候选中依次找,
		for i := index; i < len(candidates); {
			if v := candidates[i]; v <= target {
				tmp = append(tmp, v)
				helper(i+1, target-v)
				tmp = tmp[:len(tmp)-1]
			} else {
				break
			}
			for i = i + 1; i < len(candidates) && candidates[i-1] == candidates[i]; i++ {
			}
		}
	}

	helper(0, target)

	return res
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
