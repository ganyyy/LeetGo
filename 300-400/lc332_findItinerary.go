package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	// 构建一个映射
	m := make(map[string][]string)
	for _, v := range tickets {
		m[v[0]] = append(m[v[0]], v[1])
	}
	for k := range m {
		sort.Strings(m[k])
	}

	var res = make([]string, 0, len(m)+1)
	// 从头开始走
	var helper func(addr string)
	// 对于起点 JFK 而言, 出度一定比入度大1
	// 对于终点 而言, 入度一定比出度大1
	helper = func(addr string) {
		for len(m[addr]) != 0 {
			to := m[addr][0]
			m[addr] = m[addr][1:]
			helper(to)
		}
		// 后边的都处理完了, 将当前的起点作为 前边的终点加入到 队列中
		// 按照终点出现的顺序入队, 最终结果需要进行逆序处理
		res = append(res, addr)
	}
	helper("JFK")
	reverse(res)
	return res
}

func reverse(res []string) {
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
}

func main() {
	tmp := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	fmt.Println(findItinerary(tmp))
}
