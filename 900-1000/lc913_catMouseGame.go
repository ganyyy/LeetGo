package main

const (
	draw     = 0
	mouseWin = 1
	catWin   = 2
)

func catMouseGame(graph [][]int) int {
	n := len(graph)
	// 鼠在节点i, 猫在节点j, 第k轮次的结果
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n*2)
			for k := range dp[i][j] {
				// 初始-1
				dp[i][j][k] = -1
			}
		}
	}

	var getResult, getNextResult func(int, int, int) int
	getResult = func(mouse, cat, turns int) int {
		if turns == n*2 {
			return draw
		}
		res := dp[mouse][cat][turns]
		if res != -1 {
			return res
		}
		if mouse == 0 {
			// 走到了终点
			res = mouseWin
		} else if cat == mouse {
			// 鼠被猫抓到了
			res = catWin
		} else {

			res = getNextResult(mouse, cat, turns)
		}
		dp[mouse][cat][turns] = res
		return res
	}
	getNextResult = func(mouse, cat, turns int) int {
		curMove := mouse
		if turns%2 == 1 {
			curMove = cat
		}
		defaultRes := mouseWin
		if curMove == mouse {
			defaultRes = catWin
		}
		res := defaultRes
		// 当前棋子在当前位置可选的移动路径
		for _, next := range graph[curMove] {
			if curMove == cat && next == 0 {
				continue
			}
			nextMouse, nextCat := mouse, cat
			if curMove == mouse {
				nextMouse = next
			} else if curMove == cat {
				nextCat = next
			}
			nextRes := getResult(nextMouse, nextCat, turns+1)
			if nextRes != defaultRes {
				res = nextRes
				if res != draw {
					break
				}
			}
		}
		return res
	}
	// 起始状态下, 猫在2, 鼠在1, 第0轮(鼠先走)
	return getResult(1, 2, 0)
}
