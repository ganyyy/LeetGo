package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int

	var c func(int)
	t := make([]int, 0, k)

	c = func(start int) {
		if len(t) == k {
			res = append(res, append([]int(nil), t...))
			return
		}
		for i := start; i <= n; i++ {
			t = append(t, i)
			c(i + 1)
			t = t[:len(t)-1]
		}
	}
	c(1)
	return res
}

func main() {
	fmt.Println(combine(5, 4))
}
