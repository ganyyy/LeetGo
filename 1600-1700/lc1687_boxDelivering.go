package main

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	// n: 数量
	n := len(boxes)
	// p: port[i]
	p := make([]int, n+1)
	// w: weight[i]
	w := make([]int, n+1)
	// neg: 到i的时候, 中途需要变换几次方向
	neg := make([]int, n+1)
	// W: 所有箱子的前缀和
	W := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = boxes[i-1][0]
		w[i] = boxes[i-1][1]
		if i > 1 {
			neg[i] = neg[i-1]
			if p[i-1] != p[i] {
				neg[i]++
			}
		}
		W[i] = W[i-1] + w[i]
	}

	// 当前卡车中的箱子
	opt := []int{0}

	// f[i] = min{f[j]+neg(j+1,i)+2} // 从[j..i]中, 选取转向次数最少的方案
	//      = min{f[j]+neg[i]-neg[j+1]+2}, 即从[j..i]中, 需要更换几次方向
	// 搬运到第i个箱子, 需要多少趟
	f := make([]int, n+1)
	// g[j] = f[j]-neg[j+1]
	// f[i] = min{g[j]}+neg[i]+2
	g := make([]int, n+1)

	for i := 1; i <= n; i++ {
		// 随着i的递增, 会逐渐的更新f和g
		// 事实上, i直到首次不满足运输条件之前, opt会一直增加
		for i-opt[0] > maxBoxes || W[i]-W[opt[0]] > maxWeight {
			// 如果超过了可搬运箱子的数量/重量的上限, 就将首个箱子出队, 直到满足一次运输为止
			opt = opt[1:]
		}

		// 搬运到第i个箱子, 需要多少趟
		// 因为搬运的箱子都在opt中, 所以当前f[i]的值就是
		// f[opt[0]]+neg[i]-neg[opt[0]+1]+2
		f[i] = g[opt[0]] + neg[i] + 2

		if i != n {
			// 关键是这个f[i]-neg[i+1]
			// 可以理解为, 从 i 位置进行一次分割, 然后判断是不是会减小整体的转移次数
			g[i] = f[i] - neg[i+1]
			for len(opt) > 0 && g[i] <= g[opt[len(opt)-1]] {
				opt = opt[:len(opt)-1]
			}
			opt = append(opt, i)
		}
	}

	return f[n]
}
