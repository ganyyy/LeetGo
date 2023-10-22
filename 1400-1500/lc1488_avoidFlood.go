package main

import "sort"

func avoidFlood(rains []int) []int {
	// 所有为0的天数
	var buf [16]int
	var droughtDay = buf[:0]

	length := len(rains)
	ans := make([]int, length)
	rainToDay := map[int]int{}

	for i := range ans {
		ans[i] = 1
	}

	// rains[i] == 0 表示第 i 天没有下雨
	// rains[i] > 0 表示第 i 天下雨了, 并且下在了 rains[i] 这个湖泊上
	for day, lake := range rains {
		if lake == 0 {
			// 这一天可以抽水
			droughtDay = append(droughtDay, day)
		} else {
			ans[day] = -1
			if lastDay, ok := rainToDay[lake]; ok {
				// 如果这一天之前这个湖已经下过雨了, 说明该抽水了
				// 这里相当于查找 droughtDay 中是否存在 >= lastDay 的元素. 因为是不可能相同的, 所以只能是大于
				// 相当于从[lastDay, day)中找到一个最小的元素, 然后抽水
				idx := sort.SearchInts(droughtDay, lastDay)
				if idx >= len(droughtDay) {
					// 在上一次下雨和今天之间没有多余可用的抽水天数
					return nil
				}
				// 这里应该不可能发生dryDay >= len(ans)的情况
				// 因为dryDay本质上是一个下标, 而ans的长度就是下标的最大值
				ans[droughtDay[idx]] = lake
				copy(droughtDay[idx:], droughtDay[idx+1:])
				droughtDay = droughtDay[:len(droughtDay)-1]
			}
			// 标记下雨
			rainToDay[lake] = day
		}
	}
	return ans
}
