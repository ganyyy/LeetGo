package main

func distributeCandies(candyType []int) int {
	var set = make(map[int]struct{}, len(candyType)>>1)

	for _, v := range candyType {
		set[v] = struct{}{}
	}

	if len(set) > len(candyType)>>1 {
		return len(candyType) >> 1
	}
	return len(set)
}
