package main

func convertToTitle(columnNumber int) string {
	// 0-A
	// 1-B
	// 25-Z

	var ret []byte
	for columnNumber != 0 {
		columnNumber--
		ret = append(ret, byte(columnNumber%26)+'A')
		columnNumber /= 26
	}

	// 反转一下
	for l, r := 0, len(ret)-1; l < r; l, r = l+1, r-1 {
		ret[l], ret[r] = ret[r], ret[l]
	}

	return string(ret)
}
