package main

func countVowelPermutation(n int) int {
	// 印象中, 有个多规则DP来着?
	var pa = 1
	var pe = 1
	var pi = 1
	var po = 1
	var pu = 1

	const MOD = 1e9 + 7

	var get = func(v int) int {
		return v % MOD
	}

	for i := 2; i <= n; i++ {
		pa, pe, pi, po, pu =
			get(pe+pi+pu),
			get(pa+pi),
			get(pe+po),
			get(pi),
			get(pi+po)
	}

	return get(pa + pe + pi + po + pu)
}
