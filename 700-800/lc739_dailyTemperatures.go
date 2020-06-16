package _00_800

func dailyTemperatures(T []int) []int {
	// 从后向前走,

	// 很明显, 最后一项是 0
	// 对于 i而言,
	// 如果 T[i] <= T[i+1], 那么 R[i] = 1
	// 如果 T[i] > T[i+1]
	// 		如果 T[i+1] == 0, 那么 R[i] = 0
	// 		如果 T[i+1] != 0, 那么就比较 T[i] 和 T[i+1+R[i+1]]
	//			T[]

	ln := len(T)
	if 1 > ln {
		return nil
	}
	res := make([]int, len(T))
	res[ln-1] = 0
	for i := ln - 2; i >= 0; i-- {
		for t := i + 1; t < ln; t += res[t] {
			if T[i] < T[t] {
				res[i] = t - i
				break
			}
			if res[t] == 0 {
				res[i] = 0
				break
			}
		}
	}

	return res
}
