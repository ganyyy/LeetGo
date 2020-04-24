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

func main() {

}
