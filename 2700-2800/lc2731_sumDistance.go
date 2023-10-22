package main

import "sort"

func sumDistance(nums []int, s string, d int) int {
	const MOD = 1_000_000_007
	// 这里要怎么理解呢?
	//
	/*
	   假设在某一时刻,

	       a->   <-b
	   ...-1...0...1...
	           ⬇
	       a,b之间相撞
	           ⬇
	     <-a       b->
	   从图上来看,
	   完全可以继续将a看成b,b看成a,
	   这样完全不影响最终的结果
	*/
	for idx, dir := range s {
		// L & 2 = 0 - 1 = -1
		// R & 2 = 2 - 1 =  1
		nums[idx] += d * int(dir&2-1)
	}
	sort.Ints(nums)

	// sum 前缀和
	// 计算两两之间的距离时:
	// 设当前位置为 idx, 值为num
	// 那么idx左边存在idx个数不超过num,
	//  那么 nums[idx]和前边的数的两两之间的距离差为
	//  (nums[idx]-nums[0]) + (nums[idx]-nums[1]) + ...
	//      = idx * nums[idx] - (sum(nums[:idx]))
	// 累加到最后, 就是所有的距离差
	var sum, ret int
	for idx, num := range nums {
		ret = (ret + idx*num - sum) % MOD
		sum += num
	}

	return ret
}
