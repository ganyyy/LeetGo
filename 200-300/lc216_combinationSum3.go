package main

func combinationSum3(k int, n int) [][]int {
	var res [][]int

	if n/k >= 9 && k != 1 {
		return res
	}

	var helper func(i, target int)
	var tmp = make([]int, 0, k)
	helper = func(i, target int) {
		if len(tmp) > k || i > 10 {
			return
		}
		if target == 0 && len(tmp) == k {
			t := make([]int, k)
			copy(t, tmp)
			res = append(res, t)
			return
		}
		for c := i; c <= target; c++ {
			tmp = append(tmp, c)
			helper(c+1, target-c)
			tmp = tmp[:len(tmp)-1]
		}
	}

	helper(1, n)
	return res
}
