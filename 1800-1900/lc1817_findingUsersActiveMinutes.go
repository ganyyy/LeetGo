package main

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	// [id] -> [minute]
	mp := map[int]map[int]struct{}{}
	for _, p := range logs {
		id, t := p[0], p[1]
		if mp[id] == nil {
			mp[id] = map[int]struct{}{}
		}
		mp[id][t] = struct{}{}
	}
	ans := make([]int, k+1)
	for _, m := range mp {
		ans[len(m)]++
	}
	return ans[1:]
}
