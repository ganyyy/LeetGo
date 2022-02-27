package main

import "math/bits"

func maximumRequests(n int, requests [][]int) (ans int) {
	// 简单而言, 就是枚举所有可能的转移方式, 然后基于状态进行各种枚举
next:
	for mask := 0; mask < 1<<len(requests); mask++ {
		// cnt是要计算的请求的数量.
		// mask是需要计算的路径, 为1表示需要计算
		cnt := bits.OnesCount(uint(mask))
		if cnt <= ans {
			continue
		}
		// 最终的n个房子的状态. 因为要成环, 所以最终这个值应该是0
		delta := make([]int, n)
		for i, r := range requests {
			// 这是一个被选中的转移方案, 进行房间转移
			if mask>>i&1 == 1 {
				delta[r[0]]++
				delta[r[1]]--
			}
		}
		// 如果本次分配可以成功, 那么就记录一下当前成环的个数
		for _, d := range delta {
			if d != 0 {
				// 这不是一个合法的状态, 看下一个
				continue next
			}
		}
		ans = cnt
	}
	return
}

func maximumRequests2(n int, requests [][]int) int {
	// 枚举所有可能的选项
	var ret int
	var delta = make([]int, n)
next:
	for mask := 0; mask < 1<<len(requests); mask++ {
		var cnt = bits.OnesCount(uint(mask))
		if cnt <= ret {
			continue
		}
		for i := range delta {
			delta[i] = 0
		}
		// 枚举状态
		for i, r := range requests {
			// 这是一个被选中的转移方案, 进行房间转移
			if mask>>i&1 == 1 {
				delta[r[0]]++
				delta[r[1]]--
			}
		}

		for _, v := range delta {
			if v != 0 {
				continue next
			}
		}
		ret = cnt
	}

	return ret
}

func maximumRequests3(n int, requests [][]int) int {
	// 变换后，每个节点的出度和入度是相等的
	// 一共2^16种可能性， 判断每一种可能性
	// 可能性判断，用degree来记录，
	// dfs
	res := 0
	picked := 0
	zeroCnt := n
	degrees := make([]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(idx int)
	dfs = func(idx int) {
		if zeroCnt == n {
			res = max(res, picked)
		}
		if idx == len(requests) {
			// fmt.Println(zeroCnt, picked)
			return
		}
		// 如果出入相等，必选
		if requests[idx][0] == requests[idx][1] {
			picked++
			dfs(idx + 1)
			picked--
			return
		}

		// 不取这个request
		dfs(idx + 1)

		// 取这个request
		if degrees[requests[idx][0]] == 0 {
			zeroCnt--
		}
		degrees[requests[idx][0]]--
		if degrees[requests[idx][0]] == 0 {
			zeroCnt++
		}
		if degrees[requests[idx][1]] == 0 {
			zeroCnt--
		}
		degrees[requests[idx][1]]++
		if degrees[requests[idx][1]] == 0 {
			zeroCnt++
		}
		picked++
		dfs(idx + 1)
		// 恢复现场
		picked--

		if degrees[requests[idx][0]] == 0 {
			zeroCnt--
		}
		degrees[requests[idx][0]]++
		if degrees[requests[idx][0]] == 0 {
			zeroCnt++
		}
		if degrees[requests[idx][1]] == 0 {
			zeroCnt--
		}
		degrees[requests[idx][1]]--
		if degrees[requests[idx][1]] == 0 {
			zeroCnt++
		}
	}
	dfs(0)
	return res

}

func main() {
	maximumRequests2(5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}})
}
