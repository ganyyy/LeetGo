package main

const (
	Shift = 14
	Mask  = (1 << Shift) - 1
	Max   = 1e4
)

var dir = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func escapeGhostsBad(ghosts [][]int, target []int) bool {
	// BFS
	var targetPos = genPos(target[0], target[1])
	if targetPos == 0 {
		return true
	}
	var ghostsQueue = make([]int, 0, len(ghosts))

	var set = map[int]struct{}{}
	for _, g := range ghosts {
		var p = genPos(g[0], g[1])
		if p == targetPos {
			return false
		}
		ghostsQueue = append(ghostsQueue, p)
		addToSet(set, p)
	}

	var nextQueue = []int{0}
	var visited = map[int]struct{}{}
	addToSet(visited, 0)

	for len(nextQueue) != 0 {
		// 先遍历所有阻挡的可选路程
		var ln = len(ghostsQueue)
		for _, p := range ghostsQueue[:ln] {
			var x, y = parsePos(p)
			for _, d := range dir {
				var nx, ny = x + d[0], y + d[1]
				if nx > Max || nx < -Max || ny > Max || ny < -Max {
					continue
				}
				var np = genPos(nx, ny)
				if np == targetPos {
					return false
				}
				if checkInSet(set, np) {
					continue
				}
				addToSet(set, np)
				ghostsQueue = append(ghostsQueue, np)
			}
		}
		ghostsQueue = ghostsQueue[ln:]

		// 再看自己的
		ln = len(nextQueue)
		for _, p := range nextQueue[:ln] {
			var x, y = parsePos(p)
			for _, d := range dir {
				var nx, ny = x + d[0], y + d[1]
				if nx > Max || nx < -Max || ny > Max || ny < -Max {
					continue
				}
				var np = genPos(nx, ny)
				if np == targetPos {
					return true
				}
				if checkInSet(set, np) {
					continue
				}
				if checkInSet(visited, np) {
					continue
				}
				addToSet(visited, np)
				nextQueue = append(nextQueue, np)
			}
		}
	}
	return false
}

func addToSet(set map[int]struct{}, x int) {
	set[x] = struct{}{}
}

func checkInSet(set map[int]struct{}, x int) bool {
	_, ok := set[x]
	return ok
}

func parsePos(p int) (x, y int) {
	return p >> Shift, p & Mask
}

func genPos(x, y int) int {
	return (x << Shift) | y
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	// 如果鬼能够更快的到达终点, 那么玩家一定会失败.
	var dis = abs(target[0]) + abs(target[1])
	for _, g := range ghosts {
		if abs(g[0]-target[0])+abs(g[1]-target[1]) <= dis {
			return false
		}
	}
	return true
}
