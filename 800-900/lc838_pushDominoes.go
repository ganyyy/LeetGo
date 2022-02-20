package main

func pushDominoes(dominoes string) string {
	var n = len(dominoes)
	var bs = []byte(dominoes)

	var l, left = 0, byte('L') // 左哨兵

	for l < n {
		var i, j int
		for i = l; i < n && bs[i] == '.'; i++ {
		} // 找到连续的 .

		var right = byte('R') // 右哨兵
		if i < n {
			right = bs[i] // 在范围内会替换
		}

		if left == right {
			for j = l; j < i; j++ { // 相同的符号, 批量替换
				bs[j] = right
			}
		} else if left == 'R' && right == 'L' {
			var k int // 左边→, 右边←, 两边替换中间怼
			for j, k = l, i-1; j < k; j, k = j+1, k-1 {
				bs[j] = 'R'
				bs[k] = 'L'
			}
		}
		l = i + 1 // 交换下一步的指针
		left = right
	}

	return string(bs)
}
