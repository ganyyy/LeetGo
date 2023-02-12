package main

import "sort"

func alertNames(keyName, keyTime []string) (ans []string) {
	timeMap := map[string][]int{}
	for i, name := range keyName {
		t := keyTime[i]
		hour := int(t[0]-'0')*10 + int(t[1]-'0')
		minute := int(t[3]-'0')*10 + int(t[4]-'0')
		timeMap[name] = append(timeMap[name], hour*60+minute)
	}
	// 有优化空间. 一次排序, 然后再迭代.
	for name, times := range timeMap {
		sort.Ints(times)
		for i, t := range times[2:] {
			// 最多只需要添加一次, 这里对比的就是 (times[i+2]-times[i])
			if t-times[i] <= 60 {
				ans = append(ans, name)
				break
			}
		}
	}
	sort.Strings(ans)
	return
}
