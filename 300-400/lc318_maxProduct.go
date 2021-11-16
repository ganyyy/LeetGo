package main

func maxProduct(words []string) int {
	var set = make([]int, len(words))

	for idx, v := range words {
		var val int
		for i := range v {
			val |= 1 << int(v[i]-'a')
		}
		set[idx] = val
	}

	var ret int

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if set[i]&set[j] != 0 {
				continue
			}
			ret = max(ret, len(words[i])*len(words[j]))
		}
	}
	return ret
}
