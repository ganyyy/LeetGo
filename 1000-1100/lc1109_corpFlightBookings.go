package main

func corpFlightBookingsGood(bookings [][]int, n int) []int {
	prefix := make([]int, n+1)
	answer := make([]int, n+1)
	for _, booking := range bookings {
		prefix[booking[0]-1] += booking[2]
		prefix[booking[1]] -= booking[2]
	}
	sum := 0
	for i, v := range prefix {
		sum += v
		answer[i] = sum
	}
	return answer[:len(answer)-1]
}

func corpFlightBookingsMedium(bookings [][]int, n int) []int {
	// 果然要用区间和

	var start, end = make(map[int]int), make(map[int]int)

	for _, b := range bookings {
		start[b[0]] += b[2]
		end[b[1]] += b[2]
	}

	var ans = make([]int, n)
	var tmp int
	for i := 0; i < n; i++ {
		tmp += start[i+1]
		ans[i] = tmp
		tmp -= end[i+1]
	}
	return ans
}

func corpFlightBookingsBad(bookings [][]int, n int) []int {
	var ret = make([]int, n+1)

	for _, booking := range bookings {
		var from, to, seats = booking[0], booking[1], booking[2]
		for i := from; i <= to; i++ {
			ret[i] += seats
		}
	}

	return ret[1:]
}
