package main

func maxEqualFreq(nums []int) (ans int) {
	freq := map[int]int{}  // 某一频次出现的数的个数
	count := map[int]int{} // 某一个值出现的次数
	maxFreq := 0           // 当前出现的最大频次
	for i, num := range nums {
		// 维护频次, 出现次数, 最大频率
		if count[num] > 0 {
			freq[count[num]]--
		}
		count[num]++
		maxFreq = max(maxFreq, count[num])
		freq[count[num]]++

		// 1. 最大频次是1, 删了都是0, 满足条件
		// 2. 到目前为止, 去掉一个最大值剩余的都是次大频率的值(因为一次只能删除1个, 所以必须要比较最大和次大, 且最大只有一个)
		// 3. 都是最大频率, 除了一个单独的
		if maxFreq == 1 ||
			freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 ||
			freq[maxFreq]*maxFreq+1 == i+1 && freq[1] == 1 {
			ans = max(ans, i+1)
		}
	}
	return
}
