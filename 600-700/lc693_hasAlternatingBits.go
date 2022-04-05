package main

func hasAlternatingBits(n int) bool {

	// 如果任意的0和1都不重复相邻, 那么 tmp 就等同于 1111...111
	// 此时tmp&(tmp+1) == 0. 反之则不成立. 因为至少有一位为1
	// 666
	var tmp = n ^ (n >> 1)
	return tmp&(tmp+1) == 0

	// var pre = n % 2

	// for n != 0 {
	//     n /= 2
	//     var cur = n%2
	//     if cur ==  pre {
	//         return false
	//     }
	//     pre = cur
	// }
	// return true
}
