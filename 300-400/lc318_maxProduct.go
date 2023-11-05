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

var buf [1001]int

func maxProduct2(words []string) int {
	var pack = func(s string) int {
		var ret int
		for i := range s {
			ret |= 1 << int(s[i]-'a')
		}
		return ret
	}

	strEncode := buf[:len(words)]

	var ret int
	for i, word := range words {
		ei := pack(word)
		strEncode[i] = ei
		for j, preWord := range words[:i] {
			if ei&strEncode[j] != 0 {
				continue
			}
			ret = max(ret, len(word)*len(preWord))
		}
	}
	return ret
}
