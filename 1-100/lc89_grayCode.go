package main

func grayCode(n int) []int {
	var ret = make([]int, 0, 1<<n)

	for i := 0; i < (1 << n); i++ {
		// 真: 记公式 i ^ (i/2)
		ret = append(ret, i^(i>>1))
	}
	return ret
}
