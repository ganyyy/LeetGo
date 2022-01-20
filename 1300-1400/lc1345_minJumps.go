package main

import "fmt"

type set map[int]struct{}

func (s set) add(v int) {
	s[v] = struct{}{}
}

func (s set) has(v int) bool {
	var _, ok = s[v]
	return ok
}

func zip(arr []int) []int {
	// 压缩算法有问题. 多个循环节无法处理
	var preIdx int
	var start int
	for i, v := range arr {
		if v == arr[preIdx] {
			continue
		}
		// 中间差了超过2个位置
		preIdx++
		if i > start+1 {
			arr[preIdx] = arr[preIdx-1]
			preIdx++
		}
		arr[preIdx] = v
		start = i
	}

	if len(arr) > start+1 {
		preIdx++
	}

	return arr[:preIdx+1]
}

func minJumps(arr []int) int {
	arr = zip(arr)
	if len(arr) == 1 {
		return 0
	}

	var next = make([]set, len(arr))
	// 这是一个图, 当前节点可以到达的所有位置的集合

	var same = make(map[int][]int)

	for i, v := range arr {
		next[i] = make(map[int]struct{}, 2)
		if i > 1 {
			next[i].add(i - 1)
		}
		if i < len(arr)-1 {
			next[i].add(i + 1)
		}
		same[v] = append(same[v], i)
	}

	for _, v := range same {
		for _, idx := range v {
			var n = next[idx]
			for _, i := range v {
				if i == idx {
					continue
				}
				// 需要压缩一下
				n.add(i)
			}
		}
	}

	var visited set = make(map[int]struct{}, len(arr))

	var src = []int{0}
	var end = len(arr) - 1
	var cur int
	for len(src) != 0 {
		cur++
		var idx = len(src)
		for _, i := range src[:idx] {
			if visited.has(i) {
				continue
			}
			visited.add(i)
			var n = next[i]
			if n.has(end) {
				return cur
			}
			for nn := range n {
				if visited.has(nn) {
					continue
				}
				src = append(src, nn)
			}
		}
		src = src[idx:]
	}
	return 1
}

func minJumpsGood(arr []int) int {
	var ln = len(arr)
	var m = make(map[int][]int, ln)

	// 按照值进行划分位置
	for i, v := range arr {
		var n = m[v]
		if n == nil {
			n = make([]int, 0, 4)
		}
		m[v] = append(n, i)
	}

	// fmt.Println(m)

	// 所有访问过的位置
	var visited = make([]bool, ln)
	visited[0] = true
	// 当前方位的位置和步数
	var queue = [][2]int{{0, 0}}

	for len(queue) > 0 {
		var top = queue[0]
		queue = queue[1:]

		// BFS迭代的位置和步数
		var idx, step = top[0], top[1]
		// 可选择的下一步的位置
		var next = m[arr[idx]]

		// 如果到达了终点...
		if idx == ln-1 {
			return step
		}

		// 增加idx+1,idx-1
		if idx > 0 && !visited[idx-1] {
			next = append(next, idx-1)
		}
		if idx < ln-1 && !visited[idx+1] {
			next = append(next, idx+1)
		}
		for _, p := range next {
			if visited[p] {
				continue
			}
			visited[p] = true
			// 加入到备选节点中
			queue = append(queue, [2]int{p, step + 1})
		}

		// 必须要清空!
		// 因为同样的值不需要迭代第二遍!
		m[arr[idx]] = nil
	}

	return -1
}

func main() {
	var src = []int{1, 1, 1, 1, 1, 2, 2, 3, 3, 3, 4}

	fmt.Println(zip(src))
}
