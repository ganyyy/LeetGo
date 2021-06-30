package main

func numWays(n int, relation [][]int, k int) int {
	// dfs

	var cur, cnt int
	var next = make([][]int, n)
	for _, r := range relation {
		next[r[0]] = append(next[r[0]], r[1])
	}
	var dfs func(i int)

	dfs = func(i int) {
		if cur == k {
			if i == n-1 {
				cnt++
			}
		} else {
			for _, n := range next[i] {
				cur++
				dfs(n)
				cur--
			}
		}
	}
	dfs(0)
	return cnt
}
