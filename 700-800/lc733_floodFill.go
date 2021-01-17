package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	old := image[sr][sc]

	// 处理相同的情况
	if old == newColor {
		return image
	}
	lr := len(image)
	lc := len(image[0])

	var helper func(int, int)

	helper = func(r, c int) {
		image[r][c] = newColor
		v := r - 1
		if v >= 0 && image[v][c] == old {
			helper(v, c)
		}
		v = r + 1
		if v < lr && image[v][c] == old {
			helper(v, c)
		}

		v = c - 1
		if v >= 0 && image[r][v] == old {
			helper(r, v)
		}
		v = c + 1
		if v < lc && image[r][v] == old {
			helper(r, v)
		}
	}

	helper(sr, sc)
	return image
}
