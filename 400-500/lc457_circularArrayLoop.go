package main

func circularArrayLoop(nums []int) bool {
	const InfMax = 1000
	n := len(nums)
	for index, num := range nums {
		if num >= -InfMax && num <= InfMax {
			lastIndex := index
			i := index

			if nums[i] > 0 {
				// 先看大于0的正方向前进
				for nums[i] > 0 && nums[i] <= InfMax {
					lastIndex = i
					i = (nums[i] + i) % n
					// 每一个到达的点都进行一个特殊的标记, 表示是不是从当前索引出发
					// 并回到原点的
					nums[lastIndex] = InfMax + index + 1
				}

				if lastIndex != i && nums[i] == InfMax+index+1 {
					return true
				}
			} else {
				// 这个看的是负数索引的逻辑
				for nums[i] < 0 && nums[i] >= -InfMax {
					lastIndex = i
					i = (n - (-nums[i] % n) + i) % n
					// 道理相同, 逻辑相似
					nums[lastIndex] = -InfMax - index - 1
				}
				if lastIndex != i && nums[i] == -InfMax-index-1 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	print(circularArrayLoop([]int{3, 1, 2}))
}
