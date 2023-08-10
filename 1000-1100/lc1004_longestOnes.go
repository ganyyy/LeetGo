package main

import "fmt"

func longestOnes(A []int, K int) int {
	var l, r int
	var cur = K
	var res int
	for ; r < len(A); r++ {

		// 等于1不用处理
		if A[r] == 1 {
			// 此时结果包括当前位置
			res = max(res, r-l+1)
			continue
		}

		// 如果一个替换的都没有的话, 直接走就好了
		if K == 0 {
			// 所以为什么有时候是 r, 有时候是r+1呢?
			l = r + 1
			continue
		}

		// 有多余的就给它补上
		if cur > 0 {
			// 此时结果包括当前位置
			res = max(res, r-l+1)
			cur--
			continue
		}

		// 去掉第一个被替换的0, 放到当前位置上
		for ; l < r; l++ {
			if A[l] == 0 {
				l++
				break
			}
		}
		// 此时结果不包括当前位置
		res = max(res, r-l+1)
	}
	return res
}

func longestOnes2(A []int, K int) int {
	var l, r int
	var res int
	for ; r < len(A); r++ {
		// 等于1不用处理
		if A[r] == 1 {
			// 此时结果包括当前位置
			res = max(res, r-l+1)
			continue
		}

		// 有多余的就给它补上
		if K > 0 {
			// 此时结果包括当前位置
			res = max(res, r-l+1)
			K--
			continue
		}

		// 去掉第一个被替换的0, 放到当前位置上
		for A[l] == 1 {
			l++
		}
		l++
		// 此时结果不包括当前位置
		res = max(res, r-l+1)
	}
	return res
}

func longestOnes3(A []int, K int) (ans int) {
	// lsum表示[0, left]里0的数量, rsum表示[0, right]中0的数量
	var left, lsum, rsum int
	for right, v := range A {

		// 注意, 这里遇到0会+1, 遇到1会+0
		rsum += 1 - v

		// 如果后半段里0的数量超过了K可以补充的最大值, 就要把left向前缩
		for lsum < rsum-K {
			// 这里相当于找到第一个0
			lsum += 1 - A[left]
			left++
		}
		// 统计的数据就是[left, right]这个窗口的大小
		ans = max(ans, right-left+1)
	}
	return
}

func main() {
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
