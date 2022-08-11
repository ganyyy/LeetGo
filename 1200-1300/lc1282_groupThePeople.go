package main

func groupThePeople(groupSizes []int) [][]int {
	var group = make(map[int][]int)

	var total int
	for i, g := range groupSizes {
		group[g] = append(group[g], i)
	}

	var ret [][]int
	for g, mem := range group {
		total += len(mem) / g
	}
	ret = make([][]int, 0, total)
	for g, mem := range group {
		for i := 0; i < len(mem); i += g {
			ret = append(ret, mem[i:i+g])
		}
	}
	return ret
}
