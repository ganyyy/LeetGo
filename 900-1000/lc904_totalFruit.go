package main

func totalFruit(fruits []int) int {
	// 贪心? DP?

	const (
		DEFAULT = -1
	)

	// 最终结果
	var ret int

	// 起始位置
	var start int

	// 选取的两个类型
	fa, fb := DEFAULT, DEFAULT
	// 这两个类型最后出现的位置
	lastA, lastB := DEFAULT, DEFAULT
	// 两个类型最后连续出现的次数
	var seqA, seqB int

	for i, fruit := range fruits {
		if fa == DEFAULT {
			fa = fruit
		}
		if fa == fruit {
			lastA = i
			seqA++
			seqB = 0
			continue
		}
		if fb == DEFAULT {
			fb = fruit
		}
		if fb == fruit {
			seqB++
			seqA = 0
			lastB = i
			continue
		}
		if fruit == fa || fruit == fb {
			// 两种类型之一
			continue
		}
		// 出现了第三种类型
		// 通过lastA/lastB选取最接近的进行复用
		ret = max(i-start, ret)
		// fmt.Println("(",fa, fb,")","(",lastA, lastB,")","(",seqA, seqB,")", ret)
		if lastA < lastB {
			lastA = lastB
			fa = fb
			seqA = seqB
		}
		fb, lastB, seqB = fruit, i, 1
		start = lastA - seqA + 1 // 去掉本身占用的1个位置
		seqA = 0                 // 当前是B, A的长度要重置
		// fmt.Println(start)
	}

	return max(ret, len(fruits)-start)
}
