package main

import "sort"

func sumOfPower(nums []int) (ans int) {
	const mod = 1_000_000_007
	sort.Ints(nums)
	s := 0
	// 只看贡献
	// 假设有 a,b,c,d,e 五个数, 从小到大依次递增
	// 假设当前最大值是 d
	// a..d 中间数有2个, d^2 * a * 2^2
	// b..d 中间数有1个, d^2 * b * 2^1
	// c..d 中间数有0个, d^2 * c * 2^0
	// d    自身的计算 , d^2 * d
	// 累加: d^3 + d^2*(a*2^2 + b*2^1 + c*2^0)
	// 令 a*2^2 + b*2^1 + c*2^0 = s,
	// 得出 d^3 + d^2 * s = d^2*(d + s)

	// 那么当最大值为e的时候, a*2^3 + b*2^2 + c*2^1 + d*2^0
	// = 2*(a*2^2 + b*2^1 + c*2^0) + d = 2*s+d

	// s是滚动累加的!
	for _, x := range nums { // x 作为最大值
		ans = (ans + x*x%mod*(x+s)) % mod // 中间模一次防止溢出
		s = (s*2 + x) % mod               // 递推计算下一个 s
	}
	return
}
