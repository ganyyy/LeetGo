//go:build ignore

package main

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func containVirus(isInfected [][]int) (ans int) {
	m, n := len(isInfected), len(isInfected[0])
	for {
		neighbors := []map[pair]struct{}{}
		firewalls := []int{}
		for i, row := range isInfected {
			for j, infected := range row {
				if infected != 1 {
					continue
				}
				// bfs查看周围
				q := []pair{{i, j}}
				neighbor := map[pair]struct{}{}
				firewall, idx := 0, len(neighbors)+1
				row[j] = -idx
				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					for _, d := range dirs {
						if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n {
							if isInfected[x][y] == 1 {
								// 检查标记
								q = append(q, pair{x, y})
								isInfected[x][y] = -idx
							} else if isInfected[x][y] == 0 {
								// 预感染标记
								firewall++
								neighbor[pair{x, y}] = struct{}{}
							} else {
								// 这里就是2的情况了, 啥都不用做
							}
						}
					}
				}
				neighbors = append(neighbors, neighbor)
				firewalls = append(firewalls, firewall)
			}
		}

		// 都被感染了
		if len(neighbors) == 0 {
			break
		}

		// 找出预感染最大的区域
		// 上边的idx范围是[1,n]
		// 这里的idx范围是[0,n)
		idx := 0
		for i := 1; i < len(neighbors); i++ {
			if len(neighbors[i]) > len(neighbors[idx]) {
				idx = i
			}
		}

		// 上防火墙
		ans += firewalls[idx]

		// 最后一片区域了, 可以直接跳出
		if len(neighbors) == 1 {
			break
		}

		// 复原原始数据
		for _, row := range isInfected {
			for j, v := range row {
				if v < 0 {
					if v != -idx-1 {
						// 下次继续参与计算
						row[j] = 1
					} else {
						// 已经被围死了, 下次不参与计算
						row[j] = 2
					}
				}
			}
		}

		// 感染周围区域
		for i, neighbor := range neighbors {
			if i != idx {
				for p := range neighbor {
					isInfected[p.x][p.y] = 1
				}
			}
		}

	}
	return
}
