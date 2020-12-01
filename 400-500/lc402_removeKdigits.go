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
		if len(bs) == 0 {
			return "0"
		}
	}
	return *(*string)(unsafe.Pointer(&bs))
}

func removeKdigitsStack(num string, k int) string {
	var remain = len(num) - k

	if remain == 0 {
		return "0"
	}

	var stack = make([]byte, 0, len(num))

	// 需要注意一下提前退出的情况
	for i := 0; i < len(num); i++ {
		// 从前向后找, 如果栈顶存在且栈顶大于当前元素, 就将栈顶出栈
		for len(stack) != 0 && stack[len(stack)-1] > num[i] && k > 0 {
			stack = stack[:len(stack)-1]
			// 移除了一个元素
			k--
		}
		// 将当前元素入栈
		stack = append(stack, num[i])
	}
	var i int
	// 去掉前导0
	for ; i < len(stack) && stack[i] == '0'; i++ {
	}
	stack = stack[i:]

	if len(stack) == 0 {
		return "0"
	} else if len(stack) > remain {
		// 如果还存在多余的, 就直接取前 remain个
		return string(stack[:remain])
	} else {
		// 这种情况直接返回就好了
		return string(stack)
	}

}

func main() {
	var testCases = []string{
		"1432999",
		"1432119",
		"1432911",
		"10",
		"100",
		"1000",
	}

	for _, n := range testCases {
		fmt.Println(removeKdigits(n, 1))
	}
}
