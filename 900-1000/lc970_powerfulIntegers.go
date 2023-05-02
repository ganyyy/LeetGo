package main

func powerfulIntegers(x int, y int, bound int) []int {
	res := make(map[int]bool)
	value1 := 1
	// 标准的二重循环
	for i := 0; i < 21; i++ {
		value2 := 1
		for j := 0; j < 21; j++ {
			value := value1 + value2
			if value <= bound {
				res[value] = true
			} else {
				break
			}
			value2 *= y
		}
		if value1 > bound {
			break
		}
		value1 *= x
	}
	var result []int
	for k := range res {
		result = append(result, k)
	}
	return result
}
