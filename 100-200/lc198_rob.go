package main

func rob(nums []int) int {
	// a为前一个没选
	// b为选了前一个
	var a, b int
	for i := 0; i < len(nums); i++ {
		a, b = b, max(nums[i]+a, b)
	}
	return b
}

func main() {

}
