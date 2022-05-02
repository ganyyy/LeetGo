//go:build ignore

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	// 层次遍历+合并呗...

	var inOrder = func(root *TreeNode) (ret []int) {
		var stack []*TreeNode

		for len(stack) != 0 || root != nil {
			if root == nil {
				root = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				ret = append(ret, root.Val)
				root = root.Right
			} else {
				stack = append(stack, root)
				root = root.Left
			}
		}

		return
	}

	var merge = func(nums1, nums2 []int) []int {
		var ret = make([]int, 0, len(nums1)+len(nums2))
		var left, right int

		for left < len(nums1) && right < len(nums2) {
			if nums1[left] < nums2[right] {
				ret = append(ret, nums1[left])
				left++
			} else {
				ret = append(ret, nums2[right])
				right++
			}
		}

		for ; left < len(nums1); left++ {
			ret = append(ret, nums1[left])
		}
		for ; right < len(nums2); right++ {
			ret = append(ret, nums2[right])
		}
		return ret
	}

	return merge(inOrder(root1), inOrder(root2))
}
