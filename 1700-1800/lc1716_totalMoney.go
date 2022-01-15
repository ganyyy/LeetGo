package main

const WEEKDAY = 1 + 2 + 3 + 4 + 5 + 6 + 7

func totalMoney(n int) int {
	var week = n / 7
	var day = n % 7

	var weekDay int
	var addDay int
	if week > 0 {
		weekDay = WEEKDAY*week + week*(week-1)/2*7
		addDay = (week) * day
	}
	if day == 0 {
		return weekDay
	}

	return weekDay + day*(day+1)/2 + addDay
}
