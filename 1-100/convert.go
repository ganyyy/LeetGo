package main

import (
	"bytes"
	"fmt"
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

func main() {
	fmt.Println(convert("A", 2))
	//LDREOEIIECIHNTSG
	//LDREOEIIECIHNTSG
}
