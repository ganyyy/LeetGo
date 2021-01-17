package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	// 别说了, 直接并查集是吧

	// 首先按照邮箱建立基础的并查集

	// 每个邮箱的父节点
	var fa = make(map[string]string, len(accounts))
	var fn = make(map[string]string, len(accounts))

	for _, account := range accounts {
		var name = account[0]
		for _, email := range account[1:] {
			fa[email] = email
			fn[email] = name
		}
	}

	var find func(string) string

	find = func(email string) string {
		var res = fa[email]
		if res != email {
			res = find(res)
			fa[email] = res
		}
		return res
	}

	// 将有关系的邮箱合并到一起
	var union = func(from, to string) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	for _, account := range accounts {
		// 第一波合并, 找出所有的节点的根节点
		var fe = account[1]
		for _, email := range account[2:] {
			union(email, fe)
		}
	}

	var tmp = make(map[string][]string, len(fa))
	for email := range fa {
		var root = find(email)
		tmp[root] = append(tmp[root], email)
	}

	var res = make([][]string, 0)
	for re, emails := range tmp {
		sort.Strings(emails)
		res = append(res, append([]string{fn[re]}, emails...))
	}

	return res
}

func accountsMerge2(accounts [][]string) (ans [][]string) {
	// 优化方向: 将string hash转为 int hash.
	emailToIndex := map[string]int{}
	emailToName := map[string]string{}
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, has := emailToIndex[email]; !has {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = name
			}
		}
	}

	parent := make([]int, len(emailToIndex))
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(emailToIndex[email], firstIndex)
		}
	}

	indexToEmails := map[int][]string{}
	for email, index := range emailToIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}

	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return
}
