//go:build ignore

package main

func longestArithSeqLength(a []int) (ans int) {
	n := len(a)
	// [0, 500]之间,  差值的取值范围在 [-500, 500] 共计 1001 个数字
	f := make([][1001]int, n)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- { // 倒序迭代
			d := a[i] - a[j] + 500 // +500 防止出现负数
			if f[i][d] == 0 {      // 为啥每个公差只计算一次呢?
				// 可以理解为: 尽量选最右边的. 这样才能保证在公差为0的情况下, 可以选取最多的子数组
				f[i][d] = f[j][d] + 1 // 默认的 1 (自己本身) 在下面返回时加上
				ans = max(ans, f[i][d])
			}
		}
	}
	return ans + 1
}
