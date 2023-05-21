//go:build ignore

package main

func maxEqualRowsAfterFlips(matrix [][]int) (ans int) {
	// 由于只有0和1, 所以可以使用异或来判断是否相等
	// 核心是: 取每一行的第一个元素, 然后异或后, 将结果放到一个数组中(相当于首位是1的时候, 才会发生异或)
	// 将这个数组作为key, 然后统计出现的次数, 最后取最大值即可
	cnt := map[[5]uint64]int{}
	for _, row := range matrix {
		r := [5]uint64{}
		for i, x := range row {
			r[i/64] |= uint64(x^row[0]) << (i % 64)
		}
		cnt[r]++
	}
	for _, c := range cnt {
		ans = max(ans, c)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
