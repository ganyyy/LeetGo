package main

import "math"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))

	var max int = math.MinInt32
	for _, v := range candies {
		if max < v {
			max = v
		}
	}
	for i, v := range candies {
		if v+extraCandies >= max {
			res[i] = true
		}
	}
	return res
}

func main() {

}
