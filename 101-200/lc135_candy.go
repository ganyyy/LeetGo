package main

func candy(ratings []int) int {
	// 贪心的思路去搞
	// 第一位糖果为1
	// 和前边的数据进行对比, 如果比前一位要大, 那么糖果+=1
	// 如果比前一位小, 那么这一位就从1开始
	// 此时需要判断一下, 如果前一位也是1, 说明当前时递减的, 那么前一位也需要糖果+1
	// 一直持续到递减序列结束时. 如果峰值和+1后的递减序列值第一个值相同, 那么就把峰值并入到递减序列中
	// 难点在于如何记录峰值
	// 还有遇到相同值时如何处理? 尽量往少的给

	var res = 1 // 总的糖果数量
	var pre = 1 // 前一个同学的糖果数量
	var dec = 0 // 当前递减序列的长度
	var inc = 1 // 最近的递增序列的长度

	// 从第二个同学开始遍历
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			// 重置递减序列的长度
			dec = 0
			if ratings[i] == ratings[i-1] {
				// 如果相等, 那么就往少的分, 也就是给一个
				pre = 1
			} else {
				// 比前一个同学多一个
				pre++
			}
			// 更新当前同学的糖果数量
			res += pre
			// 更新一下递增序列当前的长度
			inc = pre
		} else {
			// 更新一下递减序列的长度
			dec++
			// 如果当前递减序列的长度已经和前一个递增序列相等了
			// 就将原来的递增序列的最后一位并一个过来
			if dec == inc {
				dec++
			}
			// 这里相当于把原来的递减序列的所有值都多加了一个糖果
			res += dec
			// 只要是递减的, 就从1开始
			pre = 1
		}
	}
	return res
}

func candy3(ratings []int) int {
	n := len(ratings)
	// inc: 递增序列的长度
	// dec: 递减序列的长度
	// pre: 前一个节点所得的糖果值
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			// 当前处于一个递增序列, 递减序列的长度清0
			dec = 0
			if ratings[i] == ratings[i-1] {
				// 如果相等, 那么可以从1开始分配
				pre = 1
			} else {
				// 否则就必须要大于前一个节点的值
				pre++
			}
			ans += pre
			inc = pre
		} else {
			// 当前递减
			dec++
			if dec == inc {
				// 特殊情况1: 递减序列的长度等同于上一次递增序列, 那么峰值需要并过来
				dec++
			}
			// 这个怎么理解呢?
			// 1 => +1
			// 2, 1 => +1,+2
			// 3, 2, 1 => +1,+2,+3
			// 2,3,4,3,2,1 => (inc == dec) => +1,+2,+3(+4),+3,+2,+1. 所以需要补1
			ans += dec
			pre = 1
		}
	}
	return ans
}

func candy2(ratings []int) int {
	var ln = len(ratings)
	if ln == 0 {
		return 0
	}
	var res int

	// 两次遍历, 找到满足同时大于左右的点

	var left = make([]int, ln)
	left[0] = 1
	for i := 1; i < ln; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	var right = 1
	res += max(1, left[ln-1])
	for i := ln - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		// 取最大值
		res += max(right, left[i])
	}

	return res
}
