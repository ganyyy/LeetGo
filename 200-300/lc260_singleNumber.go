package main

import (
	"fmt"
	"math/bits"
)

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

func singleNumbers(nums []int) []int {
	var odd int
	for _, i := range nums {
		odd ^= i
	}

	// 取最后一位1: 不行, 因为可能会发生溢出(?)
	flag := odd & (-odd)

	// 记录两个值其中的一个
	var res int
	for _, i := range nums {
		if i&flag != 0 {
			res ^= i
		}
	}
	// 输出结果
	return []int{res, odd ^ res}
}

func singleNumber3(nums []int) []int {
	// 还是得3
	var num int

	for _, v := range nums {
		num ^= v
	}

	fmt.Println(num)

	var index = bits.Len(uint(num)) - 1

	fmt.Println(index)

	var a int
	for _, v := range nums {
		if v&(1<<index) != 0 {
			a ^= v
		}
	}

	fmt.Println(a)

	return []int{a, a ^ num}

}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("Len:%v, val:%b\n", bits.Len(uint(i)), i)
	}
}
