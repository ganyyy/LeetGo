package main

func findClosestElements(arr []int, k int, x int) []int {

	l, r := 0, len(arr)-k

	for l < r {
		mid := (l + r) / 2
		// 找到左端点
		// 对比绝对值, 因为[l:l+k]这段区间内和x的绝对值差一定小于两边的端点
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return arr[l : l+k]
}
