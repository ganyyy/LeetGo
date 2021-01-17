package main

func partitionLabels(S string) []int {
	// 每个字母在S中占据的最远的位置
	var dict = [26]int{}
	var last, pre int
	for i := 0; i < len(S); i++ {
		dict[int(S[i]-'a')] = i + 1
	}
	// 找到同一区间内最远的位置, 就是分界点
	var res []int
	var curLast int
	for i := 0; i < len(S); i++ {
		if curLast = dict[S[i]-'a']; curLast != 0 && last != 0 && last == i {
			res = append(res, last-pre)
			pre = last
			last = curLast
		} else {
			last = max(last, curLast)
		}
	}
	// 末尾的加上去
	res = append(res, last-pre)

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
