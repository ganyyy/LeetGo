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

func main() {

}
