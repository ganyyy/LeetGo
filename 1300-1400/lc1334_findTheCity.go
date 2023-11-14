package main

import (
	"fmt"
	"math"
)

func findTheCityF(n int, edges [][]int, distanceThreshold int) int {
	ans := []int{0x3f3f3f3f, -1}
	// mp[i][j] 表示 i 到 j 的最短距离
	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mp[i][j] = 0x3f3f3f3f
		}
	}

	for _, eg := range edges {
		u, v, w := eg[0], eg[1], eg[2]
		mp[u][v], mp[v][u] = w, w
	}
	for k := 0; k < n; k++ {
		// 自己到自己的距离为0
		mp[k][k] = 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// 通过k点, 更新i到j的最短距离:
				// i到k的距离 + k到j的距离 和 i到j的距离 取最小值
				mp[i][j] = min(mp[i][j], mp[i][k]+mp[k][j])
			}
		}
	}
	for i := 0; i < n; i++ {
		// 计算以i为起点的城市, 能到达的城市数量
		cnt := 0
		for j := 0; j < n; j++ {
			if mp[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= ans[0] {
			ans[0], ans[1] = cnt, i
		}
	}
	return ans[1]
}

func findTheCityD(n int, edges [][]int, distanceThreshold int) int {
	ans := []int{math.MaxInt32 / 2, -1}
	mp := make([][]int, n)
	dis := make([][]int, n)
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, n)
		dis[i] = make([]int, n)
		vis[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			mp[i][j] = math.MaxInt32 / 2
			dis[i][j] = math.MaxInt32 / 2
			vis[i][j] = false
		}
	}

	for _, eg := range edges {
		from, to, weight := eg[0], eg[1], eg[2]
		mp[from][to], mp[to][from] = weight, weight
	}
	for i := 0; i < n; i++ {
		dis[i][i] = 0
		for j := 0; j < n; j++ {
			t := -1
			for k := 0; k < n; k++ {
				if !vis[i][k] && (t == -1 || dis[i][k] < dis[i][t]) {
					t = k
				}
			}
			for k := 0; k < n; k++ {
				dis[i][k] = min(dis[i][k], dis[i][t]+mp[t][k])
			}
			vis[i][t] = true
		}
	}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if dis[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= ans[0] {
			ans[0] = cnt
			ans[1] = i
		}
	}
	return ans[1]
}

// 家花不如野花香
func findTheCity2(n int, edges [][]int, distanceThreshold int) int {
	var path = make([][][2]int, n)
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		path[from] = append(path[from], [2]int{to, weight})
		from, to = to, from
		path[from] = append(path[from], [2]int{to, weight})
	}
	var visited = make([]bool, n)
	var dist = make([]int, n)

	var dijkstra = func(root int) {
		clear(visited)
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		// 起点
		dist[root] = 0
		for range visited {
			var next = -1
			for p := 0; p < n; p++ {
				if visited[p] {
					continue
				}
				if next == -1 || dist[p] < dist[next] {
					// 找到距离最近的点
					next = p
				}
			}
			if next == -1 {
				break
			}
			visited[next] = true
			for _, edge := range path[next] {
				// 通过next节点, 更新next的邻接节点的最短距离
				dist[edge[0]] = min(dist[edge[0]], dist[next]+edge[1])
			}
		}
	}
	const debug = false
	var minCount = math.MaxInt32
	var minCity = -1
	for i := range visited {
		dijkstra(i)
		if debug {
			fmt.Println("\nroot", i)
			for n, d := range dist {
				fmt.Printf("{%v, %v}, ", n, d)
			}
		}
		var count int
		for _, d := range dist {
			if d <= distanceThreshold {
				count++
			}
		}
		if minCount >= count {
			minCount = count
			minCity = i
		}
	}
	return minCity
}
