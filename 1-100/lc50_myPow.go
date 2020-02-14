package main

func myPow(x float64, n int) float64 {
	var res float64 = 1
	// 不管怎么算, 不停的/2 一定会得到1
	for i := n; i != 0; i /= 2 {
		// 如果i是奇数, 必须多乘一个x(包括了最后为1的情况),
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

func main() {

}
