//go:build ignore

package main

import "strings"

func areSentencesSimilar(sentence1, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")
	i, n := 0, len(words1)
	j, m := 0, len(words2)
	// 左向右
	for i < n && i < m && words1[i] == words2[i] {
		i++
	}
	// 右向左
	for j < n-i && j < m-i && words1[n-j-1] == words2[m-j-1] {
		j++
	}
	// 最长相似的字符串长度加一块等同于最短的那个句子
	return i+j == min(n, m)
}
