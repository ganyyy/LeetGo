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
