package main

import (
	"math"
	"sort"
)

var DIR = []int{-1, 0, 1, 0, -1}

func maximumMinutes(grid [][]int) int {

	// bfs 变种吧..
	// 但是, 怎么计算是个问题(?)

	// 比如: 先计算出多少步烧到人, 然后呆的时间肯定小于 这个步数
	// 然后二分测试?
	// 迭代火蔓延的BFS可以记录到达每一块草地的步数, 记录在高32位

	rowCount := len(grid)
	if rowCount == 0 {
		return Safe
	}
	colCount := len(grid[0])
	if colCount == 0 {
		return Safe
	}

	var current, nextStep []int
	// 获取所有的起点. x高32位, y低32位
	for x, row := range grid {
		for y, cell := range row {
			if cell == Fire {
				current = append(current, pack(x, y))
			}
		}
	}

	// 火势蔓延到所有位置
	var step int
	for len(current) != 0 {
		step++
		for _, pos := range current {
			x, y := unpack(pos)
			for nx, ay := range DIR[1:] {
				ax := DIR[nx]
				ax += x
				ay += y
				if uint(ax) >= uint(rowCount) || uint(ay) >= uint(colCount) {
					continue
				}
				state := grid[ax][ay]
				if state != Grass {
					continue
				}
				nextStep = append(nextStep, pack(ax, ay))
				grid[ax][ay] = pack(step, state)
			}
		}
		current, nextStep = nextStep, current[:0]
	}

	// for _, row := range grid {
	//     for _, v := range row {
	//         fire, _ := unpack(v)
	//         fmt.Printf("%v\t", fire)
	//     }
	//     fmt.Println()
	// }

	// 火烧到起点的时间. 如果不为0, 停留的时间肯定小于这个时间
	startFire, _ := unpack(grid[0][0])
	// 火烧到终点的时间
	endFire, _ := unpack(grid[rowCount-1][colCount-1])

	var maxSearchStep = startFire
	if maxSearchStep == 0 {
		maxSearchStep = rowCount*colCount - 1
	}

	// fmt.Println("maxSearchStep fire", startFire)

	endX, endY := rowCount-1, colCount-1

	var visited = make([][]bool, rowCount)
	for i := range visited {
		visited[i] = make([]bool, colCount)
	}

	// 二分判断是不是能到... 嗯?

	// minValid表示最小的有效停留时间
	var minStopValid = math.MaxInt32

	// firstNotValidStop 表示第一个不合法的停留时间. 也就是说, firstNotValidStop-1是最小的有效停留时间
	// 当stop满足时, 返回false, 表示可以增加停留时间; 当stop不满足时, 返回true, 表示需要减少停留时间
	// false会使stop增大, true会使stop减小
	firstNotValidStop := sort.Search(maxSearchStep, func(stop int) bool {
		if minStopValid != math.MaxInt32 && endFire == 0 {
			// 无脑返回即可
			return true
		}
		for _, row := range visited {
			clear(row)
		}
		visited[0][0] = true
		current, nextStep = current[:0], nextStep[:0]
		current = append(current, 0)
		var step int
		for len(current) != 0 {
			step++
			for _, pos := range current {
				x, y := unpack(pos)
				for nx, ay := range DIR[1:] {
					ax := x + DIR[nx]
					ay += y
					if uint(ax) >= uint(rowCount) || uint(ay) >= uint(colCount) {
						continue
					}
					fire, state := unpack(grid[ax][ay])
					if state != Grass {
						continue
					}
					// 到达这个方块的时间: 当前步数+停留时间
					arrival := step + stop
					if ax == endX && ay == endY && (fire == 0 || arrival <= fire) {
						// 可以到达终点
						// 应该算是同时到达的吧..?
						minStopValid = min(minStopValid, stop)
						return false
					}
					if fire != 0 && arrival >= fire || visited[ax][ay] {
						// 这个方块被烧了, 得绕行了
						continue
					}
					visited[ax][ay] = true
					nextStep = append(nextStep, pack(ax, ay))
				}
			}
			current, nextStep = nextStep, current[:0]
		}
		// 停留时间太长了, 需要缩短一些
		return true
	})

	if minStopValid != math.MaxInt32 {
		// 能到达终点, 并且火没有蔓延到起点, 说明无论怎么停留, 都是安全的
		if startFire == 0 {
			return Safe
		}
		return max(firstNotValidStop-1, 0)
	}
	return -1
}

