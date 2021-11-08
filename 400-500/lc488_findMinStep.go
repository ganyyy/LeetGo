package main

import "math"

func ParseIndex(b byte) int {
	switch b {
	case 'R':
		return 1
	case 'Y':
		return 2
	case 'B':
		return 3
	case 'G':
		return 4
	case 'W':
		return 0
	}
	return -1
}

func ParseColor(i int) byte {
	switch i {
	case 1:
		return 'R'
	case 2:
		return 'Y'
	case 3:
		return 'B'
	case 4:
		return 'G'
	case 0:
		return 'W'
	}
	return '-'
}

func findMinStep(board string, hand string) int {
	var set [5]int

	for i := range hand {
		set[ParseIndex(hand[i])]++
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var removeSrc = func(src []byte, start, end int) []byte {
		var tmp = make([]byte, len(src))
		copy(tmp, src)
		return append(tmp[:start], tmp[end+1:]...)
	}

	var result = math.MaxInt32

	var dfs func(src []byte, step int)

	dfs = func(src []byte, step int) {
		if step >= result {
			return
		}
		if len(src) == 0 {
			result = min(result, step)
			return
		}

		for i := range src {
			//1. 获取当前最长的连续字串
			var j = i
			var idx = ParseIndex(src[i])
			for j+1 < len(src) && src[j+1] == src[i] {
				j++
			}

			if j == i && set[idx] >= 2 {
				// 只有一个球, 并且手中有两个以上的备选球
				set[idx] -= 2
				dfs(Eliminate(removeSrc(src, i, j)), step+2)
				set[idx] += 2
			} else if j == i+1 {
				// 两个球
				if set[idx] >= 1 {
					var tmp = make([]byte, len(src))
					copy(tmp, src)
					set[idx] -= 1
					dfs(Eliminate(removeSrc(src, i, j)), step+1)
					set[idx] += 1
				}

				// 尝试向两个颜色相同的球中间插入一个颜色不同的球
				for c := 0; c < 5; c++ {
					if c == idx {
						continue
					}
					if set[c] <= 0 {
						continue
					}
					var tmp = make([]byte, 0, len(src)+1)
					tmp = append(tmp, src[:i+1]...)
					tmp = append(tmp, ParseColor(c))
					tmp = append(tmp, src[i+1:]...)
					set[c]--
					dfs(Eliminate(tmp), step+1)
					set[c]++
				}
			}
		}
	}

	dfs([]byte(board), 0)

	if result == math.MaxInt32 {
		return -1
	}
	return result
}

//Eliminate 去除所有三个以上的组合
func Eliminate(src []byte) []byte {
	var flag = true
	for flag {
		flag = false
		for i := 0; i < len(src); {
			j := i + 1
			for ; j < len(src) && src[j] == src[i]; j++ {
			}
			if j-i >= 3 {
				flag = true
				src = append(src[:i], src[j:]...)
			} else {
				i = j
			}
		}
	}
	return src
}

func main() {

	println(findMinStep("RRYGGYYRRYGGYYRR", "GGBBB"))
}
