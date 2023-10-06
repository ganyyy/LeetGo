package main

func myPow(x float64, n int) float64 {
	var res float64 = 1
	// 不管怎么算, 不停的/2 一定会得到1
	for i := n; i != 0; i /= 2 {
		// 如果i是奇数, 必须多乘一个x(包括了最后为1的情况),
		// 这个x不一定是原来的x!!!
		if i&1 != 0 {
			res *= x
		}
		// 直接翻倍就行
		x *= x
	}
	if n < 0 {
		return 1.0 / res
	} else {
		return res
	}
}

func myPow2(x float64, n int) float64 {
	// 二分?

	// 这么理解
	// 2^10 = (((2^2)^2)^2) * 2 * 2
	/*
		2^10 = 2^5 * 2^5 = 4^5
		4^5 = 4^2 * 4^2 * 4 = 16^2 * 4 => res *= 4
		16^2 = 256^1 => res *= 256 = 1024
	*/
	// 这里通过二分找快速的找到最接近的 2的指数次幂
	// 中间的每一个单数都相当于 多乘以一次
	// 不管怎么计算, 最终 n 都会到1, 此时 的res就是最终的结果

	var res float64 = 1
	var i = n
	for n != 0 {
		if n&1 != 0 {
			res *= x
		}
		n /= 2
		x *= x
	}

	if i < 0 {
		return 1.0 / res
	}
	return res
}

func myPow3(x float64, n int) float64 {
	var ret float64 = 1
	var cur = x
	for n := n; n != 0; n /= 2 {
		if n&1 == 1 {
			ret *= cur
		}
		cur *= cur
	}
	if n < 0 {
		ret = 1 / ret
	}
	return ret
}

func main() {

}
