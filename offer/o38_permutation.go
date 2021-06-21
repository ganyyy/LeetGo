package main

func permutation(s string) []string {
	var dfs func(i int)

	var ret []string

	var tmp = []byte(s)

	dfs = func(i int) {
		if i == len(s)-1 {
			ret = append(ret, string(tmp))
			return
		}
		var set = make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if !set[tmp[j]] {
				set[tmp[j]] = true
			} else {
				continue
			}
			tmp[i], tmp[j] = tmp[j], tmp[i]
			dfs(i + 1)
			tmp[j], tmp[i] = tmp[i], tmp[j]
		}
	}

	dfs(0)

	return ret
}
