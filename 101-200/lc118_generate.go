package main

import "fmt"

func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	var t []int
	for i := 1; i <= numRows; i++ {
		t = make([]int, i)
		t[0], t[i-1] = 1, 1
		for j := 1; j < i-1; j++ {
			t[j] = res[i-2][j-1] + res[i-2][j]
		}
		res[i-1] = t
	}

	return res
}

func main() {
	for _, v := range generate(10) {
		fmt.Println(v)
	}
}
