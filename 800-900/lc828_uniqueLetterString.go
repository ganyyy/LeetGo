package main

func uniqueLetterString(s string) (ans int) {
	// last[c][1]:上次出现的位置
	// last[c][0]:当前出现的位置
	sum, last := 0, [26][2]int{}
	for i := range last {
		// 初始情况下, 全是-1
		last[i] = [2]int{-1, -1}
	}
	for i, c := range s {
		c -= 'A'
		// (i-[0])-([0]-[1])
		// 好神奇啊
		sum += i - last[c][0]*2 + last[c][1]
		ans += sum
		last[c][1] = last[c][0]
		last[c][0] = i
	}
	return
}
