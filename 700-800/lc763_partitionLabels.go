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

func partitionLabels2(s string) (ret []int) {
	if len(s) < 2 {
		return []int{len(s)}
	}
	// 获取区间最远的位置
	var cnt [26]int
	idx := func(b byte) int {
		return int(b - 'a')
	}
	for i := range s {
		cnt[idx(s[i])] = i
	}
	// fmt.Println(cnt)
	var last int
	var cur int
	for i := range s {
		c := cnt[idx(s[i])]
		if c == last {
			// abc
			ret = append(ret, 1)
			last = i + 1
		} else if cur == i && i != 0 {
			// aba
			ret = append(ret, i-last+1)
			last = i + 1
		}
		cur = max(cur, c)
	}
	if last != len(s) {
		ret = append(ret, len(s)-last)
	}
	return
}

func partitionLabels3(s string) (partition []int) {
	// 这个解法好评
	var lastPos [26]int
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}
