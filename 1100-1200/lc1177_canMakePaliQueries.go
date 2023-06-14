package main

func canMakePaliQueries(s string, queries [][]int) []bool {
	// 统计区间内每个字符的个数, 然后判断能否构成回文?
	// 毕竟可以随意重排诶

	// 这里只需要关心区间内各个字符中, 奇数个字符的个数即可
	// 因为如果是偶数的字符, 可以直接添加到两边形成回文
	// 奇数个数的字符可能需要进行特殊的处理. 原则上回文字符串至多拥有1个奇数个数的字符
	// 比如子串是 abcdd , 那么奇数字符数为3, 偶数字符数为1,
	// 如果想要其成为回文, 则至少需要1个(替换a,b,c中任意一个为其他两个之一), 最多需要3个(整体替换)
	// 所以, 对于替换k个数字而言, 最少能使得长度为 2k+1的字符串成为一个回文串?
	// 道理很简单: 不管是偶/奇数种字符, 都至少需要len/2次替换. 那么上限就是 2*k+1
	// ab -> 1
	// abc -> 1
	// abcd -> 2
	// abcde -> 2
	// abcdef -> 3
	// abcdefg -> 3

	// 前缀和计算区间

	n := len(s)
	count := make([]int, n+1)
	for i := 0; i < n; i++ {
		count[i+1] = count[i] ^ (1 << (s[i] - 'a'))
	}
	res := make([]bool, len(queries))
	for i, query := range queries {
		l, r, k := query[0], query[1], query[2]
		var bits int
		x := count[r+1] ^ count[l]
		for x > 0 {
			x &= x - 1
			bits++
		}
		res[i] = bits <= 2*k+1
	}
	return res
}
