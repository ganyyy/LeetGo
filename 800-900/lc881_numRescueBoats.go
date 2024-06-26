package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	// 开头+结尾进行组合, 保证数量最少?

	var left, right = 0, len(people) - 1

	var ret int

	for left <= right {
		// 相等的时候, 二倍体重是否超过limit无所谓
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		ret++
	}
	return ret
}
