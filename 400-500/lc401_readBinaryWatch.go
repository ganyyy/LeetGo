package main

import "fmt"

// 性能极差版
func readBinaryWatch2(turnedOn int) []string {
	// 最高同时亮8个灯(11:59)

	if turnedOn >= 9 {
		return nil
	}

	if turnedOn == 0 {
		return []string{"0:00"}
	}

	var ret []string

	// 这是一道递归题目
	var dfs func(on int)
	// [0:6]表示分钟, [6:10] 表示小时
	var val int

	var m = map[string]bool{}

	dfs = func(on int) {
		if on > turnedOn {
			// 根据当前已选择的数字组合成一个答案
			var minute, hour int
			for i := 0; i < 6; i++ {
				if val&(1<<i) != 0 {
					minute += 1 << i
				}
			}
			if minute > 59 {
				return
			}
			for i := 6; i < 10; i++ {
				if val&(1<<i) != 0 {
					hour += 1 << (i - 6)
				}
			}
			if hour > 11 {
				return
			}
			var str = fmt.Sprintf("%d:%02d", hour, minute)
			if m[str] {
				return
			}
			m[str] = true
			ret = append(ret, str)
			return
		}
		for i := 0; i < 10; i++ {
			if val&(1<<i) == 0 {
				val |= 1 << i
				dfs(on + 1)
				val &^= 1 << i
			}
		}
	}

	dfs(1)

	return ret
}

func readBinaryWatch(turnedOn int) []string {
	// 暴力穷举版, 这个性能很好

	if turnedOn >= 9 {
		return nil
	}
	if turnedOn == 0 {
		return []string{"0:00"}
	}

	var res []string

	var count = func(i int) int {
		var cnt int
		for i != 0 {
			i &= i - 1
			cnt++
		}
		return cnt
	}

	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if count(i)+count(j) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}

	return res
}
