package main

import (
	"strconv"
	"unsafe"
)

func monotoneIncreasingDigits(N int) int {
	// 987654321 -> 899999999
	// 981234567 -> 899999999
	// 991234567 -> 899999999

	var bs = []byte(strconv.Itoa(N))

	// 转成字符串, 好处理一点

	// 情况1 单调递增数字 -> 就是本身

	var i int
	for i = 0; i < len(bs)-1; i++ {
		if bs[i] > bs[i+1] {
			break
		}
	}
	if i == len(bs)-1 {
		return N
	}

	// 情况2 中间存在拐点 -> 从拐点前一位开始-1, 如果减去之后 小于上一位数字, 上一位-1, 这一位置成9
	// 减到头为值
	// , 往后全填9即可

	// 首先把末尾置成9
	for j := i + 1; j < len(bs); j++ {
		bs[j] = '9'
	}

	// i就是拐点, 从后向前赋值
	for ; i >= 0; i-- {
		if i > 0 && bs[i]-1 < bs[i-1] {
			bs[i] = '9'
		} else {
			bs[i] -= 1
			break
		}
	}

	var ret, _ = strconv.Atoi(toString(bs))

	// 情况2 中间存在拐点 -> 从拐点前一位开始-1, 如果减去之后 小于上一位数字, 上一位-1, 这一位置成9
	// 减到头为值
	// , 往后全填9即可

	return ret

}

func monotoneIncreasingDigitsBigLao(N int) int {
	var i = 1
	var res = N

	for i <= res/10 {
		var n = res / i % 100 // 每次取两位出来
		i *= 10               // 加一个进位
		if n/10 > n%10 {      // 如果高位大于低位
			res = res/i*i - 1 // 置成 XX999的格式, 不停的减一位
		}
	}
	return res
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func toBytes(s string) []byte {
	var x = (*[2]uintptr)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&[3]uintptr{x[0], x[1], x[1]}))
}
