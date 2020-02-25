package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	p1, p2 := 1, 2
	for i := 3; i <= n; i++ {
		p2, p1 = p1+p2, p2
	}
	return p2
}

func main() {
	fmt.Println(climbStairs(10))
}
