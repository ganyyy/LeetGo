//go:build ignore

package main

import (
	"fmt"
	"unsafe"
)

func addStrings(num1 string, num2 string) string {
	// 1是短的一方
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	// 保留一个进位
	res := make([]byte, len(num2)+1)
	sub := len(num2) - len(num1)
	// 遍历字符串1, 两两相加
	var add int
	for i := len(num1) - 1; i >= 0; i-- {
		total := num1[i] - '0' + num2[i+sub] - '0' + byte(add)
		if total >= 10 {
			total -= 10
			add = 1
		} else {
			add = 0
		}
		res[i+sub+1] = byte(total) + '0'
	}
	// 遍历较长的字符串剩下的
	for i := sub - 1; i >= 0; i-- {
		if t := num2[i] - '0' + byte(add); t >= 10 {
			add = 1
			res[i+1] = t - 10 + '0'
		} else {
			res[i+1] = t + '0'
			add = 0
		}
	}
	if add == 1 {
		res[0] = '1'
	} else {
		res = res[1:]
	}
	return string(res)
}

func addStrings2(num1, num2 string) string {
	sb := make([]byte, max(len(num1), len(num2))+1)
	carry, i, j := 0, len(num1)-1, len(num2)-1
	var cur int
	// 倒叙加到正序中, 在进行反转
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
			j--
		}
		sb[cur] = byte(carry%10) + '0'
		carry /= 10
		cur++
	}
	// 反转一下
	for head, tail := 0, len(sb)-1; head < tail; head, tail = head+1, tail-1 {
		sb[head], sb[tail] = sb[tail], sb[head]
	}

	if sb[0] == 0 {
		sb = sb[1:]
	}
	return string(sb)
}

func addStrings3(num1 string, num2 string) string {
	// nums1 为较小的一方
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	var bs1, bs2 = stringToBytes(num1), stringToBytes(num2)
	var ret = make([]byte, 1+len(bs2))

	var as, bs, rs = len(bs1) - 1, len(bs2) - 1, len(ret) - 1
	var add int

	var addToRet = func(sum int) {
		var b byte
		if sum >= 10 {
			b = byte(sum-10) + '0'
			add = 1
		} else {
			add = 0
			b = byte(sum) + '0'
		}
		ret[rs] = b
		rs--
	}
	for as >= 0 {
		addToRet(int(bs1[as]-'0') + add + int(bs2[bs]-'0'))
		as--
		bs--
	}

	for bs >= 0 {
		addToRet(add + int(bs2[bs]-'0'))
		bs--
	}
	if add == 0 {
		ret = ret[1:]
	} else {
		ret[rs] = '1'
	}
	return bytesToString(ret)
}

func stringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			s string
			l int
		}{
			s: s, l: len(s),
		}),
	)
}

func bytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(addStrings2("19", "1"))
}
