package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	var sp1, sp2 = strings.Split(s1, " "), strings.Split(s2, " ")

	var t1 = make(map[string]int, len(sp1))
	var t2 = make(map[string]int, len(sp2))

	for _, s := range sp1 {
		t1[s]++
	}
	for _, s := range sp2 {
		t2[s]++
	}

	var ret = make([]string, 0, (len(s1)+len(s2))/2)
	for s, c := range t1 {
		if c > 1 {
			continue
		}
		if _, ok := t2[s]; ok {
			continue
		}
		ret = append(ret, s)
	}

	for s, c := range t2 {
		if c > 1 {
			continue
		}
		if _, ok := t1[s]; ok {
			continue
		}
		ret = append(ret, s)
	}
	return ret
}
