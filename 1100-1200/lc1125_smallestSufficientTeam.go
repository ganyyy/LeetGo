package main

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	n, m := len(reqSkills), len(people)
	skillIndex := make(map[string]int)
	for i, skill := range reqSkills {
		skillIndex[skill] = i
	}
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = m
	}
	dp[0] = 0
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
		for prev := 0; prev < (1 << n); prev++ {
			comb := prev | curSkill
			// 这里就是类似于DP, 最起码需要保证不能低于前置技能
			// > 1 说明结合这名员工的技能后, 拥有了新的进展
			if dp[comb] > dp[prev]+1 {
				dp[comb] = dp[prev] + 1
				prevSkill[comb] = prev
				prevPeople[comb] = i
			}
		}
	}
	res := []int{}
	i := (1 << n) - 1
	for i > 0 {
		res = append(res, prevPeople[i])
		i = prevSkill[i]
	}
	return res
}
