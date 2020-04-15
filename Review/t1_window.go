package main

import "fmt"

func getTarget(target int) [][]int {
	i, j, t := 1, 2, 3
	res := make([][]int, 0)
	for i < j {
		if t < target {
			t += j + 1
			j++
		} else if t > target {
			t -= i
			i++
		} else {
			tmp := make([]int, j-i+1)
			for k := i; k <= j; k++ {
				tmp[k-i] = k
			}
			res = append(res, tmp)
			t -= i
			i++
		}
	}
	return res
}

func main() {
	fmt.Println(getTarget(15))
}
