package main

func kthGrammar(n, k int) (ans int) {

	// 1
	// 10
	// 1001
	// 1001 0110
	// 1001 0110 0110 1001

	// 如果 k > 2^(n-2), 就需要进行反转, 否则就

	if n == 1 {
		return 0
	}
	if k > (1 << (n - 2)) {
		return 1 ^ kthGrammar(n-1, k-(1<<(n-2)))
	}
	return kthGrammar(n-1, k)

	// 进一步优化, 起始就是求 k-1 中 1的个数...
	// 为啥呢?  因为反转就是不停的去掉 最右边的 1 !

	// return bits.OnesCount(uint(k-1)) & 1
}
