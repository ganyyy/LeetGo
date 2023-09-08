package main

import (
	"bytes"
	"unsafe"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	size := len(s)
	res := make([]byte, 0, size)
	sub := (numRows - 1) * 2
	for i := 0; i < numRows; i++ {
		j := i
		mul := 1
		for j < size {
			res = append(res, s[j])
			if i != 0 && i != numRows-1 {
				// 不是开头和结尾的情况, 求定点的和在减去当前值即为中间值
				val := sub*(2*mul-1) - j
				if val < size {
					res = append(res, s[val])
				}
			}
			j += sub
			mul += 1
		}
	}
	return string(res)
}

func convert2(s string, numRows int) string {
	// 还是有点进步的蛤~
	if numRows == 1 {
		return s
	}
	var next = [2]int{0, (numRows - 1) * 2}
	var ret = bytes.NewBuffer(nil)
	ret.Grow(len(s))
	for i := 0; i < numRows; i++ {
		var cur = i
		var cnt int
		for cur < len(s) {
			ret.WriteByte(s[cur])
			cnt++
			if n := next[cnt&1]; n == 0 {
				cnt++
			}
			cur += next[cnt&1]
		}
		next[0], next[1] = next[0]+2, next[1]-2
	}

	return ret.String()
}

func convert3(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	// 一个N字形为一组
	/*
	   首先可以确认的是: 最大差值是 (numRows-1)*2. (从A->G)
	   next[0]表示这一组中第一个额外添加的, next[1]表示的是这一组中的第二个额外添加的
	   开头和结尾是有一些特殊的. 因为他们在一组中只额外存在一次. 为了规律, 分别将其放置在前后
	   A     G    [6, 0]
	   B   F H    [4, 2]
	   C E   I    [2, 4]
	   D     J    [0, 6]
	*/
	var next = [2]int{0, (numRows - 1) * 2}
	var ret = make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		var cur = i
		var cnt int
		for cur < len(s) {
			ret = append(ret, s[cur])
			cnt++
			if n := next[cnt&1]; n == 0 {
				cnt++
			}
			cur += next[cnt&1]
		}
		next[0], next[1] = next[0]+2, next[1]-2
	}

	return *(*string)(unsafe.Pointer(&ret))
}
