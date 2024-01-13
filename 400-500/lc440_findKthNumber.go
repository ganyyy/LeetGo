package main

import "fmt"

func getSteps(cur, n int) (steps int) {

	/*
	     1     2     3     4     5     6     7     8     9
	    / \
	   10 11  12 13

	*/

	// 1, 1, 1
	// 10, 19, 11
	// 100, 199 111
	// n是最大值,
	// 每一层级的最大值, 也就是最后一个数, 也就是cur*10+9
	first, last := cur, cur
	for first <= n {
		// 上限是n, 但是如果last > n, 那么就需要取n
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return
}

func findKthNumber(n, k int) int {

	/*
		- 比如n=30, k = 15. 此时的字典序是
		1(,10,11,12,13,14,15,16,17,18,19),2(,20,21,22,23,24,25,26,27,28,29),3(,30),4,5,6,7,8,9
		此时的查找流程是:
		首先需要定位处于哪一层. 通过getSteps(cur, n)可以得到距离n最近的层级. 首先可以确定的是, 这个无法到达第三层, 因为到达第三层意味着n >= 100, 但是n=30
		1(,10,11,12,13,14,15,16,17,18,19)是11个数, < 15, 所以需要找下一个根节点, 即cur = 2, k = 14 - 11 = 3
		再次定位以2为根节点的层, 2(,20,21,22,23,24,25,26,27,28,29)是11个数, > 3, 所以需要找到具体的层, 即cur = 20, k = 3 - 1 = 2
		以20为根节点的层, 只有一个数, 所以返回的steps = 1, k = 2 - 1 = 1, cur = 20+1 = 21
		以21为根节点的层, 只有一个数, 所以返回的steps = 1, k = 1 - 1 = 0, cur = 21+1 = 22
	*/

	cur := 1 // 从第一层开始, 根节点分别为 1,2,3,4,5
	// 下一层分别为 10,11,12...   20,21,22...   30,31,32...
	// 个位数的情况下,
	k--
	// 整体的查找方向, 就是 由左上到右下的方向查找
	for k > 0 {
		steps := getSteps(cur, n) // 先确定大概的层数, 在确定对应的步数
		fmt.Println("steps:", steps, "cur:", cur, "k:", k)
		if steps <= k { // 如果k >= steps, 说明要找下一个根
			k -= steps
			cur++
		} else { // 否则, 就需要找到具体的层
			cur *= 10
			// 1 -> 1X -> 1XX -> 1XXX, 这种在字典序上相当于是连续的, 所以可以直接--
			k--
		}
	}
	return cur
}

func findKthNumber2(n int, k int) int {
	var cur = 1 // 第一个数
	k--         // 算上第一个

	for k > 0 {
		var step = getSteps(cur, n)
		if step <= k {
			// fmt.Println("less K:", step, cur, k)
			k -= step // 第k个数不在当前基数对应的子树中,
			cur++     // 查找下一个基数
		} else {
			// 这种情况啥时候会出现呢?
			// 因为getSteps获取的到达某一层首个数字经历的步数
			// 举个栗子: 就是某一层在前n中, 但是k 不属于下一个基数的子树
			// fmt.Println("big K:", step, cur, k)
			cur *= 10 // 第k个数在当前基数中的某个子树中,
			k--       // 取当前层的第0个, 所以减少了一步
		}
	}
	return cur
}

func main() {
	println(findKthNumber(30, 15))
}
