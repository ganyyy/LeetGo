package main

func minimumOperations(num string) int {
	n := len(num)
	var found0, found5 bool
	for i := n - 1; i >= 0; i-- {
		c := num[i]
		// 首个找到的一定是最优解!
		if found0 && (c == '0' || c == '5') ||
			found5 && (c == '2' || c == '7') {
			return n - i - 2
		}
		if c == '0' {
			found0 = true
		} else if c == '5' {
			found5 = true
		}
	}
	if found0 {
		return n - 1
	}
	return n
}
