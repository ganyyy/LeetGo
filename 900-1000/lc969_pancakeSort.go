package main

func pancakeSort(arr []int) []int {
	// 将最大的数移动到后边
	var ret []int
	var n = len(arr)

	for n != 0 {
		var idx = Index(arr, n)
		if idx == n-1 {
			n--
			continue
		}
		Reverse(arr[:idx+1])
		Reverse(arr[:n])
		ret = append(ret, idx+1, n)
		n--
	}
	// fmt.Println(arr)
	return ret
}

func Index(arr []int, n int) int {
	for i, v := range arr {
		if v == n {
			return i
		}
	}
	return -1
}

func Reverse(arr []int) {
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
	// fmt.Println(arr)
}
