package main

// var VALID = [][2]int{
//     {-1, -2},
//     {-2, -1},
//     {-2, 1},
//     {-1, 2},
//     {1, 2},
//     {2, 1},
//     {2, -1},
//     {1, -2},
// }

func checkValidGrid(grid [][]int) bool {
	// 按照顺序, 反向映射到对应的坐标
	n := len(grid)
	if n < 2 {
		return true
	}
	var seq = make([][2]int8, n*n)

	for i, row := range grid {
		for j, step := range row {
			seq[step] = [2]int8{int8(i), int8(j)}
		}
	}

	// 验证有效性
	var pre = seq[0]
	if pre[0] != pre[1] || pre[0] != 0 {
		return false
	}
	for _, step := range seq[1:] {
		var valid bool
		// 优化: 绝对值相加等于3!
		// for _, n := range VALID {
		//     if pre[0]+n[0] == step[0] && pre[1]+n[1] == step[1] {
		//         valid = true
		//         break
		//     }
		// }
		valid = abs(step[0]-pre[0])+abs(step[1]-pre[1]) == 3
		if !valid {
			return false
		}
		pre = step
	}
	return true
}

func abs(a int8) int8 {
	if a < 0 {
		return -a
	}
	return a
}
