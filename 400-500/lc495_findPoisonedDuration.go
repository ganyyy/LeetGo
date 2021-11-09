package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	var next int
	var cnt int
	for i := 0; i < len(timeSeries); i++ {
		if next > timeSeries[i] && i > 0 {
			cnt += timeSeries[i] - timeSeries[i-1]
		} else {
			cnt += duration
		}
		next = timeSeries[i] + duration
	}

	return cnt
}
