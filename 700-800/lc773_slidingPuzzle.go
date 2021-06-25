package main

const (
	result = ((((1<<3+2)<<3+3)<<3+4)<<3+5)<<3 + 0
)

func slidingPuzzle(board [][]int) int {
	// BFS

	// 一种枝减方式:

	// 根据0所处的位置找到所有可以移动的地方

	// 直接改成整数进行运算

	var tmp = []int{1, 2, 3, 4, 5, 0}
	var getVal = func() int {
		return ((((tmp[0]<<3+tmp[1])<<3+tmp[2])<<3+tmp[3])<<3+tmp[4])<<3 + tmp[5]
	}

	var result = getVal()
	var cur int
	for _, row := range board {
		for _, v := range row {
			cur = cur<<3 + v
		}
	}

	if cur == result {
		return 0
	}

	var parseTmp = func(cur int) {
		tmp[0] = (cur & (7 << 15)) >> 15
		tmp[1] = (cur & (7 << 12)) >> 12
		tmp[2] = (cur & (7 << 9)) >> 9
		tmp[3] = (cur & (7 << 6)) >> 6
		tmp[4] = (cur & (7 << 3)) >> 3
		tmp[5] = (cur & (7 << 0)) >> 0
	}

	var getNext = func(cur int) []int {
		parseTmp(cur)
		//fmt.Println(tmp)

		// 找到0, 并获取所有可以转换的数字, 返回
		var ret []int
		if tmp[0] == 0 || tmp[3] == 0 {
			ret = make([]int, 0, 2)
			var base = 0
			var pair = 3
			if tmp[3] == 0 {
				base, pair = pair, base
			}
			tmp[base], tmp[pair] = tmp[pair], tmp[base]
			ret = append(ret, getVal())
			tmp[base], tmp[pair] = tmp[pair], tmp[base]

			tmp[base], tmp[base+1] = tmp[base+1], tmp[base]
			ret = append(ret, getVal())
			tmp[base], tmp[base+1] = tmp[base+1], tmp[base]

		} else if tmp[2] == 0 || tmp[5] == 0 {
			ret = make([]int, 0, 2)
			var base = 2
			var pair = 5
			if tmp[5] == 0 {
				base, pair = pair, base
			}
			tmp[base], tmp[pair] = tmp[pair], tmp[base]
			ret = append(ret, getVal())
			tmp[base], tmp[pair] = tmp[pair], tmp[base]

			tmp[base], tmp[base-1] = tmp[base-1], tmp[base]
			ret = append(ret, getVal())
			tmp[base], tmp[base-1] = tmp[base-1], tmp[base]
		} else {
			// 1, 4 为0, 有三种结果

			ret = make([]int, 0, 3)
			var base = 1
			var pair = 4
			if tmp[4] == 0 {
				base, pair = pair, base
			}

			tmp[base], tmp[base-1] = tmp[base-1], tmp[base]
			ret = append(ret, getVal())
			tmp[base], tmp[base-1] = tmp[base-1], tmp[base]

			tmp[base], tmp[base+1] = tmp[base+1], tmp[base]
			ret = append(ret, getVal())
			tmp[base], tmp[base+1] = tmp[base+1], tmp[base]

			tmp[base], tmp[pair] = tmp[pair], tmp[base]
			ret = append(ret, getVal())
			tmp[base], tmp[pair] = tmp[pair], tmp[base]
		}

		return ret
	}

	var visited = make(map[int]bool)

	var queue1, queue2 []int

	queue1 = []int{cur}

	var step int
	for len(queue1) != 0 {
		for _, c := range queue1 {
			if c == result {
				return step
			}
			visited[c] = true
			for _, n := range getNext(c) {
				//parseTmp(n)
				//fmt.Println(c, tmp)
				if !visited[n] {
					queue2 = append(queue2, n)
				}
			}
		}
		step++
		queue1, queue2 = queue2, queue1
		queue2 = queue2[:0]
	}

	return -1
}

// 指定可选的位置数组
var nei = [][]int{
	{1, 3},
	{0, 2, 4},
	{1, 5},
	{0, 4},
	{1, 3, 5},
	{2, 4},
}

func slidingPuzzle2(board [][]int) int {
	// BFS

	// 一种枝减方式:

	// 根据0所处的位置找到所有可以移动的地方

	// 直接改成整数进行运算

	var tmp = []int{1, 2, 3, 4, 5, 0}
	var getVal = func() int {
		return ((((tmp[0]<<3+tmp[1])<<3+tmp[2])<<3+tmp[3])<<3+tmp[4])<<3 + tmp[5]
	}

	var result = getVal()
	var cur int
	for _, row := range board {
		for _, v := range row {
			cur = cur<<3 + v
		}
	}

	if cur == result {
		return 0
	}

	var parseTmp = func(cur int) {
		tmp[0] = (cur & (7 << 15)) >> 15
		tmp[1] = (cur & (7 << 12)) >> 12
		tmp[2] = (cur & (7 << 9)) >> 9
		tmp[3] = (cur & (7 << 6)) >> 6
		tmp[4] = (cur & (7 << 3)) >> 3
		tmp[5] = (cur & (7 << 0)) >> 0
	}

	var getNext = func(cur int) []int {
		parseTmp(cur)
		//fmt.Println(tmp)

		// 找到0, 并获取所有可以转换的数字, 返回
		var ret []int
		for i, v := range tmp {
			if v == 0 {
				ret = make([]int, 0, len(nei[i]))
				for _, n := range nei[i] {
					tmp[i], tmp[n] = tmp[n], tmp[i]
					ret = append(ret, getVal())
					tmp[i], tmp[n] = tmp[n], tmp[i]
				}
				break
			}
		}
		return ret
	}

	var visited = make(map[int]bool)

	var queue1, queue2 []int

	queue1 = []int{cur}

	var step int
	for len(queue1) != 0 {
		for _, c := range queue1 {
			if c == result {
				return step
			}
			visited[c] = true
			for _, n := range getNext(c) {
				if !visited[n] {
					queue2 = append(queue2, n)
				}
			}
		}
		step++
		queue1, queue2 = queue2, queue1
		queue2 = queue2[:0]
	}

	return -1
}

func main() {
	println(slidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}}))
}
