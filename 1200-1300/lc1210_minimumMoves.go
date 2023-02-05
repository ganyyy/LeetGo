package main

type Node struct {
	X int // 蛇尾的X
	Y int // 蛇尾的Y
	S int // 当前是水平状态(0)or垂直状态(1)
}

var dirs = []Node{
	{X: 1}, // 向右(蛇头蛇尾一起动)
	{Y: 1}, // 向下(蛇头蛇尾一起动)
	{S: 1}, // 转向(只动蛇头)
}

func minimumMoves(g [][]int) int {
	n := len(g)
	// 访问标记
	vis := make([][][2]bool, n)
	for i := range vis {
		vis[i] = make([][2]bool, n)
	}
	vis[0][0][0] = true // 蛇尾的初始位置
	var queue []Node
	queue = append(queue, Node{}) // {0, 0, 0}
	for step := 1; len(queue) > 0; step++ {
		tmp := queue
		queue = nil
		for _, t := range tmp {
			for _, d := range dirs {
				//					| 	d & t      |  newX  |  newY  |  newState  |  x2  |  y2  |
				// 1. 水平方向+水平走, |{d.Y=1, t.S=0}|   X    |  Y+1   |     0      |  X   |  Y+2 |
				// 2. 竖直方向+竖直走, |{d.X=1, t.S=1}|  X+1   |   Y    |     1      |  X+2 |  Y   |
				// 3. 水平方向+竖直,  |{d.Y=1, t.S=1}|   X    |  Y+1   |      1     |  X+1 |  Y+1 |
				// 4. 竖直方向+水平,  |{d.X=1, t.S=0}|  X+1   |   Y    |      0     |  X+1 |  Y+1 |
				// 5. 水平转竖直,     |{d.S=1, t.S=0}|   X    |   Y    |     1     |  X+1 |   Y  |
				// 6. 竖直转水平,     |{d.S=1, t.S=1}|   X    |   Y    |     0     |   X  |  Y+1 |
				newX, newY, newState := t.X+d.X, t.Y+d.Y, t.S^d.S // 下一步蛇尾的位置, 以及是否转向
				x2, y2 := newX+newState, newY+(newState^1)        // 下一步蛇头的位置(如果水平转向, [x,y+1]; 如果竖直转向, [x+1, y])
				if x2 < n && y2 < n &&                            // 合法性检查
					!vis[newX][newY][newState] && // 没有访问过的位置
					g[newX][newY] == 0 && // 蛇尾没有障碍
					g[x2][y2] == 0 && // 蛇头没有障碍
					(d.S == 0 || g[newX+1][newY+1] == 0) /*非旋转, 或者蛇尾的右下角位置为空*/ {
					if newX == n-1 && newY == n-2 { // 此时蛇头一定在 (n-1,n-1)
						return step
					}
					vis[newX][newY][newState] = true
					queue = append(queue, Node{newX, newY, newState})
				}
			}
		}
	}
	return -1
}
