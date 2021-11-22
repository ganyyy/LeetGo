package main

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	var set [26]int
	var cnt [26]int
	var diffCnt int

	for i := range s {
		var a, b = s[i], goal[i]
		set[a-'a']++
		set[b-'a']--
		cnt[a-'a']++
		if a == b {
			continue
		}
		// 保证相差的个数不大于3
		diffCnt++
		if diffCnt > 2 {
			return false
		}
	}

	if s == goal {
		// 相等时, 对比是否有重复的字符串
		for _, c := range cnt {
			if c >= 2 {
				return true
			}
		}
		return false
	}

	// 不等时, 等同于二者的差集为空
	for _, v := range set {
		if v != 0 {
			return false
		}
	}

	// 相等的字符串, 如果存在两个以上相同的字符, 也算作是可以交换

	return true
}
