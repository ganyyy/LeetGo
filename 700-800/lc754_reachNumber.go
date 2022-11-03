package main

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	k := 0
	for target > 0 {
		k++
		target -= k
	}

	// delta := target-Σ(1,k)

	if target%2 == 0 {
		// 如果delta是偶数, 比如是 -2, 那么只需要将 +1 变成 -1即可
		return k
	}
	// 如果是delta奇数,
	// delta+(k+1) 是奇数, 那么就需要走两步才能变成偶数 -(k+1)+(k+2) = +1. 此时k是奇数
	// 如果delta+(k+1) 是偶数, 那么只需要走一步就能变成偶数. 此时k是偶数
	return k + 1 + k&1
}
