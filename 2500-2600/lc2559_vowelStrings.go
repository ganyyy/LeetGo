package main

func vowelStrings(words []string, queries [][]int) []int {
	var preSum = make([]int, len(words)+1)

	for i, word := range words {
		cur := preSum[i]
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			cur++
		}
		preSum[i+1] = cur
	}

	ret := make([]int, 0, len(queries))

	for _, query := range queries {
		start, end := query[0], query[1]
		ret = append(ret, preSum[end+1]-preSum[start])
	}

	return ret

}

func isVowel(c uint8) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}
