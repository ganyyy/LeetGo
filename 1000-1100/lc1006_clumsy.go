package main

func clumsy(N int) int {
	// 4个一组
	// 除了第一组, 其他的首个数字都是负数

	var res int
	var t int
	for i := N; i >= 4; i -= 4 {
		t = i
		if i != N {
			t = -i
		}
		res += t*(i-1)/(i-2) + i - 3
	}
	// 末尾的几个
	var r = N % 4
	t = r
	if r != N {
		t = -t
	}
	switch r {
	case 1:
		res += t
	case 2:
		res += t * (r - 1)
	case 3:
		res += t * (r - 1) / (r - 2)
	}

	return res
}

// 这题怎么成了找规律了...
func clumsy2() {
	// N > 3时
	// N+1 == N*(N-1)/(N-2)
	// 没这个必要把....
}

func main() {
	println(clumsy(9))
}
