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

func snakesAndLadders2(board [][]int) int {
	// mark
	n := len(board)

	var id2rc = func(id int) (r, c int) {
		r, c = (id-1)/n, (id-1)%n
		if r%2 == 1 {
			// 奇数行, 从右往左
			c = n - 1 - c
		}
		r = n - 1 - r
		return
	}

	vis := make([]bool, n*n+1)
	type pair struct{ id, step int }
	q := []pair{{1, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for i := 1; i <= 6; i++ {
			nxt := p.id + i
			if nxt > n*n { // 超出边界
				break
			}
			r, c := id2rc(nxt)   // 得到下一步的行列
			if board[r][c] > 0 { // 存在蛇或梯子
				nxt = board[r][c]
			}
			if nxt == n*n { // 到达终点
				return p.step + 1
			}
			if !vis[nxt] {
				vis[nxt] = true
				q = append(q, pair{nxt, p.step + 1}) // 扩展新状态
			}
		}
	}
	return -1
}
