package main

import (
	"fmt"
	"unsafe"
)

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	var bs = []byte(num)

	// 1432999 -> 1432
	// 1432119 -> 1119
	// 1432911 -> 1211

	// 从左到右. 找第一个比后面大的字符, 然后删掉它

	var j int
	for i := 0; i < k; i++ {
		// 此时, bs[j] < bs[j-1], bs[j-1]当前最大, 删掉
		for j = 1; j < len(bs) && bs[j] >= bs[j-1]; j++ {
		}
		bs = append(bs[:j-1], bs[j:]...)
		for j = 0; j < len(bs) && bs[j] == '0'; j++ {
		}
		bs = bs[j:]
	}
	return *(*string)(unsafe.Pointer(&bs))
}

func main() {
	var testCases = []string{
		"1432999",
		"1432119",
		"1432911",
	}

	for _, n := range testCases {
		fmt.Println(removeKdigits(n, 3))
	}
}
