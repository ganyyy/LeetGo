package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	var res [][]int
	var helper func(n, i int, before []int)

	helper = func(n, i int, before []int) {
		if n == 0 {
			t := make([]int, len(before))
			copy(t, before)
			res = append(res, t)
			return
		}
		for j := i; j <= len(nums)-n; j++ {
			before[len(before)-n] = nums[j]
			helper(n-1, j+1, before)
		}
	}

	for i := 0; i <= len(nums); i++ {
		tmp := make([]int, i)
		helper(i, 0, tmp)
	}
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
