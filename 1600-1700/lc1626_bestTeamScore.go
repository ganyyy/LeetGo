package main

import "sort"

func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)
	people := make([][]int, n)
	for i := range scores {
		people[i] = []int{scores[i], ages[i]}
	}
	sort.Slice(people, func(i, j int) bool {
		// 先按照得分
		if people[i][0] < people[j][0] {
			return true
		} else if people[i][0] > people[j][0] {
			return false
		}
		// 再按照年龄
		return people[i][1] < people[j][1]
	})
	// 前i位的最大值
	dp := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		// [i].score > [j].score
		for j := 0; j < i; j++ {
			// 如果 [j].age < [i].age
			if people[j][1] <= people[i][1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		// 加上[i].score
		dp[i] += people[i][0]
		res = max(res, dp[i])
	}
	return res
}
