package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	// s2中 锁定一个长度大小为 s1 的窗口, 保证窗口内的字符和s1的组成相同
	if len(s1) > len(s2) {
		return false
	}

	// 保存出现过的字符 和对应的次数
	var set = make(map[int]int, 26)
	for i := range s1 {
		set[int(s1[i]-'a')]++
	}

	var l, r int
	for ; r < len(s2); r++ {
		var c = int(s2[r] - 'a')
		if set[c] > 0 {
			set[c]--
			if r-l+1 == len(s1) {
				return true
			}
			continue
		}

		// 不能无脑递增, 需要存在一个跳出/替换的条件
		for l <= r {
			var lc = int(s2[l] - 'a')
			l++
			if v, ok := set[lc]; ok {
				if lc != c {
					set[lc] = v + 1
				} else {
					break
				}
			}
		}
	}

	return false
}

func checkInclusion2(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		// 如果比0大, 说明不在s1中
		// 这种方式就不需要处理替换的问题了
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab",
		"eidboaoo"))
}
