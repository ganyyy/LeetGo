package main

import "fmt"

func shipWithinDaysTimeout(weights []int, D int) int {
	// 毫无意外的时间超了...

	// 一个一个测试 不同的载重?

	// 首先, 最小值应该大于所有货物的最大值

	var minLoad int
	var sum int
	for i, v := range weights {
		sum += v
		weights[i] = sum
		if minLoad < v {
			minLoad = v
		}
	}

	// 暴力穷举, 就是从最小值开始, 一直遍历到 d == D 时

	var load int
	var nextAdd int
	for load = minLoad; ; load += nextAdd {
		var pre int
		var d int
		var i int
		for i = 0; i < len(weights); {
			fmt.Println(weights[i], pre, load)
			if weights[i]-pre > load {
				pre = weights[i-1]
				if d == 0 {
					nextAdd = weights[i]
				}
				d++
				if d == D {
					break
				}
				continue
			}
			i++
		}
		if i == len(weights) {
			break
		}
	}

	return load
}

func shipWithinDays(weights []int, D int) int {
	// 我日, 竟然要用二分...

	// 修改原始数组, 计算前缀和
	var sum, max int
	for i, v := range weights {
		sum += v
		weights[i] = sum
		if max < v {
			max = v
		}
	}
	// 二分查找欸..

	var checkDay = func(load int) bool {
		var day int
		var cur int
		var idx int
		for idx = 0; idx < len(weights); {
			if weights[idx]-cur > load {
				cur = weights[idx-1]
				day++
				if day == D {
					break
				}
				continue
			}
			idx++
		}

		return idx == len(weights)
	}

	// 起点为 最多的运载天数
	// 终点为 一天运过去, 即存放所有的货物
	for max < sum {
		var load = max + (sum-max)/2
		if checkDay(load) {
			// 如果能运过去, 就尝试缩小 有边界
			sum = load
		} else {
			max = load + 1
		}
	}

	return max
}

func main() {
	println(shipWithinDays([]int{500, 500, 500, 500, 500, 500, 500, 500, 500}, 1))
}
