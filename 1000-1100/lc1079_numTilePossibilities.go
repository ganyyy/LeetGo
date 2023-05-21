package main

const mx = 8

var c [mx][mx]int

func init() {
	for i := 0; i < mx; i++ {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			/*
			   × .
			   × √
			*/
			c[i][j] = c[i-1][j-1] + c[i-1][j] // 预处理组合数
		}
	}
}

func numTilePossibilities(tiles string) (ans int) {
	counts := map[rune]int{}
	for _, ch := range tiles {
		counts[ch]++ // 统计每个字母的出现次数
	}
	// 字母需要去重
	n, m := len(tiles), len(counts)
	// f[i][j]: 前i种字符构建长度为j的序列的个数(n >= j >= i >= 1)
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	f[0][0] = 1 // 构造空序列的方案数
	i := 1
	for _, cnt := range counts { // 枚举第 i 种字母
		f[i] = make([]int, n+1)
		for j := 0; j <= n; j++ { // 枚举序列长度 j
			for k := 0; k <= j && k <= cnt; k++ { // 枚举第 i 种字母选了 k 个
				// 长度为j, 插入k个, 剩余的数量就是 j-k
				// f[i-1][j-k]: 前i-1种字符构建长度为j-k的序列的个数
				// i-1: 不选用这个字符时所有的方案数
				// j-k: 选用这个字符时, 可以在j-k个位置中任意插入k个值
				f[i][j] += f[i-1][j-k] * c[j][k] // 选用这个字符时, 可以在j个位置中任意插入k个值
			}
		}
		i++
	}
	for _, x := range f[m][1:] {
		ans += x
	}
	return
}
