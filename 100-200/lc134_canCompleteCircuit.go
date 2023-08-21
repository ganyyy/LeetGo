package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	// 首先, 可以保证车子能够正常往返的必要条件是 sum(gas) >sum(cost)

	var rest, run, start int
	for i := range gas {
		run += gas[i] - cost[i]
		rest += gas[i] - cost[i]
		// 注意 cost[i]的意义: 从 i -> i+1 点所消耗的汽油数量
		// [0:i]均是大于0的, 但是无法越过 i点, 说明i点的消耗太大, 此时不管从 [0:i+1]中的任何一个点出发都不能走完全程
		// 此时需要以i为起点重新出发
		if run < 0 {
			// 以i+1为起点重新出发,
			// 清空走过的路
			run = 0
			start = i + 1
		}
	}

	// 可以将整个数组抽象的分为两部分
	// [0:i], [i:len]
	// 如果整个路程中, gas的和大于cost, 那么 sum([0:i]) + sum[i:len])一定是大于0的
	// 可以将 i 理解为一个分界点, 其中一边可能小于0, 另一边一定大于0

	// 总油量是不够的
	if rest < 0 {
		return -1
	} else {
		// 总油量够了, 就看起点在什么地方
		return start
	}
}

// 错误的思路, 耗费大量时间且无法正确判断能否正确走完全程
func canCompleteCircuitFail(gas []int, cost []int) int {
	// 背包问题的变种
	// 但是起点位置不确定

	// 首先需要保证起点的油足够支撑到下一个地点
	var mark = make([]bool, len(cost))
	// 每个点都作为起点, 然后无脑遍历?

	var curGas int
	var check func(i int) bool
	check = func(i int) bool {
		var before = curGas
		mark[i] = true
		for k, m := range mark {
			if m {
				continue
			}
			if curGas += gas[k] - cost[k]; curGas >= 0 && check(k) {
				return true
			}
			curGas = before
		}
		mark[i] = false
		return false
	}

	for i := 0; i < len(gas); i++ {
		curGas = gas[i]
		if check(i) {
			return i
		}
	}
	return -1
}

func main() {
	/*
		[1,2,3,4,5]
		[3,4,5,1,2]
	*/
	var gas = []int{1, 2, 3, 4, 5}
	var cost = []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}
