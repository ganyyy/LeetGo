package main

var dir = [][2]int16{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	// 应该是DP吧,,,,

	// 洪范法, 应该可以求解

	var mm = make([][]int, m)
	for i := range mm {
		mm[i] = make([]int, n)
	}

	const shift = 7
	const mask = (1 << shift) - 1

	var bind = func(r, c int16) int16 {
		return r<<shift + c
	}

	var unbound = func(v int16) (r, c int16) {
		return v >> shift, v & mask
	}

	const mod = 1e9 + 7

	//bfs
	var step int
	var cnt uint
	var oldSet = make(map[int16]int)
	var tmpSet = make(map[int16]int)
	oldSet[bind(int16(startRow), int16(startColumn))] = 1
	for len(oldSet) != 0 && step < maxMove {
		for v, cc := range oldSet {
			// 判断是否可以出界, 如果可以的话, 更新mm中的值
			var p = cc
			var r, c = unbound(v)
			if mm[r][c] != 0 {
				// fmt.Println(mm[r][c], p, mm[r][c]*p)
				cnt = cnt + uint(mm[r][c]*p)
			} else {
				// 判断出界的个数
				var tmp int
				if r-1 < 0 {
					tmp++
				}
				if r+1 >= int16(m) {
					tmp++
				}
				if c-1 < 0 {
					tmp++
				}
				if c+1 >= int16(n) {
					tmp++
				}
				if tmp != 0 {
					mm[r][c] = tmp
					cnt = cnt + uint(mm[r][c]*p)
				}
			}
			// 临近点入队
			for _, d := range dir {
				if nr, nc := r+d[0], c+d[1]; nr >= 0 && nr < int16(m) && nc >= 0 && nc < int16(n) {
					tmpSet[bind(nr, nc)] = (tmpSet[bind(nr, nc)] + p) % mod
				}
			}
			delete(oldSet, v)
		}
		tmpSet, oldSet = oldSet, tmpSet
		step++
	}

	return int(cnt % uint(mod))
}

func findPaths2(m int, n int, maxMove int, startRow int, startColumn int) int {
	// 应该是DP吧,,,,

	// 洪范法, 应该可以求解

	var mm = make([][]int, m)
	for i := range mm {
		mm[i] = make([]int, n)
	}

	const shift = 6
	const mask = (1 << shift) - 1

	var bind = func(r, c int16) int16 {
		return r<<shift + c
	}

	var unbound = func(v int16) (r, c int16) {
		return v >> shift, v & mask
	}

	var queue []int16 // 当前轮次可以行走的路径点
	queue = append(queue, bind(int16(startRow), int16(startColumn)))

	const mod = 1e10 + 9

	//bfs
	var step, cnt int
	for len(queue) != 0 && step < maxMove {
		var ln = len(queue)
		for i := 0; i < ln; i++ {
			// 判断是否可以出界, 如果可以的话, 更新 m中的值
			var r, c = unbound(queue[i])
			if mm[r][c] != 0 {
				cnt = (cnt + mm[r][c]) % mod
			} else {
				var tmp int
				if r-1 < 0 {
					tmp++
				}
				if r+1 >= int16(m) {
					tmp++
				}
				if c-1 < 0 {
					tmp++
				}
				if c+1 >= int16(n) {
					tmp++
				}
				mm[r][c] = tmp
				cnt = (cnt + tmp) % mod
			}
			// 临近点入队
			if r-1 >= 0 {
				queue = append(queue, bind(r-1, c))
			}
			if r+1 < int16(m) {
				queue = append(queue, bind(r+1, c))
			}
			if c-1 >= 0 {
				queue = append(queue, bind(r, c-1))
			}
			if c+1 < int16(n) {
				queue = append(queue, bind(r, c+1))
			}
		}
		queue = queue[ln:]
		//fmt.Println(queue, mm, cnt)
		step++
	}

	return cnt
}

func findPathsDP(m int, n int, maxMove int, startRow int, startColumn int) int {
	if startRow < 0 || startRow >= m || startColumn < 0 || startColumn >= n {
		return 0
	}
	dirs := [4][2]int64{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	//bfs
	q := make([][]uint64, m)
	tq := make([][]uint64, m)
	for i := 0; i < m; i++ {
		q[i] = make([]uint64, n)
		tq[i] = make([]uint64, n)
	}
	q[startRow][startColumn] = 1
	res := uint64(0)
	for maxMove > 0 {
		for i := int64(0); i < int64(m); i++ {
			for j := int64(0); j < int64(n); j++ {
				if q[i][j] > 0 {
					for _, d := range dirs {
						if i+d[0] < 0 || i+d[0] >= int64(m) || j+d[1] < 0 || j+d[1] >= int64(n) {
							res += q[i][j]
							res %= 1e9 + 7
						} else {
							tq[i+d[0]][j+d[1]] += q[i][j]
							tq[i+d[0]][j+d[1]] %= 1e9 + 7
						}
					}
					q[i][j] = 0
				}
			}
		}
		q, tq = tq, q
		res %= 1e9 + 7
		maxMove--
	}
	return int(res)
}

func main() {
	println(findPaths(36,
		5,
		50,
		15,
		3))
	//println(findPaths2(4, 7, 15, 0, 1))
}
