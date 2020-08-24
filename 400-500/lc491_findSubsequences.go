package main

func findSubsequences(nums []int) [][]int {

	var res [][]int
	var empty = struct{}{}

	var dfsHelper func(sub []int, cur int)

	dfsHelper = func(sub []int, cur int) {
		set := map[int]struct{}{}
		if len(sub) > 1 {
			// copy 一份, 防止在另外的地方修改
			tmp := make([]int, len(sub))
			copy(tmp, sub)
			res = append(res, tmp)
		}
		for i := cur; i < len(nums); i++ {
			v := nums[i]
			if _, ok := set[v]; ok {
				continue
			}
			if ln := len(sub); ln == 0 || v >= sub[ln-1] {
				set[v] = empty
				sub = append(sub, v)
				dfsHelper(sub, i+1)
				sub = sub[:ln]
			}
		}
	}
	dfsHelper(nil, 0)
	return res
}
