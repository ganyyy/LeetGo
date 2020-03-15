package main

func singleNumber(nums []int) []int {
	// 首先全部异或一边, 得到的结果就是两个不同的数的异或的结果
	var a, b, xor int
	for _, v := range nums {
		xor ^= v
	}
	// 找到xor第一个非0位, 代表着两位数在这一位上一个为0, 一个为1
	var bit = 1
	for xor&1 == 0 {
		xor >>= 1
		bit <<= 1
	}
	// 以这个为标准, 在进行一次异或操作, 可以将结果分为两组, 返回即可
	for _, v := range nums {
		if v&bit == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}

func main() {

}
