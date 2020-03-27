package main

func findRepeatNumber(nums []int) int {
	// 这需要消耗on的空间,
	// 如果对空间有需求, 可以
	set := make([]bool, len(nums))
	for _, v := range nums {
		if !set[v] {
			set[v] = true
		} else {
			return v
		}
	}
	return -1
}

func main() {

}
