package main

import (
	"fmt"
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	// 数位DP!

	// 直接CV

	digNum := make([]int, 0, len(digits))
	for _, s := range digits {
		v, _ := strconv.Atoi(s)
		digNum = append(digNum, v)
	}

	num := strconv.Itoa(n)
	ln := len(num)
	dp := make([]int, ln)

	// 初始化
	for i := range dp {
		dp[i] = -1
	}

	var layer []byte

	// idx: 当前DP的位置
	// isLimit: 当前位置是否存在上限
	// isNum: 是否选取了数字(空串不算数字)
	var dfs func(idx int, isLimit, isNum bool) int

	dfs = func(idx int, isLimit, isNum bool) int {
		layer = append(layer, ' ')
		defer func() { layer = layer[:len(layer)-1] }()
		fmt.Println(string(layer), idx, isLimit, isNum)
		if idx == ln {
			// 1. 到达了数字结尾, 判定是否是一个合法的数字
			if isNum {
				return 1
			}
			return 0
		}
		if !isLimit && isNum && dp[idx] >= 0 {
			// 2. 存在DP的结果, 直接返回
			// 只能是没有限制, 并且已经选取了数字, 和已经DP过的情况
			return dp[idx]
		}
		var res int
		if !isNum {
			// !isNum表示没有选取任何数字, 所以需要看下一位
			// 同样的, 因为当前位数不满足上限, 所以等同于是没有上限
			res = dfs(idx+1, false, false)
			fmt.Println(string(layer), idx, res)
		}
		// 当前位置的可达上限
		up := 9
		if isLimit {
			up = int(num[idx] - '0')
		}

		for _, dn := range digNum {
			if dn <= up {
				res += dfs(idx+1, dn == up && isLimit, true)
				continue
			}
			// 非递减序列, 不满足条件可以直接跳出
			break
		}
		if !isLimit && isNum {
			// 保留每一位上的最大值
			// 为啥呢? 因为每一位只会被DP一次
			fmt.Println(string(layer), "set dp", idx, res)
			dp[idx] = res
		}
		return res
	}

	return dfs(0, true, false)
}
