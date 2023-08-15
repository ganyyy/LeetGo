package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构建队列, 初始节点都是 入度为0的 节点

	var empty = struct{}{}

	// 构建邻接表
	m := make(map[int]map[int]struct{}, numCourses)
	// 构建入读为0的队列
	queue := make([]int, numCourses)

	for _, v := range prerequisites {
		a, b := v[0], v[1]
		s, ok := m[b]
		if !ok {
			s = make(map[int]struct{})
			m[b] = s
		}
		// 构建邻接表, 意思是存在一条边由 b->a
		s[a] = empty
		// 入度+1
		queue[a]++
	}

	// 统计所有入度为0的点作为将要执行的点
	todo := make([]int, 0, numCourses)
	for i, v := range queue {
		if v == 0 {
			todo = append(todo, i)
		}
	}

	// 统计所有的安全点(入度为0的点)
	var cnt int
	for len(todo) != 0 {
		h := todo[0]
		todo = todo[1:]

		cnt++
		for i, _ := range m[h] {
			// 入度-1
			queue[i]--
			// 如果入度变成了0, 那么就成为了安全点, 加入到todo队列中
			if queue[i] == 0 {
				todo = append(todo, i)
			}
		}
	}
	return cnt == numCourses
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses)
		// visited有三种状态
		// 0 未访问
		// 1 正在访问
		// 2 访问完成
		visited = make([]int, numCourses)
		// result []int
		valid = true
		dfs   func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				// 如果是没有访问过的点, 访问
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				// 如果是正在访问的点, 报错
				// 说明此时产生了环, 无法完成学习
				valid = false
				return
			}
		}
		visited[u] = 2
		// result = append(result,u)
	}

	// info[1] 的后置课程
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

func canFinish3(numCourses int, prerequisites [][]int) bool {
	var empty = struct{}{}

	// m[a][b] 表示 a是b的前置课程
	var m = make(map[int]map[int]struct{}, numCourses)

	// dep表示入度
	var dep = make([]int, numCourses)

	// 首先构建邻接表
	for _, p := range prerequisites {
		var a, b = p[0], p[1]
		var bm, ok = m[b]
		if !ok {
			bm = make(map[int]struct{})
			m[b] = bm
		}
		// b->a, dep[a]++
		bm[a] = empty
		dep[a]++
	}

	// 找出所有入度为0的课程, 这些课程都可以作为发起点
	var todo = make([]int, 0, numCourses)
	for i, d := range dep {
		if d == 0 {
			todo = append(todo, i)
		}
	}

	// 最后, 遍历所有的发起点, 比较完成的课程和总课程之间的数量差值
	var num int

	for len(todo) != 0 {
		var t = todo[0]
		todo = todo[1:]

		num++

		for i := range m[t] {
			// 如果这个课程的入度已经为0了, 说明存在环, 直接返回错误即可
			// if dep[i] == 0 {
			//     return false
			// }
			dep[i]--
			// 如果某一个课程减到0了, 说明可以作为一个发起点
			if dep[i] == 0 {
				todo = append(todo, i)
			}
		}
	}

	// 比较统计的课程数量和总的课程数量
	return num == numCourses
}

func canFinish4(numCourses int, prerequisites [][]int) bool {
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
		if dep == 0 {
			queue = append(queue, i)
		}
	}

	// 判断起始的入度为0的点
	if len(queue) == 0 {
		return false
	}

	var cnt int // 访问完成的课程的数量

	for len(queue) != 0 {
		var p = queue[0]
		queue = queue[1:]
		cnt++
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

	// 再次统计所有节点的入度是否变成了0
	// for _, dep := range visits {
	//     if dep != 0 {
	//         return false
	//     }
	// }
	// return true
	return cnt == numCourses
}

func canFinish5(numCourses int, prerequisites [][]int) bool {
	// 所有入度为0的课程, 就是起始课程

	// pre -> [post]
	var nextCourses = make([][]int, numCourses)
	// [post]
	var inDegrees = make([]int, numCourses)
	// 累计学习完成的课程
	var past int

	for _, prerequisite := range prerequisites {
		post, pre := prerequisite[0], prerequisite[1]
		inDegrees[post]++
		nextCourses[pre] = append(nextCourses[pre], post)
	}

	var cur []int
	var next []int

	for course, ind := range inDegrees {
		if ind == 0 {
			cur = append(cur, course)
		}
	}

	for len(cur) != 0 {
		past += len(cur)
		for _, course := range cur {
			for _, nextCourse := range nextCourses[course] {
				inDegrees[nextCourse]--
				if inDegrees[nextCourse] <= 0 {
					next = append(next, nextCourse)
				}
			}
		}

		cur, next = next, cur[:0]
	}

	return past == numCourses
}
