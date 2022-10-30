package main

func letterCasePermutation(s string) (ret []string) {
	src := []byte(s)
	var tmp []byte

	var dfs func(int)

	addByte := func(b byte, idx int) {
		tmp = append(tmp, b)
		dfs(idx + 1)
		tmp = tmp[:len(tmp)-1]
	}

	inRange := func(v, start, end byte) bool {
		return v >= start && v <= end
	}

	dfs = func(idx int) {
		if idx == len(s) {
			ret = append(ret, string(append([]byte(nil), tmp...)))
			return
		}
		b := src[idx]
		addByte(b, idx)
		if inRange(b, 'a', 'z') {
			addByte(b-'a'+'A', idx)
		} else if inRange(b, 'A', 'Z') {
			addByte(b-'A'+'a', idx)
		}
	}

	dfs(0)

	return
}
