package main

func missingNumber(nums []int) int {
	res := 0
	// 从0到n以功n+1个数, 两两异或, 剩下的就是缺失的那个数
	for i := 0; i < len(nums); i++ {
		res ^= i ^ nums[i]
	}
	return res ^ len(nums)
}

func main() {

}
