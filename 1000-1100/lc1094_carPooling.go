package main

func carPooling(trips [][]int, capacity int) bool {
	// 差分数组
	var road [1001]int
	var mx int
	for _, trip := range trips {
		p := trip[0]
		road[trip[1]] += p
		road[trip[2]] -= p
		mx = max(mx, trip[2])
	}
	var sum int
	for _, num := range road[:min(1000, mx)+1] {
		sum += num
		if sum > capacity {
			return false
		}
	}
	return true
}
