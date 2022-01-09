package main

import "math"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	for d := 2; d*d <= num; d++ {
		// 找到所有因数
		if num%d == 0 {
			sum += d
			// 因为 d ∈ [2, sqrt(num)],
			if d*d < num {
				sum += num / d
			}
		}
	}
	return sum == num

}

func checkPerfectNumberYYDS(num int) bool {
	switch num {
	case 6, 28, 496, 8128, 33550336:
		return true
	default:
		return false
	}
}

func checkPerfectNumberNew(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	var limit = int(math.Sqrt(float64(num)))
	for d := 2; d <= limit; d++ {
		if num%d == 0 {
			// 先加上小于 limit 的因数
			sum += d
			if d*d < num {
				// 再加上大于limit的因数
				sum += num / d
			}
		}
	}
	return sum == num
}
