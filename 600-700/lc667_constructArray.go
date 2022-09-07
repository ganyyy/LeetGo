package main

func constructArray(n, k int) []int {
	ans := make([]int, 0, n)
	// n-k个1
	for i := 1; i < n-k; i++ {
		// i1, i2, i3
		// |i1-i2| = 1
		// |i2-i3| = 1
		ans = append(ans, i)
	}

	// k个差值
	for i, j := n-k, n; i <= j; i++ {
		// i1, j1, i2=i1+1, j2=j1-1
		// |i1-j1| = k
		// |j1-i2| = k-1
		// |i2-j2| = k-2
		ans = append(ans, i)
		if i != j {
			ans = append(ans, j)
		}
		j--
	}
	return ans
}
