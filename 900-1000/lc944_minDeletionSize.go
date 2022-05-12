package main

func minDeletionSize(strs []string) int {
	var n = len(strs)
	var m = len(strs[0])
	if n <= 1 {
		return 0
	}

	if m < 1 {
		return 0
	}

	var ret int
	for i := 0; i < m; i++ {
		var begin = strs[0][i]

		for j := 1; j < n; j++ {
			if begin > strs[j][i] {
				ret++
				break
			}
			begin = strs[j][i]
		}

	}

	return ret
}
