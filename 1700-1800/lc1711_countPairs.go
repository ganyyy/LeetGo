package main

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
)

const (
	BIT = 22
	MOD = 1e9 + 7
)

func countPairs(deliciousness []int) int {
	// 统计每个数字的数量, 进行一次压缩

	var m = make(map[int]int)

	for _, v := range deliciousness {
		m[v]++
	}

	var visited = make(map[int]map[int]bool)
	for v := range m {
		visited[v] = map[int]bool{}
	}

	// 两两相加, 那么最大值就是 2^40 = 1 << 41
	var cnt int

	for v, n := range m {
		for i := 0; i < BIT; i++ {
			var cur = 1 << i
			if v > cur {
				continue
			}
			var sub = cur - v
			if visited[v][sub] {
				continue
			}
			if n2 := m[sub]; n2 != 0 {
				if sub == v {
					n2 = (n * (n - 1)) / 2
				} else {
					n2 *= n
				}
				cnt = (cnt + n2) % MOD
				visited[sub][v] = true
			}
		}
	}

	fmt.Println(m, visited)

	return cnt
}

func countPairs2(deliciousness []int) int {
	num := len(deliciousness)
	sort.Ints(deliciousness)

	numMap := map[int]int{}
	prevCnt := 0
	numMap[deliciousness[0]] = 1
	for i := 1; i < num; i++ {
		n := deliciousness[i]
		// 这个函数用来返回首个1所处的索引. 0的话是32

		// n0 表示的是 n <= 最近的二的整数次幂
		n0 := 1<<(32-bits.LeadingZeros32(uint32(n))) - n
		// 简单的来说, 1<->1, 3<->1, 7<->1, 5<->3
		// 肯定能配对上. 只要存在的话. 这样就避免了计算是否标记的问题
		cnt := numMap[n0]

		// 这里是当n恰好为 二的整数次幂的时候, 就加上0的数量
		// 前提是本身不能为0
		if deliciousness[0] == 0 && (n > 0 && (n&(n-1)) == 0) {
			cnt += numMap[0]
		}
		prevCnt = (prevCnt + cnt) % (1e9 + 7)
		// 当前数字的计数+1
		numMap[n] += 1
	}
	//fmt.Printf("maps:%+v\n", numMap)
	return prevCnt
}

func main() {
	fmt.Println(bits.LeadingZeros32(uint32(0)))
	fmt.Println(bits.LeadingZeros32(uint32(1)))
	fmt.Println(bits.LeadingZeros32(uint32(3)))
	fmt.Println(bits.LeadingZeros32(uint32(7)))
	fmt.Println(bits.LeadingZeros32(uint32(math.MaxUint32)))
}
