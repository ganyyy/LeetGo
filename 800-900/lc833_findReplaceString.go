package main

import (
	"sort"
	"strings"
)

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	var sb strings.Builder
	var order = make([]int, len(indices))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool { return indices[order[i]] < indices[order[j]] })

	var writeIdx int
	for _, i := range order {
		index := indices[i]
		source := sources[i]
		target := targets[i]
		if writeIdx < index {
			sb.WriteString(s[writeIdx:index])
			writeIdx = index
		}
		if strings.HasPrefix(s[index:], source) {
			sb.WriteString(target)
			writeIdx += len(source)
		}
	}
	sb.WriteString(s[writeIdx:])
	return sb.String()
}
