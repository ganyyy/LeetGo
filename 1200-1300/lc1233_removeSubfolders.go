package main

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) (ans []string) {
	// 按照字典序排序, 天然符合父文件夹和子文件夹之间的顺序
	sort.Strings(folder)
	ans = append(ans, folder[0])
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			// 没有共同前缀, 那么就肯定不是同一个父文件夹 /a/b /a/c
			// 有共同签注, 但是多了一点 /a/b /a/bc
			ans = append(ans, f)
		}
	}
	return
}