const (
	Shift = 32
	MASK  = (1 << Shift) - 1

	Safe        = 1e9
	DefaultStep = int(1_0000_0000)
)

const (
	None = iota - 1
	Grass
	Fire
)

func unpack(n int) (int, int) { return n >> Shift, n & MASK }
func pack(v1, v2 int) int     { return (v1 << Shift) | v2 }

func maximumMinutes2(grid [][]int) int {
	rowCount := len(grid)
	if rowCount == 0 {
		return Safe
	}
	colCount := len(grid[0])
	if colCount == 0 {
		return Safe
	}

	var startQueue, nextQueue []int

	resetAndFilter := func(filter, v int, grid [][]int) {
		for x, row := range grid {
			for y, cell := range row {
				_, state := unpack(cell)
				row[y] = pack(v, state)
				if state == filter {
					startQueue = append(startQueue, pack(x, y))
				}
			}
		}
	}

	type WalkGridResult struct {
		Arrival     int // 到达终点(右下)的最短时间
		LeftArrival int // 到达终点左边的最短时间
		TopArrival  int // 到达终点右边的最短时间
	}

	walk := func() (ret WalkGridResult) {
		ret.Arrival = DefaultStep
		var step int
		for len(startQueue) != 0 {
			step++
			for _, pos := range startQueue {
				curX, curY := unpack(pos)
				for indexX, newY := range DIR[1:] {
					newY += curY
					newX := curX + DIR[indexX]
					if uint(newX) >= uint(rowCount) || uint(newY) >= uint(colCount) {
						continue
					}
					newStep, state := unpack(grid[newX][newY])
					if newStep != DefaultStep {
						continue
					}
					if state != Grass {
						continue
					}
					grid[newX][newY] = pack(step, state)
					nextQueue = append(nextQueue, pack(newX, newY))
				}
			}
			startQueue, nextQueue = nextQueue, startQueue[:0]
		}
		startQueue = startQueue[:0]
		nextQueue = nextQueue[:0]

		// 终点
		ret.Arrival, _ = unpack(grid[rowCount-1][colCount-1])
		// 左边, 列-1
		ret.LeftArrival, _ = unpack(grid[rowCount-1][colCount-2])
		// 上边, 行-1
		ret.TopArrival, _ = unpack(grid[rowCount-2][colCount-1])
		return
	}

	// defer resetAndFilter(-1, 0, grid)

	// 人
	resetAndFilter(None, DefaultStep, grid)

	startQueue = append(startQueue, 0)
	humanWalk := walk()
	if humanWalk.Arrival == DefaultStep {
		// 人不可达终点
		return -1
	}

	// 火
	resetAndFilter(Fire, DefaultStep, grid)
	fireWalk := walk()
	if fireWalk.Arrival == DefaultStep {
		// 火不可达终点
		return Safe
	}

	// 火会比人延后多久到达终点
	early := fireWalk.Arrival - humanWalk.Arrival
	if early < 0 {
		// 火比人先到达
		return -1
	}

	/*
		使用[human,fire]来表示到达某个点人和火的所消耗的最短时间

		情况1:
				  ........  ........
		........  [11, 11]  [12, 12]
		........  [10, 12]  [11, 13]
		如果延迟early(13-11=2)个周期,
		那么人在第12个周期(10+2)的时候, 人和火同时到达终点的左边,
		这是不可行的


		情况2:
				  ........  ........
		........  [0,  0 ]  [12, 12]
		........  [10, 14]  [11, 13]
		如果延迟early(13-11=2)个周期,
		那么人在第12个周期(10+2)的时候, 此时火无法烧到人,
		第13个周期(11+2)的时候, 人和火同时到达终点这是可行的
	*/

	if (humanWalk.LeftArrival != DefaultStep && humanWalk.LeftArrival+early < fireWalk.LeftArrival) ||
		(humanWalk.TopArrival != DefaultStep && humanWalk.TopArrival+early < fireWalk.TopArrival) {
		// 为什么要额外处理左边和上边呢?
		// 因为如果同时到达终点的话, 这种方法是可以逃出去的
		// 但是, 在终点之外的其他点是不允许同时到达的!
		// 所以, 这里得需要额外处理以下: 如果人能比火优先到达外围,
		// 说明延迟early可以让两者同时到达终点,
		return early
	}
	// 否则必须要提前1个周期
	return early - 1
}
