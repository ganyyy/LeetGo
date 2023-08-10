package main

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	const DAY = 24 * 60

	// 一天就1440分钟, 超过这个时间点肯定有重复的
	if len(timePoints) > DAY {
		return 0
	}
	sort.Strings(timePoints)

	var parse = func(str string) int {
		var hour, minute = str[:2], str[3:]
		var h, _ = strconv.Atoi(hour)
		var m, _ = strconv.Atoi(minute)
		return h*60 + m
	}

	var a, b = parse(timePoints[0]), parse(timePoints[1])
	var ret = b - a

	for i := 2; i < len(timePoints); i++ {
		a, b = b, parse(timePoints[i])
		ret = min(ret, b-a)
	}

	ret = min(ret, DAY+parse(timePoints[0])-b)
	return ret
}
