package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 以l1 为返回链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var add int
	c1, c2 := l1, l2
	var p1 *ListNode = c1
	for c1 != nil && c2 != nil {

		if v := c1.Val + c2.Val + add; v >= 10 {
			add = 1
			c1.Val = v - 10
		} else {
			add = 0
			c1.Val = v
		}
		p1 = c1
		c1 = c1.Next
		c2 = c2.Next
	}
	// 结尾判断
	if c1 == nil {
		p1.Next = c2
		for add != 0 && c2 != nil {
			if v := c2.Val + add; v >= 10 {
				c2.Val = v - 10
				add = 1
			} else {
				c2.Val = v
				add = 0
			}
			p1 = c2
			c2 = c2.Next
		}
	} else if c2 == nil {
		for add != 0 && c1 != nil {
			if v := c1.Val + add; v >= 10 {
				c1.Val = v - 10
				add = 1
			} else {
				c1.Val = v
				add = 0
			}
			p1 = c1
			c1 = c1.Next
		}
	}
	if add != 0 {
		p1.Next = &ListNode{Val: 1}
	}
	return l1
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var odd int
	var pre *ListNode
	var ret = l1
	for l1 != nil && l2 != nil {
		v := l1.Val + l2.Val + odd
		if v >= 10 {
			v -= 10
			odd = 1
		} else {
			odd = 0
		}
		l1.Val = v
		pre = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil && l2 != nil {
		pre.Next = l2
		l1 = l2
	}
	for ; l1 != nil && odd != 0; l1 = l1.Next {
		l1.Val += odd
		pre = l1
		if l1.Val >= 10 {
			l1.Val -= 10
			odd = 1
		} else {
			odd = 0
		}
	}

	if odd != 0 {
		pre.Next = &ListNode{
			Val: 1,
		}
	}

	return ret
}

func main() {

}
