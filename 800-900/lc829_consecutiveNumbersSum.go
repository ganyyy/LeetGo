package main

func consecutiveNumbersSum(N int) int {
	// 1个数时，必然有一个数可构成N
	// 2个数若要构成N，第2个数与第1个数差为1，N减掉这个1能整除2则能由商与商+1构成N
	// 3个数若要构成N，第2个数与第1个数差为1，第3个数与第1个数的差为2，N减掉1再减掉2能整除3则能由商、商+1与商+2构成N
	// 依次内推，当商即第1个数小于等于0时结束
	res := 0
	for i := 1; N > 0; N, i = N-i, i+1 {
		if N%i == 0 {
			res += 1
		}
	}
	return res
}

func consecutiveNumbersSum2(n int) int {
	var ret int

	// 查找指定的连续的数字的个数
	// 9 = 9 = 4+5 = 2+3+4
	// 连续1个, i = 1, n = 9, 9 % 1 == 0 -> 存在一个解. 9
	// 连续2个, i = 2, n = 8, 8 % 2 == 0 -> 存在一个解. 4, 4+1(4和4+1中间差了1个数)
	// 连续3个, i = 3, n = 6, 6 % 3 == 0 -> 存在一个解. 2, 2+1, 2+2(2和2+2中间差了2个数, 差值为3)
	// 死记硬背吧
	for i := 1; n > 0; n, i = n-i, i+1 {
		if n%i == 0 {
			ret += 1
		}
	}

	return ret
}

func main() {

}
