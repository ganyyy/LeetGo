package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type array []int

func (a array) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a array) Len() int {
	return len(a)
}

func (a array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *array) Push(x interface{}) {
	(*a) = append((*a), x.(int))
}

func (a *array) Pop() (v interface{}) {
	*a, v = (*a)[:a.Len()-1], (*a)[len(*a)-1]
	return v
}
func findKthLargest(nums []int, k int) int {
	// 用堆
	if len(nums) < k {
		return 0
	}
	var t array = nums[:k]
	heap.Init(&t)
	for _, v := range nums[k:] {
		if v > t[0] {
			t[0] = v
			heap.Fix(&t, 0)
		}
	}
	return t[0]
}

func findKthLargest2(nums []int, k int) int {
	var arr = make([]int, k)
	var heapArr = func() {
		// 这是一个小顶堆, 那么堆顶一定是最小值
		for i := k/2 - 1; i >= 0; i-- {
			var root = i
			for {
				var child = root*2 + 1
				if child >= k {
					break
				}
				if child+1 < k && arr[child+1] < arr[child] {
					child++
				}
				if arr[child] >= arr[root] {
					break
				}
				// 这一步是关键欸
				arr[root], arr[child] = arr[child], arr[root]
				root = child
			}
		}
	}

	copy(arr, nums[:k])
	heapArr()

	for _, v := range nums[k:] {
		if v <= arr[0] {
			continue
		}
		arr[0], arr[k-1] = arr[k-1], v
		heapArr()
	}

	return arr[0]
}

func findKthLargest3(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())

	var quick func(l, r int) int
	var partition = func(l, r int) int {
		// nums[r]是random选取的随机节点
		var pivot = nums[r]
		var idx = l - 1
		// idx左边的均<= pivot
		// idx右边的均 > pivot
		for l < r {
			if nums[l] <= pivot {
				idx++
				nums[idx], nums[l] = nums[l], nums[idx]
			}
			l++
		}
		// 此时idx对应的位置是 > pivot的
		nums[idx+1], nums[r] = nums[r], nums[idx+1]
		return idx + 1
	}
	var random = func(l, r int) int {
		pos := rand.Int()%(r-l+1) + l
		// pos就是要定位的点
		nums[pos], nums[r] = nums[r], nums[pos]
		return partition(l, r)
	}

	index := len(nums) - k

	quick = func(l, r int) int {
		var pos = random(l, r)
		if pos == index {
			return nums[pos]
		} else if pos < index {
			return quick(pos+1, r)
		}
		return quick(l, pos-1)
	}

	return quick(0, len(nums)-1)
}

func up(nums []int, child int) {
	for {
		// 找到父节点
		parent := (child - 1) / 2
		if parent == child || nums[parent] < nums[child] {
			// 没有父节点, 或者父节点小于自身, 说明不需要继续上升
			break
		}
		nums[parent], nums[child] = nums[child], nums[parent]
		child = parent
	}
}

func down(nums []int, start int) bool {
	parent := start
	length := len(nums)
	for {
		// 左子节点
		child := parent*2 + 1
		if child >= length || child < 0 {
			break
		}
		// 选出两个子节点的最小值
		if right := child + 1; right < length && nums[right] < nums[child] {
			child = right
		}
		if nums[child] > nums[parent] {
			// 子节点 > 父节点
			break
		}
		nums[child], nums[parent] = nums[parent], nums[child]
		parent = child
	}
	return parent > start
}

func fix(nums []int, idx int) {
	// idx始终都是0, 所以不需要down
	if !down(nums, idx) {
		up(nums, idx)
	}
}

func initHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		down(nums, i)
	}
}

func findKthLargestHeap(nums []int, k int) int {
	var arr = nums[:k]
	initHeap(arr)
	for _, num := range nums[k:] {
		if arr[0] < num {
			arr[0] = num
			fix(arr, 0)
		}
	}
	return arr[0]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
