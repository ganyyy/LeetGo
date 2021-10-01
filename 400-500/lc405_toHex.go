package main

import (
	"bytes"
	"fmt"
)

var TAB = [...]byte{
	'0',
	'1',
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
	'a',
	'b',
	'c',
	'd',
	'e',
	'f',
}

func toHex(num int) string {
	const MAX = 0xffffffff + 1

	if num < 0 {
		num += MAX
	}
	var ret bytes.Buffer
	for num != 0 {
		ret.WriteByte(TAB[num%16])
		num /= 16
	}
	var buffer = ret.Bytes()
	for i, j := 0, ret.Len()-1; i < j; i, j = i+1, j-1 {
		buffer[i], buffer[j] = buffer[j], buffer[i]
	}

	return ret.String()
}

func toHex2(num int) string {
	const MAX = 0xffffffff + 1

	if num == 0 {
		return "0"
	}

	if num < 0 {
		num += MAX
	}

	var ret = make([]byte, 8)
	var i = 8

	for num != 0 {
		i--
		ret[i] = TAB[num&15]
		num /= 16
	}
	return string(ret[i:])
}

func main() {
	fmt.Println(toHex(-1))
	fmt.Println(toHex(100))
	fmt.Println(toHex(16))
}
