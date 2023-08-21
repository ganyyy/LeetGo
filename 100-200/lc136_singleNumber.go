package main

func singleNumber(nums []int) int {
	// 异或处理, 如果把所有数字异或一边, 最终的结果就是唯一的那个
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func main() {

}
