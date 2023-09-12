package main

// 关于计算ID的方式.
// Golang中整数的除法都是靠向0的.
// 假设 w = 5. 那么当x == 0时, 返回的是0. 所以[0,5]6个数是一组
// 对应的, [-6, -1]应该是一组. 且组ID应该为-1
// 要满足这个条件, 当x = -6/-1时都要返回-1. 所以负数计算ID需要通过以下方式计算
// 只要满足每个数字的组ID唯一不重合即可, 计算方式多种多样.

func getID(x, w int) int {
	// 对于正负数的不同处理方式
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	// map维护了大小为k的区间内, 所有值的 "ID"
	mp := map[int]int{}
	for i, x := range nums {

		// ID的计算方式可以理解为
		// 为啥这里是t+1 呢?
		// 和 x 匹配的值在范围 [x-t, x+t]中
		// 一个桶最多保存t+1个元素 x + [0, t]一共t+1个元素
		id := getID(x, t+1)

		// 如果出现在同一个桶中, 那么可以直接返回
		if _, has := mp[id]; has {
			return true
		}

		// 对比前一个桶和后一个桶是否存在, 如果存在的话就看看是不是满足条件
		if y, has := mp[id-1]; has && abs220(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs220(x-y) <= t {
			return true
		}
		// 这里为什么先赋值, 然后再删除?
		// 如果x == nums[i-k], 那该怎么处理?
		// 如果这样的话, 第一个条件就会成功判定了.
		// 所以这里的 nums[i-k]在整个区间内的ID一定是唯一的
		mp[id] = x
		// 维护大小为k的窗口
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}

func abs220(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func containsNearbyAlmostDuplicate2(nums []int, k, t int) bool {
	// map维护了大小为k的区间内, 所有值的 "ID"
	mp := map[int]int{}
	for i, x := range nums {

		// ID的计算方式可以理解为
		// 为啥这里是t+1 呢?
		// 和 x 匹配的值在范围 [x-t, x+t]中
		// 一个桶最多保存t+1个元素 x + [0, t]一共t+1个元素
		id := getID(x, t+1)

		// 如果出现在同一个桶中, 那么可以直接返回
		if _, has := mp[id]; has {
			return true
		}

		// 还是因为匹配的半径问题. [x-t, x+t] 最多会跨两个桶, 可以细节区分是前半部分还是后半部分
		// 对比前一个桶和后一个桶是否存在, 如果存在的话就看看是不是满足条件
		if y, has := mp[id-1]; has && abs220(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs220(x-y) <= t {
			return true
		}
		// 这里为什么先赋值, 然后再删除?
		// 如果x == nums[i-k], 那该怎么处理?
		// 如果这样的话, 第一个条件就会成功判定了.
		// 所以这里的 nums[i-k]在整个区间内的ID一定是唯一的
		mp[id] = x
		// 维护大小为k的窗口
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}
