package main

func findTheWinner(n int, k int) int {
	// 模拟法
	var queue = make([]int, n)
	for i := range queue {
		queue[i] = i + 1
	}

	var cur int
	for len(queue) != 1 {
		cur = (cur + k - 1) % len(queue)
		queue = append(queue[:cur], queue[cur+1:]...)
	}
	return queue[0]
}

func mathSolve(n, k int) int {
	// 求最后一个剩下的数, 在原始数组中的位置
	/*
	   k退1
	   n  pos
	   1  0
	   2  (f(1)+k)%2
	   3  (f(2)+k)%3
	   ...
	   n  (f(n-1)+k)%n
	*/
	var pos int
	for i := 2; i <= n; i++ {
		pos = (pos + k) % i
	}
	return pos + 1 // 位置+1就是对应的顺序
}
