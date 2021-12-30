package main

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	for d := 2; d*d <= num; d++ {
		if num%d == 0 {
			sum += d
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
