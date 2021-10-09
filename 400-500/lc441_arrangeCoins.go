package main

import "math"

func arrangeCoins(n int) int {
	// row * (row + 1)/2 >= n

	// row*(row+1) >= n/2

	// row^2+row - n/2 >= 0

	return int(math.Floor(math.Sqrt(1+float64(8*n))-1) / 2)
}
