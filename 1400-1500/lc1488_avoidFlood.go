package main

import "sort"

func avoidFlood(rains []int) []int {
	// 所有为0的天数
	var dryDay []int

	ln := len(rains)
	ans := make([]int, ln)
	rainToDay := map[int]int{}

	for i := range ans {
		ans[i] = 1
	}

	for day, rain := range rains {
		if rain == 0 {
			// 这一天可以抽水
			dryDay = append(dryDay, day)
		} else {
			ans[day] = -1
			if day, ok := rainToDay[rain]; ok {
				// 如果这一天之前已经下过雨了, 说明该抽水了
				// 这里相当于查找 dryDay 中是否存在 >= day 的元素. 因为是不可能相同的, 所以只能是大于
				idx := sort.SearchInts(dryDay, day)
				if idx >= len(dryDay) {
					// 在上一次下雨和今天之间没有多余可用的抽水天数
					return nil
				}
				dry := dryDay[idx]
				if dry < len(ans) {
					ans[dry] = rain
				}
				copy(dryDay[idx:], dryDay[idx+1:])
				dryDay = dryDay[:len(dryDay)-1]
			}
			// 标记下雨
			rainToDay[rain] = day
		}
	}
	return ans
}
