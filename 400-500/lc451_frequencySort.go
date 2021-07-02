package main

import (
	"bytes"
	"sort"
	"unsafe"
)

func frequencySort(s string) string {
	var cnt [128]int

	for i := range s {
		cnt[s[i]]++
	}

	var tmp = make([][2]int, 0, 26)

	for i, v := range cnt {
		if v == 0 {
			continue
		}
		tmp = append(tmp, [2]int{i, v})
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][1] > tmp[j][1]
	})

	var res = make([]byte, 0, len(s))

	for _, t := range tmp {
		res = append(res, bytes.Repeat([]byte{byte(t[0])}, t[1])...)
	}

	return *(*string)(unsafe.Pointer(&res))
}
