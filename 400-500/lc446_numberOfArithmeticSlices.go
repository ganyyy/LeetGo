package main

func numberOfArithmeticSlices(nums []int) (ans int) {

	// f[i][d] 表示nums[:i]中相对于nums[i]而言, 差值为 d 的数值的数量
	f := make([]map[int]int, len(nums))

	// 计算 [:i] 区间内的两两之差,
	// 将 nums[j], nums[i]作为等差数列的倒数第二项和倒数第一项, 将差值作为d,
	// 那么, 如果 nums[j][d] 存在的话, 说明在[:j]中有一个索引k, 满足 nums[i]-nums[j] = nums[j]-nums[k]
	// 此时 k,j,i 就是一个等差数列
	// 相对的, 也可能存在x,k,j是一个等差数列, 此时对应的次数已经计入到 f[j][d]中了
	for i, x := range nums {
		f[i] = map[int]int{}
		for j, y := range nums[:i] {
			// 计算差值
			d := x - y
			// 获取前置值的对应的差值
			cnt := f[j][d]
			// 不管是否存在, 直接加上即可
			ans += cnt
			// cnt+1表示算了了j本身
			f[i][d] += cnt + 1
		}
	}
	return
}
