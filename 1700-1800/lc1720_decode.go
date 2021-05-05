package main

func decode(encoded []int, first int) []int {
	var ret = make([]int, len(encoded)+1)

	ret[0] = first

	for i := 1; i < len(ret); i++ {
		ret[i] = encoded[i-1] ^ ret[i-1]
	}

	return ret
}
