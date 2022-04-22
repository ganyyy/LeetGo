package main

func maxRotateFunction(nums []int) int {
	var sum int
	var pre int
	for i, v := range nums {
		sum += v
		pre += i * v
	}

	var max = pre
	for i := 1; i < len(nums); i++ {
		// F(0) = 0*A0+1*A1+2*A2+3*A3+4*A4
		// F(1) =      0*A1+1*A2+2*A3+3*A4+4*A0
		// F(2) =           0*A2+1*A3+2*A4+3*A0+4*A1
		// F(0) - F(1) = A1+A2+A3+A4 - 4*A0 = sum - 5*A0
		// F(1) - F(2) = A2+A3+A4+A0 - 4*A1 = sum - 5*A1
		pre = pre - sum + len(nums)*nums[i-1]
		if pre > max {
			max = pre
		}
	}

	return max
}
