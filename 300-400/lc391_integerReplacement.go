package main

func integerReplacement(n int) int {
	var ret int
	var cur = n
	for cur != 1 {
		if cur&3 == 3 && cur != 3 {
			// 在末尾不是3的情况下
			// 末尾是0xXXX11时, +1可以获得两个0, 0xXXX00 -> 0xXXX0 -> 0xXXX
			// 如果-1, 那么相当于变成 0xXXX10 -> 0xXXX1 -> 0xXXX0 -> 0xXXX
			cur++
		} else if cur&1 == 1 {
			// 其他情况下, 直接-1即可
			cur--
		} else {
			// 末尾是0, 直接右移
			cur >>= 1
		}
		// 每次操作+1
		ret++
	}

	return ret
}
