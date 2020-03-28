package main

import (
	"fmt"
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}

	var res int

	sort.Strings(words)

	for i, n := 0, len(words)-1; i <= n; i++ {
		if i < n && strings.HasPrefix(words[i+1], words[i]) {
			continue
		} else {
			res += len(words[i]) + 1
		}
	}
	return res
}

func reverse(s string) string {
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

// 解法2: 字典树 未完成

type TrieNode struct {
	Val   byte
	Nodes [26]*TrieNode
}

func (t *TrieNode) Insert(s string) {
	cur := t
	for i := 1; i < len(s); i++ {
		v := s[i] - 'a'
		if cur.Nodes[v] == nil {
			cur.Nodes[v] = NewTrieNode(v)
		}
		cur = cur.Nodes[v]
	}
}

func NewTrieNode(val byte) *TrieNode {
	return &TrieNode{
		Val:   val,
		Nodes: [26]*TrieNode{},
	}
}

func NewTrie(s string) *TrieNode {
	if len(s) == 0 {
		return nil
	}
	root := NewTrieNode(' ')
	root.Insert(s)
	return root
}

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}
