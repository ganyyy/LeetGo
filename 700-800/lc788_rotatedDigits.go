//go:build ignore

package main

import "strconv"

var check = [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

func rotatedDigits(n int) int {
	digits := strconv.Itoa(n)
	// 最高5位, 是否边界有2个状态; 是否出现2569有2个状态
	memo := [5][2][2]int{}
	for i := 0; i < 5; i++ {
		memo[i] = [2][2]int{{-1, -1}, {-1, -1}}
	}
	var dfs func(int, bool, bool) int
	// bound表示是否要选取当前位置的边界值
	// diff表示是否存在了2569
	dfs = func(pos int, bound, diff bool) (res int) {
		if pos == len(digits) {
			return bool2int(diff)
		}
		ptr := &memo[pos][bool2int(bound)][bool2int(diff)]
		if *ptr != -1 {
			return *ptr
		}
		lim := 9
		if bound {
			// 12345, pos == 3,
			// 如果前边的是 123XX, 那么bound==true,  所以当前位选取的上限是4
			// 如果前边的是 122XX, 那么bound==false, 所以当前位的选取上限是9
			lim = int(digits[pos] - '0')
		}
		for i := 0; i <= lim; i++ {
			if check[i] != -1 {
				res += dfs(pos+1, bound && i == int(digits[pos]-'0'), diff || check[i] == 1)
			}
		}
		*ptr = res
		return
	}
	return dfs(0, true, false)
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
