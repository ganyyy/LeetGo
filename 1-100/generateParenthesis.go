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

func generateParenthesis2(n int) []string {

	res := make([]string, 0)

	tmp := make([]byte, n*2)

	var generate func(int, int)

	generate = func(left, right int) {
		if left == right && right == n {
			res = append(res, string(tmp))
			return
		}
		if right > left || left > n {
			return
		}
		tmp[left+right] = '('
		generate(left+1, right)
		tmp[left+right] = ')'
		generate(left, right+1)
	}

	generate(0, 0)

	return res
}

func main() {
	res := generateParenthesis2(3)
	fmt.Println(res)
}
