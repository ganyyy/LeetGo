package main

func decode(encoded []int) []int {
	var ret = make([]int, len(encoded)+1)

	// 需要前3个确认初始的4个组合

	// 一共N个数, 从[1..N]且中间不会重复

	// 求从1->N总的异或结果
	var p int
	for i := 1; i <= len(encoded)+1; i++ {
		p ^= i
	}

	// 从后向前消除, 只剩下最终剩下开头的唯一一个数字
	for i := len(encoded) - 1; i >= 0; i -= 2 {
		p ^= encoded[i]
	}

	// 从头构造结果
	ret[0] = p
	for i := 1; i < len(ret); i++ {
		ret[i] = ret[i-1] ^ encoded[i-1]
	}

	return ret
}
