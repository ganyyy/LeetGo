package main

func findDuplicate(nums []int) int {
	var res int
	// 将数组抽象成链表, 通过快慢指针进行解决
	// 先找到快慢指针相同的点, 再将头指针和慢指针同时移动找到环的入口
	// 即为相似点

	// 找到二者相等 时 的值
	for fast := 0; res != fast || fast == 0; {
		res = nums[res]
		fast = nums[nums[fast]]
	}
	// 在从头开始找,
	for i := 0; res != i; i = nums[i] {
		res = nums[res]
	}
	return res
}

func main() {

}
