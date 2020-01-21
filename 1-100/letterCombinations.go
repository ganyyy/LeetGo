package main

import "fmt"

var mapStr = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func recall(before []byte, digit string, start int, res *[]string) {
	if start == len(digit) {
		*res = append(*res, string(before))
		return
	}
	arr := mapStr[digit[start]]
	for _, _v := range arr {
		tmp := append(before, _v)
		recall(tmp, digit, start+1, res)
	}
}

func Test2(digit string) []string {
	res := make([]string, 0)
	if len(digit) == 0 {
		return res
	}

	begin := mapStr[digit[0]]

	for _, _v := range begin {
		before := []byte{_v}
		recall(before, digit, 1, &res)
	}
	return res
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	res = append(res, "")
	for _, _i := range digits {
		size := len(res)
		arr := mapStr[uint8(_i)]
		for _j := 0; _j < size; _j++ {
			_s := res[0]
			res = res[1:]
			for _, _l := range arr {
				res = append(res, _s+string(_l))
			}
		}
	}
	return res
}

func main() {
	res := Test2("232")
	fmt.Println(res)
}
