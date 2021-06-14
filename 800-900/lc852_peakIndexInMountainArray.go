package main

func peakIndexInMountainArray(arr []int) int {
	// 一次遍历

	// for i := 1; i < len(arr)-1; i++ {
	//     if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
	//         return i
	//     }
	// }

	// 二分查找
	var l, r = 1, len(arr)

	for l < r {
		var mid = l + (r-l)>>1
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid] > arr[mid-1] {
			// 递增的, 说明拐点在右边
			l = mid + 1
		} else {
			// 递减的, 说明拐点在左边
			r = mid
		}
	}

	return 0
}
