package main

func numberOfWeeks(milestones []int) int64 {
	// 耗时最长工作所需周数
	longest := 0
	rest := int64(0)
	for _, count := range milestones {
		longest = max(longest, count)
		rest += int64(count)
	}
	// 其余工作共计所需周数
	rest -= int64(longest)
	if int64(longest) > rest+1 {
		// 此时无法完成所耗时最长的工作
		return rest*2 + 1
	} else {
		// 此时可以完成所有工作
		return int64(longest) + rest
	}
}
