package main

import "sort"

func numberOfWeakCharactersSB(properties [][]int) int {
	// 分优先级进行排序

	sort.Slice(properties, func(i, j int) bool {
		var a, b = properties[i], properties[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	})

	// 合并相同的分组(?)

	// 现在可以保证的是: 前者的攻击力一定小于后者
	// 攻击力相同时, 前者的防御一定小于后者

	// 用单调递增栈(栈顶最小), 如果一个数成为新的栈顶, 就说明不存在比他更小的值

	var stack []int
	var ret int
	for i, p := range properties {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 {
			var top = properties[stack[len(stack)-1]]
			if top[0] < p[0] && top[1] < p[1] {
				ret++
				stack = stack[:len(stack)-1]
			} else {
				// 从头扫一遍?
				for i := 0; i < len(stack); {
					var cur = properties[stack[i]]
					if cur[0] < p[0] && cur[1] < p[1] {
						copy(stack[i:], stack[i+1:])
						stack = stack[:len(stack)-1]
						ret++
					} else {
						i++
					}
				}
				break
			}
		}
		stack = append(stack, i)
	}

	return ret
}

func numberOfWeakCharacters(properties [][]int) int {
	// 攻击降序, 防御升序. 避免了同攻击时误判的情况
	// 怎么排序是个门道
	// 多条件排序的情况下, 要分清主次关系
	sort.Slice(properties, func(i, j int) bool {
		var a, b = properties[i], properties[j]
		if a[0] != b[0] {
			return a[0] > b[0]
		}
		return a[1] < b[1]
	})

	// fmt.Println(properties)

	var top = properties[0]
	var maxDef = top[1]
	var cnt int
	// 攻击降序保证了前者攻击>=后者
	// 防御降序保证了攻击相同时, 前者防御 <= 后者
	// 这种情况下, 只需要保留迭代到的最大的防御值.
	// 后续的防御值小于最大值, 那么就是严格小于的关系!
	for _, p := range properties[1:] {
		if p[1] < maxDef {
			cnt++
		} else {
			maxDef = p[1]
		}
	}

	return cnt
}
