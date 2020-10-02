package main

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{i, idx}
		}
		m[v] = i
	}
	return nil
}

func main() {

}
