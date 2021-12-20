package main

import (
	"strconv"
	"strings"
	"time"
)

func dayOfYearBad(date string) int {

	var t, _ = time.Parse("2006-01-02", date)

	return t.YearDay()
}

func dayOfYear(date string) (ret int) {
	// 事实证明: scanf 性能差的可以. 甚至不如split
	var year, month, day int
	var str = strings.Split(date, "-")
	year, _ = strconv.Atoi(str[0])
	month, _ = strconv.Atoi(str[1])
	day, _ = strconv.Atoi(str[2])

	ret = monthDay[month-1] + day
	if isLipYear(year) && month > 2 {
		ret += 1
	}

	return
}
func isLipYear(year int) bool {
	return (year%100 != 0 && year%4 == 0) || (year%400 == 0)
}

var monthDay = [12]int{
	0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334,
}

func main() {

}
