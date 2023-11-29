package main

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if word1 == word2 {
		return true
	}

	if len(word1) != len(word2) {
		return false
	}

	var cnt1, cnt2 [26]int
	var alpha1, alpha2 int
	for i := range word1 {
		ia, ib := word1[i]-'a', word2[i]-'a'
		cnt1[ia]++
		cnt2[ib]++
		alpha1 |= 1 << ia
		alpha2 |= 1 << ib
	}

	if alpha1 != alpha2 {
		// 出现的字符种类不一致
		return false
	}

	sort.Ints(cnt1[:])
	sort.Ints(cnt2[:])

	return cnt1 == cnt2
}
