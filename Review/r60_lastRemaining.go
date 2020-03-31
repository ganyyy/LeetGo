package main

// 可以通过队列实现， 时间复杂度为o(m,n)
// 同样的, 存在一个递推公式

func lastRemaining(n int, m int) int {
	// normal
	ln := make([]int, n)
	for i := 0; i < n; i++ {
		ln[i] = i
	}
	c := 0
	m -= 1
	for len(ln) != 1 {
		c += m
		if c >= len(ln) {
			c = c % len(ln)
		}
		ln = append(ln[:c], ln[c+1:]...)
	}
	return ln[0]
}

func lastRemaining2(n, m int) int {
	// 数学方法-反推回去
	// 核心是, 如果剩下一个人, 那么就是这个人
	// 如果剩下N个人, f(N,M)=(f(N−1,M)+M)%N
	// 没去掉一个, 就把剩下的向前移动即可.
	// 然后从前向后推.
	// f(n) 是从0开始数的, 那么f(n-1)就是从
	var p int
	for i := 2; i <= n; i++ {
		p = (p + m) % i
	}
	return p
}

func main() {
	print(lastRemaining2(2, 3))
}
