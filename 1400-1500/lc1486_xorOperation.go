package main

func xorOperation(n int, start int) int {
	// 凡人做法
	var ret = start
	for i := 1; i < n; i++ {
		ret = ret ^ (start + 2*i)
	}

	return ret
}

func sumXor(x int) int {
	switch x % 4 {
	case 0:
		return x
	case 1:
		return 1
	case 2:
		return x + 1
	default:
		return 0
	}
}

func xorOperationGod(n, start int) (ans int) {
	// 神仙做法
	s, e := start>>1, n&start&1
	ret := sumXor(s-1) ^ sumXor(s+n-1)
	return ret<<1 | e
}
