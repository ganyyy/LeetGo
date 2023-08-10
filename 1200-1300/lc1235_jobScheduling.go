//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func jobScheduling(startTime, endTime, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	// 按照Job的结束时间进行排序
	sort.Slice(jobs, func(i, j int) bool { return jobs[i][1] < jobs[j][1] })
	fmt.Println(jobs)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// dp[:i]中, 第一个结束时间大于jobs[i]开始时间的位置
		// 所以, k之前的所有位置的结束时间都小于等于jobs[i]的开始时间
		// k是最大的
		k := sort.Search(i, func(j int) bool { return jobs[j][1] > jobs[i-1][0] })
		fmt.Println(jobs[i-1], jobs[k], k)
		// dp[k] 等同于选取 jobs[k-1]时, 产生的利益
		dp[i] = max(dp[i-1], dp[k]+jobs[i-1][2])
	}
	return dp[n]
}
