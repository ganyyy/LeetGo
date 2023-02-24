package main

func minimumSwap(s1 string, s2 string) int {
	xy, yx := 0, 0
	n := len(s1)
	/*
	   xx yy -> 1次
	   xy yx -> 2次
	*/
	for i := 0; i < n; i++ {
		a, b := s1[i], s2[i]
		if a == 'x' && b == 'y' {
			xy++
		}
		if a == 'y' && b == 'x' {
			yx++
		}
	}
	// 如果为奇数, 那么不管怎么换, 最后都会剩下一组
	if (xy+yx)%2 == 1 {
		return -1
	}
	// 尽可能的匹配 xx yy 这种情况
	// 因为 xy + yx 是一个偶数, 要么都是奇数, 要么都是偶数
	return xy/2 + yx/2 + xy%2 + yx%2
}
