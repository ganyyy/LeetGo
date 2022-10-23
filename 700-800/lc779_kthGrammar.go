package main

func kthGrammar(n, k int) (ans int) {

	// 1
	// 10
	// 1001
	// 1001 0110
	// 1001 0110 0110 1001

	// 如果 k > 2^(n-2), 就需要进行反转, 否则就

	// 核心点是找到对应关系
	// 观察上边的序列, 可以发现: 对于任意一行, 前半部分和后半部分的对应位置的值是相反的
	// 从二进制方向考虑, 就相当于去掉高位的1, 有几个1就反转几次

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
