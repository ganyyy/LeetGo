package main

func alphabetBoardPath(target string) string {
	// 固定顺序的打表, 完全可以直接算

	var cx, cy int // 起点是(0, 0)
	var res []byte
	for _, c := range target {
		nx := int(c-'a') / 5
		ny := int(c-'a') % 5
		if nx < cx {
			// Z单独一行, 所以Z到其他位置需要先上去U
			for j := 0; j < cx-nx; j++ {
				res = append(res, 'U')
			}
		}
		if ny < cy {
			// 同样的道理, 其他字母想到Z必须要先左去L
			for j := 0; j < cy-ny; j++ {
				res = append(res, 'L')
			}
		}
		if nx > cx {
			for j := 0; j < nx-cx; j++ {
				res = append(res, 'D')
			}
		}
		if ny > cy {
			for j := 0; j < ny-cy; j++ {
				res = append(res, 'R')
			}
		}
		res = append(res, '!')
		cx = nx
		cy = ny
	}
	return string(res)
}
