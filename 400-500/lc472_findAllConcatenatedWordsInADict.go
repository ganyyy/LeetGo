package main

import "sort"

type trie struct {
	children [26]*trie
	isEnd    bool
}

func (root *trie) insert(word string) {
	node := root
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (root *trie) dfs(word string) bool {
	if word == "" {
		return true
	}
	node := root
	for i, ch := range word {
		node = node.children[ch-'a']
		if node == nil {
			return false
		}
		// 这里就是分割, 查找是否存在子串
		if node.isEnd && root.dfs(word[i+1:]) {
			return true
		}
	}
	return false
}

func findAllConcatenatedWordsInADict(words []string) (ans []string) {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })

	root := &trie{}
	for _, word := range words {
		if word == "" {
			continue
		}
		// 为啥插入找得到就不向里面插入了呢?
		// 如果现在树中存在 a, b, c. 此时查找abc肯定是找得到的
		// 如果现在又出来一个abca, 那么由a,b,c,a组成等同于abc, a组成
		// 所以满足条件的abc自然没必要进入到树中
		if root.dfs(word) {
			ans = append(ans, word)
		} else {
			root.insert(word)
		}
	}
	return
}
