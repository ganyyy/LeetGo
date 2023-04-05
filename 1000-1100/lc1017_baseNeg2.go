package main

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	var buffer []byte

	// n = 3 = 0011
	// 0011 -> 1
	// 0011 --位移--> 0001 -> 1 --取反--> -1 -> 1111
	// 1111 -> 1(算数位移补符号位)
	// 1111 --位移--> 1111 ->-1 --取反--> 1 -> 0001
	// 0001 -> 1
	// 0001 --位移--> 0000 -> 0 --取反--> 0 -> 0

	// fmt.Printf("%032b\n", -1 >> 1)

	for n != 0 {
		// +1: 0001
		// -1: 1111
		// 末尾的奇偶性保持不变
		// fmt.Printf("%032b\n", n)
		buffer = append(buffer, '0'+byte(n&1))
		// 每一次位移, 都意味着一次符号的反转.
		n = -(n >> 1)
	}
	for l, r := 0, len(buffer)-1; l < r; l, r = l+1, r-1 {
		buffer[l], buffer[r] = buffer[r], buffer[l]
	}
	return string(buffer)
}
