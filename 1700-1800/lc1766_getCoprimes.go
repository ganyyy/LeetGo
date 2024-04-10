package main

var GCD = func() [51][]int {
	var ret [51][]int
	ret[1] = []int{1}
	for i := 1; i <= 50; i++ {
		for j := 1; j <= 50; j++ {
			if i == j {
				continue
			}
			if gcd(i, j) == 1 {
				// 互质
				ret[i] = append(ret[i], j)
			}
		}
	}
	return ret
}()

func gcd(a, b int) int {
	if a == 0 || b == 0 {
		return a + b
	}
	if a > b {
		a, b = b, a
	}
	for b%a != 0 {
		a, b = b%a, a
	}
	return a
}

func getCoprimes(nums []int, edges [][]int) []int {
	nexts := make([][]int, len(nums))
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		nexts[a] = append(nexts[a], b)
		nexts[b] = append(nexts[b], a)
	}

	var ret = make([]int, len(nums))
	var dep = make([]int, len(nums)) // 深度
	var pos [51]int                  // 每位数字对应的节点id. 因为是DFS, 所以需要保留现场并进行回溯
	for i := range pos {
		pos[i] = -1
	}
	for i := range ret {
		ret[i] = -1
	}
	copy(dep, ret)
	dep[0] = 0

	var dfs func(current, parent int)
	dfs = func(current, parent int) {

		// 找到最近的互质数
		var currentVal = nums[current]
		for _, g := range GCD[currentVal] {
			if pos[g] == -1 {
				// 还未出现
				continue
			}
			// 如果当前节点还未赋值, 或者出现了一个距离current更近的互质节点
			if ret[current] == -1 || dep[ret[current]] < dep[pos[g]] {
				ret[current] = pos[g]
			}
		}
		// 保留现场
		var old = pos[currentVal]
		pos[currentVal] = current
		nextDep := dep[current] + 1
		for _, next := range nexts[current] {
			if next == parent {
				continue
			}
			dep[next] = nextDep
			dfs(next, current)
		}
		// 回溯
		pos[currentVal] = old
	}

	dfs(0, -1)

	return ret
}
