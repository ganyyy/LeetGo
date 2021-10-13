package main

func peakIndexInMountainArray(arr []int) int {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}

func peakIndexInMountainArrayBinarySearch(arr []int) int {
	var l, r = 0, len(arr) - 1
	for l < r {
		// 取中点
		mid := l + (r-l)/2
		if arr[mid] > arr[l] && arr[mid] > arr[r] {
			// 中点的值满足峰值, 但是不一定是最大值. 缩小选择空间
			l++
			r--
		} else if arr[mid] > arr[l] {
			// 中点值大于左边, 小于右边. 此时峰顶应该在mid的右边
			l = mid + 1
		} else if arr[mid] < arr[l] {
			// 中点值大于右边, 小于左边. 此时峰顶应该在mid的左边
			r = mid - 1
		} else if arr[mid] == arr[l] {
			// 相等的情况.. 如果左边小于右边, 说明峰顶在右边, 否则在左边
			if arr[l] > arr[r] {
				r--
			} else {
				l++
			}
		}
	}
	return l
}
