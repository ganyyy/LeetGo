package main

func minimumDeletions(s string) int {
	// 从头迭代,
	var cb, step int
	for _, c := range s {
		if c == 'a' {
			step = step + 1 // 删除这个a所需要的开销
			if step > cb {
				step = cb // 保留a, 删除前边所有b所需要的开销
			}
		} else {
			cb++
		}
	}

	return step
}
