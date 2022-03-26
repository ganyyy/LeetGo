package main

func reversePairs(nums []int) int {
	tmp := make([]int, len(nums))
	return reverse(nums, tmp, 0, len(nums)-1)
}

func reverse(nums, tmp []int, left, right int) int {
	// 只剩一个时一定时有序的
	if right <= left {
		return 0
	}
	// 规避可能的溢出
	mid := left + (right-left)/2

	// 左右区间分开进行合并
	cnt := reverse(nums, tmp, left, mid) + reverse(nums, tmp, mid+1, right)

	// 先看一下是不是已经有序了, 已经有序了就没必要继续排下去了
	if nums[mid] <= nums[mid+1] {
		return cnt
	}

	// 进行左右两个有序区间的合并, 同时计算逆序数对
	i, j, k := left, mid+1, left
	for ; i <= mid && j <= right; k++ {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
			// 合并右边时, 看左边剩下几个比当前数大的, 有几个就是几个逆序对
			cnt += mid - i + 1
		}
	}
	// 看看有没有剩下的

	// 如果左边有剩下的, 所有的逆序对在右边入栈时已经全部加完了, 这里不需要在进行任何处理
	for ; i <= mid; i++ {
		tmp[k] = nums[i]
		k++
	}
	// 如果右边有剩下的, 此时所有右边剩下的都比左边大, 自然也不可能存在逆序对
	for ; j <= right; j++ {
		tmp[k] = nums[j]
		k++
	}

	// 使nums变成有序数组
	for k = left; k <= right; k++ {
		nums[k] = tmp[k]
	}

	return cnt
}

func reversePairs2(nums []int) int {
	var space = make([]int, len(nums))

	var cnt int
	var merge = func(left, right []int) {
		var li, ri int
		var idx int
		for li < len(left) && ri < len(right) {
			if left[li] <= right[ri] {
				space[idx] = left[li]
				idx++
				li++
			} else {
				// 这一句是关键: 因为从这个数往后的所有数都是大于right[ri]的
				cnt += len(left) - li
				space[idx] = right[ri]
				idx++
				ri++
			}
		}

		for li < len(left) {
			space[idx] = left[li]
			idx++
			li++
		}
		for ri < len(right) {
			space[idx] = right[ri]
			idx++
			ri++
		}
	}

	var mergeSort func(start, end int)

	mergeSort = func(start, end int) {
		if end-start <= 1 {
			return
		}
		var mid = start + (end-start)/2
		mergeSort(start, mid)
		mergeSort(mid, end)
		merge(nums[start:mid], nums[mid:end])
		copy(nums[start:end], space[:end-start])
	}

	mergeSort(0, len(nums))

	// fmt.Println(nums)

	return cnt

}

func main() {

}
