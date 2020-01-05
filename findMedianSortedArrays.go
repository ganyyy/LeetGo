package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	len1, len2 := len(nums1), len(nums2)
	val1, val2 := 0, 0
	if len1 == 0 {
		mid := len2 / 2
		if len2&1 == 0 {
			return float64(nums2[mid]+nums2[mid-1]) / 2
		} else {
			return float64(nums2[mid])
		}
	}
	start1, start2 := 0, 0
	mid := (len1 + len2 + 1) / 2
	getVal := func(arr []int, pos int) int {
		if pos >= len(arr) {
			pos = len(arr) - 1
		}
		return arr[pos]
	}
	getMin := func(pos1, pos2 int) int {
		if pos1 < 0 || pos1 >= len1 {
			return nums2[pos2]
		}
		if pos2 < 0 || pos2 >= len2 {
			return nums1[pos1]
		}
		if nums1[pos1] < nums2[pos2] {
			return nums1[pos1]
		}
		return nums2[pos2]
	}

	for {
		if start1 >= len1 {
			val1 = nums2[start2+mid-1]
			val2 = getMin(start1, start2+mid)
			break
		}
		if mid == 1 {
			if nums1[start1] <= nums2[start2] {
				val1 = nums1[start1]
				val2 = getMin(start1+1, start2)
			} else {
				val1 = nums2[start2]
				val2 = getMin(start1, start2+1)
			}
			break
		}
		k := mid / 2
		mid -= k
		if getVal(nums1, start1+k-1) > getVal(nums2, start2+k-1) {
			start2 += k
		} else {
			start1 += k
			if start1 > len1 {
				mid += start1 - len1
			}
		}
	}

	if (len1+len2)&1 == 0 {
		return float64(val1+val2) / 2
	} else {
		return float64(val1)
	}
}

func main() {

}
