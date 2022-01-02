package main

import (
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	var t = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return t.Weekday().String()
}

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = func() []int {
	var t = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	for i := 1; i < len(t); i++ {
		t[i] += t[i-1]
	}
	return t
}()

func dayOfTheWeekGood(day, month, year int) string {
	days := 0
	// 输入年份之前的年份的天数贡献
	// 这里采用了一个取巧的方法: 1968年是瑞年. 往后每4年都是瑞年
	// 所以可以用(year-1968)/4 计算距离当前年份中, 瑞年的个数
	days += 365*(year-1971) + (year-1969)/4
	// 输入年份中，输入月份之前的月份的天数贡献
	day += monthDays[month-1]
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		days++
	}
	// 输入月份中的天数贡献
	days += day
	return week[(days+3)%7]
}

func main() {
	println(dayOfTheWeek(3, 1, 2022))
	println(dayOfTheWeekGood(3, 1, 2022))
}
