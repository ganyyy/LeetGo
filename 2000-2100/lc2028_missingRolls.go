package main

func missingRolls(rolls []int, mean int, n int) []int {
	var total = (len(rolls) + n) * mean

	for _, v := range rolls {
		total -= v
	}
	if total < n || total > n*6 {
		return nil
	}

	var base, rem = total / n, total % n
	var getAdd = func() int {
		if rem > 0 {
			rem--
			return 1
		}
		return 0
	}
	var ret = make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = base + getAdd()
	}
	return ret
}
