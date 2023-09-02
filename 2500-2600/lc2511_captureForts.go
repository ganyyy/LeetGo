package main

func captureForts(forts []int) int {
	// 连续0的最多个数?
	var mx int

	// 去除前边的0和后边的0
	l, r := 0, len(forts)-1
	var start = l
	for i := l; i <= r; i++ {
		if forts[i] == 0 {
			continue
		}
		// 是否是有效边界? 只能是[1,-1]才有效嘛?
		if forts[start] == -forts[i] {
			mx = max(mx, i-start-1)
		}
		start = i
	}
	return mx
}
