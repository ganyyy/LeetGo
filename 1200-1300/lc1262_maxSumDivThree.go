//go:build ignore

package main

func maxSumDivThree(nums []int) int {
	// 上一个状态, sum mod 3 = 0/1/2 对应的最大值
	f := []int{0, -0x3f3f3f3f, -0x3f3f3f3f}
	g := make([]int, 3)
	for _, num := range nums {
		// 针对当前这个数...
		// 比如num = 2
		// [i,mod] = {[0,2],[1,0],[2,1]}
		// 因为 2 % 3 == 2, 所以 可以和 x % 3 == 1的最大和合并到 当前状态的0上
		// 但是, 当前状态0可以选择使用该num(f[1]+num), 也可以不选择使用该num(f[0])
		// 同时, 当前数字的选取与否会对下一个状态产生影响
		for i := 0; i < 3; i++ {
			mod := (i + num) % 3
			// f[mod]: 不选当前位
			// f[i]+num: 选取当前位
			g[mod] = max(f[mod], f[i]+num)
		}
		f, g = g, f
	}
	// 最后返回的就是 mod 3 == 0的最大可能的集合
	return f[0]
}
