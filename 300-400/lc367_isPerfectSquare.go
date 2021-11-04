package main

func isPerfectSquare(num int) bool {
	// 二分走起

	var low, high = 1, num

	for low < high {
		var mid = low + (high-low)/2
		var mul = mid * mid
		if mul == num {
			return true
		} else if mul > num {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return false

}

func main() {
	isPerfectSquare(14)
}
