package main

func totalHammingDistance(nums []int) int {
	// 暴力算法不可取啊

	// 统计每一位上1的数量, 然后进行组合, 最终计算的结果就是答案

	var cnt [31]int

	// 这里应该可以优化一下
	for _, v := range nums {
		for i := 0; i <= 30 && v != 0; i++ {
			if v&1 == 1 {
				cnt[i]++
			}
			v = v >> 1
		}
	}
	// 111 000
	// 6个数字, 3个1, 3个0
	// 那么有 3 * 3 = 9 种不同的模式
	var total = len(nums)
	var ret int
	for _, v := range cnt {
		// 一个没有, 或者都有. 这个无法产生距离
		if v == 0 || v == total {
			continue
		}
		ret += (total - v) * v
	}
	return ret
}

func totalHammingDistance2(nums []int) (ret int) {
	var cnt int
	var ln = len(nums)
	for i := 0; i <= 30; i++ {
		cnt = 0
		for _, v := range nums {
			cnt += (v >> i) & 1
		}
		ret += (ln - cnt) * cnt
	}
	return
}
