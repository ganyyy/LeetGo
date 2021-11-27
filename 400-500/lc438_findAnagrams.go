package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var set [26]int
	var cnt int
	var src [26]bool

	var updateSet func(b byte, v int)

	updateSet = func(b byte, v int) {
		if !src[b-'a'] {
			return
		}
		set[b-'a'] += v
		cnt += v
	}

	var checkSet = func(b byte) bool {
		return set[b-'a'] > 0
	}

	for i := range p {
		src[p[i]-'a'] = true
		updateSet(p[i], 1)
	}
	var pre int
	var ret []int
	for i := range s {
		if checkSet(s[i]) {
			updateSet(s[i], -1)
			if cnt == 0 {
				// 满足条件了, 就整体后移一位
				ret = append(ret, pre)
				updateSet(s[pre], 1)
				pre++
			}
		} else {
			if src[s[i]-'a'] {
				updateSet(s[i], -1) // 先把这个扣了, 因为要选择这个位置
				// 这个元素满足条件, 数量不够了,
				for pre < i {
					updateSet(s[pre], 1)
					//就从pre开始, 查找到第一个有该元素的位置停止
					pre++
					if s[pre-1] == s[i] {
						break
					}
				}
			} else {
				// 这个元素不满足条件, 直接清空之前所有统计的数据
				for ; pre <= i; pre++ {
					updateSet(s[pre], 1)
				}
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd",
		"abc"))
}
