package main

import (
	"sort"
	"strings"
)

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	var subs sort.StringSlice
	// cnt: (+1, )-1
	// left:
	cnt, left := 0, 0
	for i, ch := range s {
		if ch == '1' {
			cnt++
		} else if cnt--; cnt == 0 {
			// 1放前边, 0放后边
			// 将其变成(+[left+1:i]+)
			// s[left] = 1
			// s[i] = 0
			//fmt.Println(left+1, i, s[left+1:i])
			subs = append(subs, "1"+makeLargestSpecial(s[left+1:i])+"0")
			// 找到一组匹配的括号后, i+1一定是(
			left = i + 1
		}
	}
	sort.Sort(sort.Reverse(subs))
	return strings.Join(subs, "")
}
