package main

import "math"

func findRestaurant(list1 []string, list2 []string) []string {

	var set1 = make(map[string]int)

	for i, v := range list1 {
		set1[v] = i
	}

	var set2 = make(map[int][]string)
	var minIdx = math.MaxInt32
	for i, v := range list2 {
		if idx, ok := set1[v]; !ok {
			continue
		} else {
			if i+idx <= minIdx {
				minIdx = i + idx
				set2[i+idx] = append(set2[i+idx], v)
			}
		}
	}
	return set2[minIdx]
}
