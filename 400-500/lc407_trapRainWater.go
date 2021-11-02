//go:build ignore

package main

type pair struct {
	x, y, h int
}

type PriQueue struct {
	capacity int
	count    int
	arr      []*pair
}

func trapRainWater(heightMap [][]int) (ans int) {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	m, n := len(heightMap), len(heightMap[0])
	p := Constructor(m * n)
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				p.add(&pair{x: i, y: j, h: heightMap[i][j]})
				visit[i][j] = true
			}
		}
	}

	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	for !p.isEmpty() {
		cur := p.poll()
		x, y, h := cur.x, cur.y, cur.h
		for k := 0; k < 4; k++ {
			newx, newy := x+dr[k], y+dc[k]
			if 0 <= newx && newx < m && 0 <= newy && newy < n && !visit[newx][newy] {
				if heightMap[newx][newy] < h {
					ans += h - heightMap[newx][newy]
				}
				p.add(&pair{x: newx, y: newy, h: max(h, heightMap[newx][newy])})
				visit[newx][newy] = true
			}
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Constructor(size int) *PriQueue {
	p := new(PriQueue)
	p.capacity = size
	p.count = 0
	p.arr = make([]*pair, size+1)
	return p
}

func (p *PriQueue) isEmpty() bool {
	return p.count == 0
}

func (p *PriQueue) add(data *pair) {
	if p.count >= p.capacity {
		return
	}
	p.count++
	p.arr[p.count] = data
	i := p.count
	for i/2 > 0 && p.arr[i/2].h > p.arr[i].h {
		swap(p.arr, i, i/2)
		i /= 2
	}
}

func (p *PriQueue) poll() *pair {
	res := p.arr[1]
	p.arr[1] = p.arr[p.count]
	p.count--
	heapify(p.arr, p.count, 1)
	return res
}

func heapify(arr []*pair, size int, i int) {
	for {
		pos := i
		if i*2 <= size && arr[i*2].h < arr[i].h {
			pos = i * 2
		}
		if i*2+1 <= size && arr[i*2+1].h < arr[pos].h {
			pos = i*2 + 1
		}
		if pos == i {
			break
		}
		swap(arr, i, pos)
		i = pos
	}
}

func swap(arr []*pair, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
