package main

func tupleSameProduct(nums []int) int {
	var cnt = make(map[int]int)
	var ret int
	for i, iv := range nums {
		for _, jv := range nums[i+1:] {
			cnt[iv*jv]++
		}
	}
	// fmt.Println(cnt)
	// 每一组有两个元素组成, 自身的排列可能为 A(2,2) = 2, 两组组成一个结果, 所以是 2*2
	// 当出现乘积相同次数 num > 1 时, 说明可以组成一个结果. 组成结果数等同于:
	// 从 num 个组合中选出两个的组合成一个结果, 对应的选取方式就是 A(num,2) = num * (num-1)
	for _, num := range cnt {
		if num == 1 {
			continue
		}
		// 2 -> 8: A(2,2) * (A(2,2)*A(2,2))
		// 3 ->24: A(3,2) * (A(2,2)*A(2,2))
		ret += num * (num - 1) * 2 * 2
	}
	return ret
}
