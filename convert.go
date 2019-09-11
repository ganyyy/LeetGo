package main

import "fmt"

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
				val := sub*(2 * mul - 1) - j
				if val < size {
					res = append(res, s[val])
				}
			}
			j += sub
			mul +=1
		}
	}
	return string(res)
}

func main() {
	fmt.Println(convert("A", 2))
	//LDREOEIIECIHNTSG
	//LDREOEIIECIHNTSG
}
