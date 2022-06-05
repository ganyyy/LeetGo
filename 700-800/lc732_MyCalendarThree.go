package main

import "sort"

type Pair struct {
	Day int
	Cnt int
}

type MyCalendarThree struct {
	calender []Pair
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{}
}

func (t *MyCalendarThree) Book(start int, end int) int {
	var sd = t.search(start)
	sd.Cnt++
	var ed = t.search(end)
	ed.Cnt--

	var cur int
	var ret int
	for _, p := range t.calender {
		cur += p.Cnt
		if cur > ret {
			ret = cur
		}
	}
	return ret

}

func (t *MyCalendarThree) search(day int) *Pair {
	var idx = sort.Search(len(t.calender), func(i int) bool {
		return t.calender[i].Day >= day
	})
	if idx == len(t.calender) {
		t.calender = append(t.calender, Pair{Day: day})
		return &t.calender[len(t.calender)-1]
	} else {
		if t.calender[idx].Day != day {
			t.calender = append(t.calender, Pair{})
			copy(t.calender[idx+1:], t.calender[idx:])
			t.calender[idx] = Pair{Day: day}
		}
		return &t.calender[idx]
	}
}

func main() {
	var c = Constructor()
	c.Book(5, 10)
	c.Book(50, 60)
	c.Book(10, 40)
	c.Book(5, 15)
	c.Book(5, 10)
	c.Book(25, 55)
}
