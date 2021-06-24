package main

import (
	"strconv"
	"unsafe"
)

func openLock(deadends []string, target string) int {
	// 好家伙, 开始BFS了..
	var dead = make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}

	// 包括目标或者 初始值的情况直接返回-1
	if dead[target] || dead["0000"] {
		return -1
	}

	// 快速转换
	var toString = func(bs []byte) string {
		return *(*string)(unsafe.Pointer(&bs))
	}

	// 获取切片的头部元素
	//var getTop = func(bss [][]byte) []byte {
	//	var top = bss[0]
	//	bss = bss[1:]
	//	return top
	//}

	// 获取到达当前值的所有可能
	var getNext = func(cur []byte) [][]byte {
		var ret = make([][]byte, 0, 8)

		for i := 0; i < 4; i++ {
			// TA, TB 分别表示到达当前值所需要的-1/+1
			var tmpA = make([]byte, 4)
			var tmpB = make([]byte, 4)
			copy(tmpA, cur)
			copy(tmpB, cur)
			if cur[i] == '0' {
				tmpA[i] = '9'
			} else {
				tmpA[i] = cur[i] - 1
			}

			if cur[i] == '9' {
				tmpB[i] = '0'
			} else {
				tmpB[i] = cur[i] + 1
			}

			ret = append(ret, tmpA)
			ret = append(ret, tmpB)
		}

		return ret
	}

	// 两个遍历方向
	var queue1, queue2 [][]byte

	// 已访问的节点
	var visited = make(map[string]bool)

	// 起始值
	var start = []byte("0000")

	queue1 = append(queue1, start)

	var step int
	for len(queue1) != 0 {
		var top = queue1[0]
		queue1 = queue1[1:]
		if toString(top) == target {
			return step
		}

		var next = getNext(top)

		for _, bs := range next {
			if !dead[toString(bs)] && !visited[toString(bs)] {
				visited[toString(bs)] = true
				queue2 = append(queue2, bs)
			}
		}

		if len(queue1) == 0 {
			queue1, queue2 = queue2, queue1
			queue2 = queue2[:0]
			step++
		}
	}

	return -1
}

func openLock2(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	var deadArray [10000]bool
	for _, s := range deadends {
		i, _ := strconv.Atoi(s)
		deadArray[i] = true
	}
	targetInt, _ := strconv.Atoi(target)
	if deadArray[0] {
		return -1
	}
	q := []int{0}
	visit := [10000]bool{0: true}
	var tq []int
	helper := func(out []int, x, rate int) {
		if k := (x / rate) % 10; k == 0 {
			out[0] = x + 9*rate
			out[1] = x + 1*rate
		} else if k == 9 {
			out[0] = x - 9*rate
			out[1] = x - 1*rate
		} else {
			out[0] = x + 1*rate
			out[1] = x - 1*rate
		}
	}
	res := 0
	for len(q) > 0 {
		res++
		for len(q) > 0 {
			n := q[len(q)-1]
			q = q[:len(q)-1]

			var next [8]int
			helper(next[:], n, 1)
			helper(next[2:], n, 10)
			helper(next[4:], n, 100)
			helper(next[6:], n, 1000)
			for _, v := range next {
				if v == targetInt {
					return res
				}
				if !deadArray[v] && !visit[v] {
					tq = append(tq, v)
					visit[v] = true
				}
			}
		}
		q, tq = tq, q
		tq = tq[:0]
	}
	return -1
}

func main() {
	println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
