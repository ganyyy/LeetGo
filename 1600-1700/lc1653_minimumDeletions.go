package main

func minimumDeletions(s string) int {
	// 从头迭代,
	// cb: b的当前数量
	// step: 完成平衡字符串的最少步数
	var cb, step int
	// 如果想要它平衡
	// 无外乎:
	// 1. A前边不能有B, 等价于删除前边所有的B
	// 2. B后边不能有A, 等价于删除前边所有的A
	for _, c := range s {
		if c == 'a' {
			step = step + 1 // 删除a所需要的开销
			if step > cb {
				step = cb // 保留a, 删除前边所有b所需要的开销
			}
		} else {
			cb++
		}
	}

	return step
}
