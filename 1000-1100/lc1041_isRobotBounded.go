//go:build ignore

package main

var dirs = [4][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func isRobotBounded(instructions string) bool {
	// 4个方向
	// 0: 北, 1: 东, 2: 南, 3: 西
	// 这个数组有点意思哈
	robot := []int{0, 0}
	d := 0
	for i := 0; i < len(instructions); i++ {
		if instructions[i] == 'G' {
			robot[0] += dirs[d][0]
			robot[1] += dirs[d][1]
		} else if instructions[i] == 'L' {
			d = (d + 3) % 4 // 左转90°
		} else {
			d = (d + 1) % 4 // 右转90°
		}
	}
	// 执行完一次全部指令后, 只要方向发生了变化; 或者回到了原点, 就认为一直在一个圈中
	return d != 0 || (robot[0] == 0 && robot[1] == 0)
}
