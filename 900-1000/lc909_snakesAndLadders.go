package main

func init() {
	/*
		写点啥呢
		也就这么着吧, 认了吧.
		只会越来越差, 到最后就是一个废物罢了
		怎么混的?
	*/
}

func snakesAndLadders(board [][]int) int {
	// 也是队列吧

	// 整个倒序, 省的这么麻烦

	var m, n = len(board), len(board[0])

	for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
		board[i], board[j] = board[j], board[i]
	}

	var target = m*n - 1

	var queue1, queue2 []int
	var visited = make(map[int]bool)

	queue1 = []int{0}
	visited[0] = true

	var step int
	for len(queue1) != 0 {
		// fmt.Println(queue1)
		for _, pos := range queue1 {
			if pos == target {
				return step
			}
			for i := pos + 1; i <= pos+6 && i <= target; i++ {
				var r = i / n
				var c = i % n
				if r&1 != 0 {
					c = n - c - 1
				}
				v := board[r][c]
				if v == -1 {
					if visited[i] {
						continue
					}
					if i == target {
						return step + 1
					}
					queue2 = append(queue2, i)
					visited[i] = true
				} else if !visited[v-1] {
					if v-1 == target {
						return step + 1
					}
					queue2 = append(queue2, v-1)
					visited[v-1] = true
				}
			}
		}
		step++

		queue1, queue2 = queue2, queue1
		queue2 = queue2[:0]
	}

	return -1
}
