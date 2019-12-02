package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	getRes(&res, make([]byte, 0), 0, 0, n)
	return res
}

func getRes(res *[]string, str []byte, l, r, total int) {
	if l > total || l < r || r > total {
		return
	} else {
		if l == r && l == total {
			*res = append(*res, string(str))
		} else {
			getRes(res, append(str, '('), l+1, r, total)
			getRes(res, append(str, ')'), l, r+1, total)
		}
	}
}

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}
