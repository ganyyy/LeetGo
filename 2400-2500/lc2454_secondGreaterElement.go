package main

func secondGreaterElement(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	// s: 递减栈, 当新加入一个x时, 会将所有比x小的元素转移到t中. 那么当新的x1到来时, 首先对比t的栈顶元素, 如果比x1小, 那么t的栈顶元素的下一个更大元素就是x1
	// 因为t中的元素的下一个更大元素在s中, 同时t中所有的元素都比s中的元素小, 所以如果t中存在小于x1的元素, t中的这个元素的下下个更大元素就是x1
	var s []int
	var t []int
	for i, x := range nums {
		for len(t) > 0 && nums[t[len(t)-1]] < x {
			ans[t[len(t)-1]] = x // t 栈顶的下下个更大元素是 x
			t = t[:len(t)-1]
		}
		j := len(s) - 1
		for j >= 0 && nums[s[j]] <= x {
			j-- // s 栈顶的下一个更大元素是 x
		}
		t = append(t, s[j+1:]...) // 把从 s 弹出的这一整段元素加到 t
		s = append(s[:j+1], i)    // 当前元素（的下标）加到 s 栈顶
	}
	return ans
}
