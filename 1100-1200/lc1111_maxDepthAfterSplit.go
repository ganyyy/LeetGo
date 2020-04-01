package main

// 说实话, 没看懂这题是要干啥
func maxDepthAfterSplit(seq string) []int {
	res := make([]int, len(seq))
	var isZero = 1
	// 相邻的(括号不是同一级的即可
	for i, v := range seq {
		if v == '(' {
			isZero ^= 1
			res[i] = isZero
		} else {
			// 右括号 和之前左括号一定是匹配的
			res[i] = isZero
			isZero ^= 1
		}
	}
	return res
}

func main() {

}
