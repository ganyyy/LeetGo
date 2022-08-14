package main

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	const (
		START = "start"
		END   = "end"
	)

	var parse = func(log string) (id int, isStart bool, t int) {
		var sp = strings.Split(log, ":")
		id, _ = strconv.Atoi(sp[0])
		t, _ = strconv.Atoi(sp[2])
		isStart = sp[1] == START

		return
	}

	var cnt = make([]int, n)
	var stack [][2]int // 函数标识
	for _, log := range logs {
		id, isStart, t := parse(log)
		if isStart {
			stack = append(stack, [2]int{id, t})
		} else if len(stack) > 0 {
			var top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 当前函数的执行时间
			var interval = t - top[1] + 1
			cnt[top[0]] += interval
			if len(stack) > 0 {
				// 去掉父函数中额外增加的时间
				// 因为t会一直增加, 所以这里需要预先扣除
				cnt[stack[len(stack)-1][0]] -= interval
			}
		}
		// fmt.Println(log, cnt)
	}
	return cnt
}
