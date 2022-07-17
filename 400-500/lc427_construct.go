//go:build ignore

package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	// 费脑子, 看不懂
	var dfs func([][]int, int, int) *Node
	dfs = func(rows [][]int, c0, c1 int) *Node {
		for _, row := range rows {
			for _, v := range row[c0:c1] {
				if v != rows[0][c0] { // 不是叶节点
					rMid, cMid := len(rows)/2, (c0+c1)/2
					return &Node{
						true,
						false,
						dfs(rows[:rMid], c0, cMid),
						dfs(rows[:rMid], cMid, c1),
						dfs(rows[rMid:], c0, cMid),
						dfs(rows[rMid:], cMid, c1),
					}
				}
			}
		}
		// 是叶节点
		return &Node{Val: rows[0][c0] == 1, IsLeaf: true}
	}
	return dfs(grid, 0, len(grid))
}

func construct2(grid [][]int) *Node {
	// 分两步
	// 第一步: 二维区域的前缀和
	var ln = len(grid)
	var sum = make([][]int, ln+1)
	sum[0] = make([]int, ln+1)
	for i, row := range grid {
		sum[i+1] = make([]int, ln+1)
		for j, v := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
		}
	}

	// 第二步: 递归求和
	var dfs func(startX, startY, endX, endY int) *Node

	dfs = func(startX, startY, endX, endY int) *Node {
		var total = sum[endX][endY] +
			sum[startX][startY] -
			sum[startX][endY] -
			sum[endX][startY]

		if total == 0 {
			// 区域内全部为0, 这是个叶子节点, 值为false
			return &Node{Val: false, IsLeaf: true}
		}
		if total == (endX-startX)*(endY-startY) {
			// 区域内全部为1, 这是个叶子节点, 值为true
			return &Node{Val: true, IsLeaf: true}
		}
		var midX = (endX + startX) / 2
		var midY = (endY + startY) / 2
		return &Node{
			IsLeaf: false,
			/*
				(startX,startY)...(startX,midY)...(startX,endY)
				...			   ...             ...
				(  midX,startY)...( midX ,midY)...(  midX,endY)
				...			   ...             ...
				(  endX,startY)...( endX ,midY)...(  endX,endY)
			*/
			TopLeft:     dfs(startX, startY, midX, midY),
			TopRight:    dfs(startX, midY, midX, endY),
			BottomLeft:  dfs(midX, startY, endX, midY),
			BottomRight: dfs(midX, midY, endX, endY),
		}
	}

	return dfs(0, 0, ln, ln)
}
