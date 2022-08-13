package main

func maxChunksToSorted(arr []int) int {
	var st []int
	for _, x := range arr {
		if len(st) == 0 || x >= st[len(st)-1] {
			// 相同的数字分在不同的组, 这样才能保证组的数量最多
			st = append(st, x)
		} else {
			// 将其放置到一个合适的位置 ( st[x-1] < x < st[top] )
			mx := st[len(st)-1]
			st = st[:len(st)-1]
			for len(st) > 0 && st[len(st)-1] > x {
				st = st[:len(st)-1]
			}
			st = append(st, mx)
		}
	}
	return len(st)
}
