package main

func isIsomorphic(s string, t string) bool {

	// 保留每个字符第一次出现的位置
	var ss, ts [128]int

	for i := 0; i < len(s); i++ {
		if ss[s[i]] == 0 {
			// 如果没出现

			// +1 是为了避免0 是一个合法值
			ss[s[i]] = i + 1
			if ts[t[i]] != 0 {
				return false
			}
			ts[t[i]] = i + 1
		} else {
			// 否则位置必须相等
			if ts[t[i]] != ss[s[i]] {
				return false
			}
		}
	}

	return true
}
