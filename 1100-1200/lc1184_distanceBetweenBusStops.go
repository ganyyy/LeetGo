package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	// 单向的, 貌似都不用考虑图论, 要么左到右, 要么右到左

	// l: 向左走,
	// r: 向右走.
	var l, r int
	var ln = len(distance)

	for cur := start; cur != destination; {
		cur = (cur - 1 + ln) % ln // 逆着走, 要先看前边的
		l += distance[cur]
	}
	for cur := start; cur != destination; {
		r += distance[cur] // 顺着走, 要先加当前的
		cur = (cur + 1) % ln
	}

	// fmt.Println(l, r)

	if l > r {
		return r
	}
	return l
}
