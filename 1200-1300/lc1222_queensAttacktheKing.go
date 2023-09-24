package main

import "math"

func queensAttackTheKing(queens [][]int, king []int) [][]int {
	// 以国王为中心, 从8个方向开始搜索

	// 二维坐标转换为一维坐标
	queenPos := make(map[int]bool, len(queens))
	for _, queen := range queens {
		x, y := queen[0], queen[1]
		queenPos[x*8+y] = true
	}

	// 答案的上限是8个
	var ans = make([][]int, 0, 8)
	// 迭代8个方向
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			kx, ky := king[0]+dx, king[1]+dy
			for uint(kx) < 8 && uint(ky) < 8 {
				pos := kx*8 + ky
				if _, ok := queenPos[pos]; ok {
					ans = append(ans, []int{kx, ky})
					break
				}
				kx += dx
				ky += dy
			}
		}
	}
	return ans
}

func queensAttackTheKingQueue(queens [][]int, king []int) [][]int {
	// 压缩+队列+贪心
	// 由于国王的位置是固定的, 所以可以将其视为原点

	// 范围是 [-1, 1]. 为了保证key的唯一性, 可以将其映射到 [0, 2]
	// 获取相对位置: 1表示在前边, 0在同一行/列, -1表示在后边
	// 这里直接加了1, 将其映射到 [0, 2]的范围内
	var sgn = func(x int) int {
		if x > 0 {
			return 2
		} else if x == 0 {
			return 1
		} else {
			return 0
		}
	}

	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var candidates [9][3]int
	for i := range candidates {
		candidates[i][2] = math.MaxInt32
	}
	var cnt int
	kx, ky := king[0], king[1]
	for _, queen := range queens {
		qx, qy := queen[0], queen[1]
		x, y := qx-kx, qy-ky
		// 同行/同列/同对角线
		if x == 0 || y == 0 || abs(x) == abs(y) {
			// 获取对应的偏置值, 这个很有意思诶
			// 总共组合有8中, 因为不会出现(0, 0)这种情况
			dx, dy := sgn(x), sgn(y)
			// 由此, 可得出的key在某个方向上肯定是唯一的
			// 修改一下编码规则.
			// dx 作为高位, dy 作为低位
			key := dx*3 + dy
			// 对比距离就是简单的对比绝对差值的和就行, 如果当前节点更近就进行替换
			if d := candidates[key][2]; d > abs(x)+abs(y) {
				if d == math.MaxInt32 {
					cnt++
				}
				candidates[key] = [3]int{qx, qy, abs(x) + abs(y)}
			}
		}
	}

	var ans = make([][]int, 0, cnt)
	for _, value := range candidates {
		if value[2] == math.MaxInt32 {
			continue
		}
		ans = append(ans, []int{value[0], value[1]})
	}
	return ans
}
