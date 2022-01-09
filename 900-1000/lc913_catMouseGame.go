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
			// 步数走够了, 往后再走的步数一定是重复的. 所以平局处理
			return draw
		}
		res := dp[mouse][cat][turns]
		if res != -1 {
			// 不是默认值可以直接返回--记忆化搜索!
			return res
		}
		if mouse == 0 {
			// 走到了终点
			res = mouseWin
		} else if cat == mouse {
			// 鼠被猫抓到了
			res = catWin
		} else {
			// 需要通过下一轮来确定结果
			res = getNextResult(mouse, cat, turns)
		}
		dp[mouse][cat][turns] = res
		return res
	}
	getNextResult = func(mouse, cat, turns int) int {
		// 当前轮次谁走棋
		curMove := mouse
		if turns%2 == 1 {
			curMove = cat
		}
		// 对于当前移动的玩家, 其对应的必败状态
		defaultRes := mouseWin
		if curMove == mouse {
			defaultRes = catWin
		}
		res := defaultRes
		// 当前棋子在当前位置可选的移动路径
		for _, next := range graph[curMove] {
			// 猫不能到达0位置
			if curMove == cat && next == 0 {
				continue
			}
			nextMouse, nextCat := mouse, cat
			if curMove == mouse {
				nextMouse = next
			} else if curMove == cat {
				nextCat = next
			}

			// 这里依赖于对手的下一步棋的结果决定当前这步棋的结果
			nextRes := getResult(nextMouse, nextCat, turns+1)
			// 如果走了当前的点, 对手能必胜, 那么就需要尝试下一个可行走的点
			if nextRes != defaultRes {
				res = nextRes
				// 只要能赢一次, 就不需要再遍历接下来的路径了
				// 如果是平局, 就需要再继续查找可能胜利的走法
				if res != draw {
					break
				}
			}
		}
		return res
	}
	// 起始状态下, 鼠在1, 猫在2, 第0轮(鼠先走)
	return getResult(1, 2, 0)
}
