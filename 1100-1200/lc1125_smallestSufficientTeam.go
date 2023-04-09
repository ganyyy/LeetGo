package main

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	n, m := len(reqSkills), len(people)
	skillIndex := make(map[string]int)
	for i, skill := range reqSkills {
		skillIndex[skill] = i
	}
	// 每一个状态的初始情况, 默认需要所有的员工填充
	// 然后在DP过程中, 逐步减少所需要的员工的数量
	// dp[i]: 技能组合到达 i 最少需要的员工数
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = m
	}
	dp[0] = 0
	// DP的核心点, 是通过
	// 前一个状态所掌握的技能, 可以用于回溯
	prevSkill := make([]int, 1<<n)
	// 前置员工, 也可以用来回溯
	prevPeople := make([]int, 1<<n)
	for i := 0; i < m; i++ {
		// 迭代每一个工作人员
		curSkill := 0
		// 获取他们对应的技能集合
		for _, s := range people[i] {
			curSkill |= 1 << skillIndex[s]
		}
		// 针对每个员工, 判断他们到指定技能组合的最小员工数
		for prev := 0; prev < (1 << n); prev++ {
			comb := prev | curSkill
			// 这里相当于要找到一个较小值
			if dp[comb] > dp[prev]+1 {
				dp[comb] = dp[prev] + 1
				prevSkill[comb] = prev
				prevPeople[comb] = i
			}
		}
	}
	var res []int
	i := (1 << n) - 1
	// 从末尾状态回溯, 获取最少需要的员工编号
	for i > 0 {
		res = append(res, prevPeople[i])
		i = prevSkill[i]
	}
	return res
}
