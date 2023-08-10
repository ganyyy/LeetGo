package main

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	// 前缀和
	preSumArr := make([]int, n+1)
	for i, num := range nums {
		preSumArr[i+1] = preSumArr[i] + num
	}
	ans := n + 1

	// 双端队列
	var q []int
	for i, curSum := range preSumArr {
		// 对头到当前位置的差就是区间和, 如果满足条件了就出队
		for len(q) > 0 && curSum-preSumArr[q[0]] >= k {
			ans = min(ans, i-q[0])
			q = q[1:]
		}
		// https://github.com/Shellbye/Shellbye.github.io/issues/41
		// 如果队尾和当前位置之间存在负数?
		// 如果都是正数, 那么一定满足 curSum > preSumArr[q[len(q)-1]]
		// 所以, 只能是 从 q[len(q)-1] - i 之间存在负数
		// 无需计算 [q[len(q)-1], i]之间的值, 因为中间存在的负数一定会导致列变长
		for len(q) > 0 && preSumArr[q[len(q)-1]] >= curSum {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if ans < n+1 {
		return ans
	}
	return -1
}
