package main

func surfaceArea(grid [][]int) int {
	// 计算总的立方体个数
	// 计算上边和左边重叠的面以及自身重叠的面
	var faces, cubes int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			v := grid[i][j]
			if v == 0 {
				continue
			}
			cubes += v
			if v > 1 {
				// 自身重叠的面
				faces += v - 1
			}
			if i > 0 {
				// 上边重叠的面
				faces += min(v, grid[i-1][j])
			}
			if j > 0 {
				// 左边重叠的面
				faces += min(v, grid[i][j-1])
			}
		}
	}
	return cubes*6 - faces*2
}

func main() {

}
