package main

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {
	var idx = sort.Search(len(letters), func(i int) bool {
		return letters[i] > target
	})
	// fmt.Println(idx)
	if idx < len(letters) {
		return letters[idx]
	}
	return letters[0]
}
