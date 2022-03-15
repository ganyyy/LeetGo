package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	m := make([][]int, numCourses)   // 课程对应的所有前置课程
	point := make([]int, numCourses) // 每一个课程的入度, 入度为0的课程表示没有任何前置课程
	for _, v := range prerequisites {
		// 第一个表示当前课程, 第二个表示前置课程

		// 表示从 前置课程 能到达 当前课程
		m[v[1]] = append(m[v[1]], v[0])
		// 入度+1
		point[v[0]]++
	}

	// 找到所有入度为0的点, 这就是始发点
	queue := make([]int, 0)
	for i, v := range point {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	// 最终结果
	res := make([]int, numCourses)
	var idx int

	// 开始找路径
	for ln := len(queue); ln != 0; ln = len(queue) {
		// queue里始终都是入度为0的课程, 通过课程找到后继
		// 对应的后继计数-1, 如果降到了0也会加入到queue中

		// 出队
		t := queue[0]
		queue = queue[1:]

		// 更新一个结果
		res[idx] = t
		idx++

		// 找出对应的后继
		lst := m[t]
		// 没有前置, 直接开始下一轮
		if 0 == len(lst) {
			continue
		}
		for _, v := range lst {
			point[v]--
			if 0 == point[v] {
				// 减到零也加入到队列中
				queue = append(queue, v)
			}
		}
	}

	// 可能会出现环?
	if numCourses == idx {
		return res
	} else {
		return []int{}
	}
}

func findOrder2(numCourses int, prerequisites [][]int) []int {
	var relations = make([][]int, numCourses)
	var visits = make([]int, numCourses)

	for _, pre := range prerequisites {
		// 构建 a->b 的映射
		// 增加 b 的入度
		var a, b = pre[1], pre[0]
		relations[a] = append(relations[a], b)
		visits[b]++
	}

	var queue []int
	for i, dep := range visits {
		if dep != 0 {
			continue
		}
		queue = append(queue, i)
	}

	// 判断起始的入度为0的点
	if len(queue) == 0 {
		return nil
	}

	var ret = make([]int, 0, numCourses)
	for len(queue) != 0 {
		var p = queue[0]
		queue = queue[1:]
		ret = append(ret, p)
		for _, next := range relations[p] {
			if visits[next] <= 0 {
				continue
			}
			visits[next]--
			if visits[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(ret) == numCourses {
		return ret
	}
	return nil
}

func main() {

}
